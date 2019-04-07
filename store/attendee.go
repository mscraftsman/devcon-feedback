package store

import (
	"encoding/json"
	"strings"

	bolt "go.etcd.io/bbolt"
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
	} else {
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
	} else {
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
