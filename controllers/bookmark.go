package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/mscraftsman/devcon-feedback/sessionize"

	"github.com/fluxynet/sequitur"
	"github.com/mscraftsman/devcon-feedback/meetup"
	"github.com/mscraftsman/devcon-feedback/store"
)

//AddBookmark allows a user to add a bookmark
func AddBookmark(w http.ResponseWriter, r *http.Request) {
	var (
		attendee *store.Attendee
		err      error
		sequence sequitur.Sequence
		request  struct {
			ID string `json:"id"`
		}
	)

	defer sequence.Catch(catchError(w, r))

	sequence.Do("loading attendee informaton", func() error {
		attendee, err = meetup.DecodeToken(r)
		return err
	})

	var resp []byte
	sequence.Do("reading request body", func() error {
		resp, err = ioutil.ReadAll(r.Body)
		return err
	})

	sequence.Do("decoding request body", func() error {
		return json.Unmarshal(resp, &request)
	})

	sequence.Do("validating bookmark information", func() error {
		if _, ok := sessionize.Sessions[request.ID]; !ok {
			return errors.New("invalid session")
		}
		return nil
	})

	sequence.Do("saving bookmark", func() error {
		attn, err := store.DB.GetAttendee(attendee.ID)
		if err == nil {
			return store.DB.SetAttendee(attn.AddBookmark(request.ID))
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
		attendee *store.Attendee
		err      error
		sequence sequitur.Sequence
		request  struct {
			ID string `json:"id"`
		}
	)

	defer sequence.Catch(catchError(w, r))

	sequence.Do("loading attendee informaton", func() error {
		attendee, err = meetup.DecodeToken(r)
		return err
	})

	var resp []byte
	sequence.Do("reading request body", func() error {
		resp, err = ioutil.ReadAll(r.Body)
		return err
	})

	sequence.Do("decoding request body", func() error {
		return json.Unmarshal(resp, &request)
	})

	sequence.Do("validating bookmark information", func() error {
		if _, ok := sessionize.Sessions[request.ID]; !ok {
			return errors.New("invalid session")
		}
		return nil
	})

	sequence.Do("removing bookmark", func() error {
		attn, err := store.DB.GetAttendee(attendee.ID)
		if err == nil {
			return store.DB.SetAttendee(attn.RemoveBookmark(request.ID))
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
	sequence.Do("loading attendee informaton", func() error {
		attendee, err = meetup.DecodeToken(r)
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
		w.Write(rs)
	})
}
