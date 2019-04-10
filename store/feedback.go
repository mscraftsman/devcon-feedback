package store

import (
	"encoding/json"
	"time"

	bolt "go.etcd.io/bbolt"
)

//Feedback for a session by an attendee
type Feedback struct {
	ID         string    `json:"id"`
	SessionID  string    `json:"session_id"`
	AttendeeID string    `json:"attendee_id"`
	Reaction1  string    `json:"reaction_1"`
	Reaction2  string    `json:"reaction_2"`
	Reaction3  string    `json:"reaction_3"`
	Reaction4  string    `json:"reaction_4"`
	CreatedAt  time.Time `json:"created_at"`
}

//AddFeedback to store
func (s Store) AddFeedback(f Feedback) error {
	attendee, err := s.GetAttendee(f.AttendeeID)
	if err != nil {
		return err
	}

	f.ID = f.SessionID + ":" + f.AttendeeID
	f.CreatedAt = time.Now()

	err = s.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketFeedbacks)
		r := b.Get([]byte(f.ID))
		if r == nil {
			return nil
		}
		return ErrorFeedbackExists
	})
	if err != nil {
		return err
	}

	attendee = attendee.AddFeedback(f.ID)
	a, err := json.Marshal(attendee)
	if err != nil {
		return err
	}

	j, err := json.Marshal(f)
	if err != nil {
		return err
	}

	_ = s.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketFeedbacks)

		if err := b.Put([]byte(f.ID), j); err != nil {
			return err
		}

		b = tx.Bucket(bucketAttendees)
		return b.Put([]byte(attendee.ID), a)
	})

	go s.UpdateRatings()
	return err
}

//GetFeedback retrieves a feedback from the store
func (s Store) GetFeedback(id string) (Feedback, error) {
	var f Feedback
	err := s.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketFeedbacks)
		raw := b.Get([]byte(id))

		if raw == nil {
			return ErrorFeedbackNotFound
		}

		return json.Unmarshal(raw, &f)
	})

	return f, err
}

//ListAttendeeFeedbacks for an attendee ID
func (s Store) ListAttendeeFeedbacks(attnID string) ([]Feedback, error) {
	var feedbacks []Feedback

	attn, err := s.GetAttendee(attnID)
	if err != nil {
		return feedbacks, err
	}

	fIDs := attn.ListFeedbacks()

	_ = s.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketFeedbacks)
		for i := range fIDs {
			var f Feedback
			raw := b.Get([]byte(fIDs[i]))

			if raw == nil {
				continue
			}

			if err := json.Unmarshal(raw, &f); err == nil {
				feedbacks = append(feedbacks, f)
			}
		}

		return nil
	})

	return feedbacks, nil
}

//ListFeedbacks for all attendees
func (s Store) ListFeedbacks() []Feedback {
	var feedbacks []Feedback

	_ = s.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketFeedbacks)
		_ = b.ForEach(func(key, raw []byte) error {
			var f Feedback
			if err := json.Unmarshal(raw, &f); err == nil {
				feedbacks = append(feedbacks, f)
			}
			return nil
		})
		return nil
	})

	return feedbacks
}
