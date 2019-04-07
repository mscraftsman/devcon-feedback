package store

import (
	"errors"
	"time"

	"github.com/mscraftsman/devcon-feedback/config"

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
