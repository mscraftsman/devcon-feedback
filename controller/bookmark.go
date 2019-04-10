package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mscraftsman/devcon-feedback/sessionize"

	"github.com/fluxynet/sequitur"
	"github.com/mscraftsman/devcon-feedback/store"
)

//AddBookmark allows a user to add a bookmark
func AddBookmark(w http.ResponseWriter, r *http.Request) {
	var (
		attendee   *store.Attendee
		err        error
		sequence   sequitur.Sequence
		bookmarkID string
		vars       = mux.Vars(r)
	)

	defer sequence.Catch(catchError(w, r))

	sequence.Do("Loading attendee information", func() error {
		attendee, err = attendeeFromRequest(r)
		return err
	})

	sequence.Do("validating bookmark information", func() error {
		var ok bool
		if bookmarkID, ok = vars["id"]; !ok {
			return errors.New("invalid session")
		}

		if _, ok = sessionize.Sessions[bookmarkID]; !ok {
			return errors.New("invalid session")
		}
		return nil
	})

	sequence.Do("adding bookmark", func() error {
		attn, err := store.DB.GetAttendee(attendee.ID)
		if err == nil {
			return store.DB.SetAttendee(attn.AddBookmark(bookmarkID))
		}
		return err
	})

	sequence.Do("saving bookmark", func() error {
		attn, err := store.DB.GetAttendee(attendee.ID)
		if err == nil {
			return store.DB.SetAttendee(attn.AddBookmark(bookmarkID))
		}
		return err
	})

	sequence.Then(func() {
		w.WriteHeader(http.StatusOK)
	})
}

//RemoveBookmark allows a user to remove a bookmark
func RemoveBookmark(w http.ResponseWriter, r *http.Request) {
	var (
		attendee   *store.Attendee
		err        error
		sequence   sequitur.Sequence
		vars       = mux.Vars(r)
		bookmarkID string
	)

	defer sequence.Catch(catchError(w, r))

	sequence.Do("Loading attendee information", func() error {
		attendee, err = attendeeFromRequest(r)
		return err
	})

	sequence.Do("validating bookmark information", func() error {
		var ok bool
		if bookmarkID, ok = vars["id"]; !ok {
			return errors.New("invalid session")
		}

		if _, ok = sessionize.Sessions[bookmarkID]; !ok {
			return errors.New("invalid session")
		}
		return nil
	})

	sequence.Do("removing bookmark", func() error {
		attn, err := store.DB.GetAttendee(attendee.ID)
		if err == nil {
			return store.DB.SetAttendee(attn.RemoveBookmark(bookmarkID))
		}
		return err
	})

	sequence.Then(func() {
		w.WriteHeader(http.StatusOK)
	})
}

//ListBookmarks allows a user to list his/her own bookmarks
func ListBookmarks(w http.ResponseWriter, r *http.Request) {
	var (
		err      error
		attendee *store.Attendee
		sequence sequitur.Sequence
		response struct {
			Bookmarks []string `json:"bookmarks"`
		}
	)

	defer sequence.Catch(catchError(w, r))
	sequence.Do("Loading attendee information", func() error {
		attendee, err = attendeeFromRequest(r)
		return err
	})

	sequence.Do("loading bookmarks", func() error {
		attn, err := store.DB.GetAttendee(attendee.ID)
		if err == nil {
			response.Bookmarks = attn.ListBookmarks()
		}
		return err
	})

	var rs []byte
	sequence.Do("encoding response", func() error {
		rs, err = json.Marshal(response)
		return err
	})

	sequence.Then(func() {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(rs)
	})
}
