package store

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/mscraftsman/devcon-feedback/config"
	uuid "github.com/satori/go.uuid"

	bolt "go.etcd.io/bbolt"
)

//DB represents the global store
var DB *Store

var (
	bucketAttendees = []byte("Attendee")
	bucketFeedbacks = []byte("Feedback")
	bucketRatings   = []byte("Rating")
)

var (
	//ErrorAttendeeNotFound error if AttendeeNotFound
	ErrorAttendeeNotFound = errors.New("attendee not found")
	//ErrorFeedbackExists error if FeedbackExists
	ErrorFeedbackExists = errors.New("attendee feedback already exists for this session")
	//ErrorFeedbackNotFound error if FeedbackNotFound
	ErrorFeedbackNotFound = errors.New("feedback not found")
)

//Store is for storing things
type Store struct {
	*bolt.DB
}

//Init initialises the data store
func Init() error {
	var store Store
	db, err := bolt.Open(config.BoltDBPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists(bucketAttendees); err != nil {
			return err
		}

		if _, err := tx.CreateBucketIfNotExists(bucketFeedbacks); err != nil {
			return err
		}

		if _, err := tx.CreateBucketIfNotExists(bucketRatings); err != nil {
			return err
		}

		return nil
	})

	store.DB = db
	DB = &store
	return err
}

//GetAttendee retrieves an attendee from the store
func (s Store) GetAttendee(id string) (Attendee, error) {
	var a Attendee
	err := s.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketAttendees)
		j := b.Get([]byte(id))

		if j == nil {
			return ErrorAttendeeNotFound
		}

		return json.Unmarshal(j, &a)
	})

	return a, err
}

//SetAttendee sets an attendee entity in the store
func (s Store) SetAttendee(a Attendee) error {
	j, err := json.Marshal(a)
	if err != nil {
		return err
	}

	return s.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketAttendees)
		return b.Put([]byte(a.ID), j)
	})
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

