package store

import (
	"encoding/json"
	"time"

	uuid "github.com/satori/go.uuid"
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
	j, err := json.Marshal(f)
	if err != nil {
		return err
	}

	id := uuid.NewV4().String()

	attn, err := s.GetAttendee(f.AttendeeID)
	if err != nil {
		return err
	}

	id += ":" + f.AttendeeID

	err = s.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketFeedbacks)
		r := b.Get([]byte(id))
		if r == nil {
			return nil
		}
		return ErrorFeedbackExists
	})

	if err != nil {
		return err
	}

	attn = attn.AddFeedback(id)
	a, err := json.Marshal(attn)
	if err != nil {
		return err
	}

	return s.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketFeedbacks)

		if err := b.Put([]byte(id), j); err != nil {
			return err
		}

		b = tx.Bucket(bucketAttendees)
		return b.Put([]byte(attn.ID), a)
	})
}

//GetFeedback retrieves a feedback from the store
func (s Store) GetFeedback(id string) (Feedback, error) {
	var f Feedback
	err := s.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketFeedbacks)
		j := b.Get([]byte(id))

		if j == nil {
			return ErrorFeedbackNotFound
		}

		return json.Unmarshal(j, &f)
	})

	return f, err
}

//ListFeedbacks for an attendee ID
func (s Store) ListFeedbacks(attnID string) ([]Feedback, error) {
	var feedbacks [][]byte

	attn, err := s.GetAttendee(attnID)
	if err != nil {
		return nil, nil
	}

	fIDs := attn.ListFeedbacks()

	s.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketFeedbacks)
		for i := range fIDs {
			if f := b.Get([]byte(fIDs[i])); f != nil {
				feedbacks = append(feedbacks, f)
			}
		}

		return nil
	})

	f := func(j [][]byte) []Feedback {
		var (
			out []Feedback
			f   Feedback
		)
		for i := range j {
			if j[i] != nil {
				if err := json.Unmarshal(j[i], &f); err == nil {
					out = append(out, f)
				}
			}
		}

		return out
	}(feedbacks)

	return f, nil
}

