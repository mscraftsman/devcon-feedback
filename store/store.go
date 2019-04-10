package store

import (
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/fluxynet/sequitur"
	"github.com/mscraftsman/devcon-feedback/config"
	log "github.com/sirupsen/logrus"

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

func openDB() error {
	var (
		store Store
		err   error
	)
	store.DB, err = bolt.Open(config.BoltDBPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	DB = &store
	return err
}

//Init initialises the data store
func Init() error {
	err := openDB()
	if err != nil {
		return err
	}

	err = DB.Update(func(tx *bolt.Tx) error {
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

	DB.UpdateRatings()
  loadAttendees()

	return err
}

//ClearDB clears the database from feedback / rating. USE WITH CAUTION
func ClearDB() error {
	var (
		sequence sequitur.Sequence
		file     []byte
		err      error
	)

	sequence.Do("opening database", func() error {
		return openDB()
	})

	sequence.Do("reading database file for backup", func() error {
		file, err = ioutil.ReadFile(config.BoltDBPath)
		return err
	})

	sequence.Do("writing database backup file", func() error {
		fname := fmt.Sprintf(
			"%s_%s.bak",
			config.BoltDBPath,
			time.Now().Format("2006-01-02T15_04_05"),
		)
		return ioutil.WriteFile(fname, file, 0600)
	})

	sequence.Do("clearing database", func() error {
		return DB.Update(func(tx *bolt.Tx) error {
			if err := tx.DeleteBucket(bucketFeedbacks); err != nil && err != bolt.ErrBucketNotFound {
				return err
			}

			if _, err := tx.CreateBucket(bucketFeedbacks); err != nil {
				return err
			}

			if err := tx.DeleteBucket(bucketRatings); err != nil && err != bolt.ErrBucketNotFound {
				return err
			}

			if _, err := tx.CreateBucket(bucketRatings); err != nil {
				return err
			}

			return nil
		})
	})

	sequence.Catch(func(name string, e error) {
		err = e
		log.WithField("error", e).Error(name)
	})

	return err
}
