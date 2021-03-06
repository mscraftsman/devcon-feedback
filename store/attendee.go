package store

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strings"

	"github.com/fluxynet/sequitur"
	log "github.com/sirupsen/logrus"
	bolt "go.etcd.io/bbolt"
)

var (
	//attendeesList is a store of attendees that are allowed to login
	attendeesList map[string]interface{}

	errNoAttendeesFile = errors.New("no attendees file")
)

//Attendee represents a devcon attendee
type Attendee struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	PhotoLink string `json:"photo_link"`
	Status    bool   `json:"status"`
	Bookmarks string `json:"bookmarks"`
	Feedbacks string `json:"feedbacks"`
}

//AddBookmark creates a new bookmark to an attendee
func (a Attendee) AddBookmark(id string) Attendee {
	b := strings.Split(a.Bookmarks, ";")
	for i := range b {
		if b[i] == id {
			return a
		}
	}

	if a.Bookmarks == "" {
		a.Bookmarks = id
	} else if !strings.Contains(a.Bookmarks, id) {
		a.Bookmarks = a.Bookmarks + ";" + id
	}

	return a
}

//RemoveBookmark removes a bookmark from an attendee
func (a Attendee) RemoveBookmark(id string) Attendee {
	var n []string
	b := strings.Split(a.Bookmarks, ";")
	for i := range b {
		if b[i] != id {
			n = append(n, b[i])
		}
	}

	a.Bookmarks = strings.Join(n, ";")
	return a
}

//ListBookmarks returns list of bookmarks
func (a Attendee) ListBookmarks() []string {
	if a.Bookmarks == "" {
		return []string{}
	}
	return strings.Split(a.Bookmarks, ";")
}

//AddFeedback adds a feedback id to the attendee (immutable)
func (a Attendee) AddFeedback(id string) Attendee {
	f := strings.Split(a.Feedbacks, ";")
	for i := range f {
		if f[i] == id {
			return a
		}
	}

	if a.Feedbacks == "" {
		a.Feedbacks = id
	} else if !strings.Contains(a.Feedbacks, id) {
		a.Feedbacks = a.Feedbacks + ";" + id
	}

	return a
}

//ListFeedbacks returns feedback ids for an attendee
func (a Attendee) ListFeedbacks() []string {
	if a.Feedbacks == "" {
		return []string{}
	}
	return strings.Split(a.Feedbacks, ";")
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
func (s Store) SetAttendee(a Attendee, overwrite ...bool) error {
	if len(overwrite) != 0 && overwrite[0] {
		//overwrite
	} else {
		//if already exists, we preserve existing and update only name and photo
		if e, err := s.GetAttendee(a.ID); err == nil {
			e.Name = a.Name
			e.PhotoLink = a.PhotoLink
			a = e
		}
	}

	j, err := json.Marshal(a)
	if err != nil {
		return err
	}

	return s.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketAttendees)
		return b.Put([]byte(a.ID), j)
	})
}

//ListAttendees returns a list of all attendees
func (s Store) ListAttendees() []Attendee {
	var attendees []Attendee

	_ = s.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketAttendees)

		_ = b.ForEach(func(k, v []byte) error {
			var a Attendee

			if err := json.Unmarshal(v, &a); err == nil {
				attendees = append(attendees, a)
			}

			return nil
		})

		return nil
	})

	return attendees
}

//IsAttendeeBanned returns a func to check if attendee is banned by id
func (s Store) IsAttendeeBanned() func(string) bool {
	a := s.ListAttendees()
	b := make(map[string]interface{})

	for i := range a {
		if !a[i].Status {
			b[a[i].ID] = nil
		}
	}
	log.WithField("banned", b).Debug("store:IsAttendeeBanned")

	return func(id string) bool {
		_, ok := b[id]
		return ok
	}
}

//IsValidAttendee check if attendee is known
func IsValidAttendee(id string) bool {
	if attendeesList == nil {
		log.WithField("id", id).Debug("IsValidAttendee:list:nil")
		return true
	}

	_, ok := attendeesList[id]
	log.WithField("id", id).WithField("ok", ok).Debug("IsValidAttendee")
	return ok
}

func loadAttendees() {
	var (
		err      error
		sequence sequitur.Sequence
		ids      map[string]interface{}
		file     struct {
			AttendeeIDS []string `json:"attendee_ids"`
		}
	)

	defer func() {
		if err := recover(); err != nil {
			log.WithField("error", err).Debug("attendees:load:panic")
		}

		sequence.Catch(func(name string, err error) {
			if err == errNoAttendeesFile {
				attendeesList = nil
				log.Debug("no attendees file found")
			} else {
				log.WithField("error", err).Fatalln(name)
			}
		})
	}()

	sequence.Do("checking if attendees file present", func() error {
		if _, err := os.Stat("attendees.json"); os.IsNotExist(err) {
			return errNoAttendeesFile
		}

		return nil
	})

	var raw []byte
	sequence.Do("reading attendees file", func() error {
		raw, err = ioutil.ReadFile("attendees.json")
		return err
	})

	sequence.Do("decoding attendees file", func() error {
		return json.Unmarshal(raw, &file)
	})

	sequence.Then(func() {
		ids = make(map[string]interface{})
		for _, id := range file.AttendeeIDS {
			ids[id] = nil
		}

		attendeesList = ids
	})
}
