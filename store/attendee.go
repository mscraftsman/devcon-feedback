package store

import (
	"strings"
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
	a.Bookmarks = a.Bookmarks + ";" + id
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
	a.Feedbacks = a.Feedbacks + ";" + id
	return a
}

//ListFeedbacks returns feedback ids for an attendee
func (a Attendee) ListFeedbacks() []string {
	return strings.Split(a.Feedbacks, ";")
}
