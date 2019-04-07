package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/fluxynet/sequitur"
	"github.com/mscraftsman/devcon-feedback/meetup"
	"github.com/mscraftsman/devcon-feedback/sessionize"
	"github.com/mscraftsman/devcon-feedback/store"
)

//AddFeedback allows a user to add feedback for a session
func AddFeedback(w http.ResponseWriter, r *http.Request) {
	var (
		attendee *store.Attendee
		err      error
		sequence sequitur.Sequence
		feedback store.Feedback
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
		err = json.Unmarshal(resp, &feedback)
		if err == nil {
			feedback.AttendeeID = attendee.ID
		}
		return err
	})

	sequence.Do("validating feedback information", func() error {
		if _, ok := sessionize.Sessions[feedback.SessionID]; !ok {
			return errors.New("invalid session")
		}

		if areResponsesInvalid(feedback) {
			return errors.New("invalid feedback")
		}
		return nil
	})

	sequence.Do("saving feedback information", func() error {
		return store.DB.AddFeedback(feedback)
	})

	sequence.Then(func() {
		w.WriteHeader(http.StatusOK)
	})
}

//ListOwnFeedback allows a user to list his/her own feedback
func ListOwnFeedback(w http.ResponseWriter, r *http.Request) {
	var (
		attendee *store.Attendee
		err      error
		sequence sequitur.Sequence
		response struct {
			Feedbacks []store.Feedback `json:"feedbacks"`
		}
	)

	defer sequence.Catch(catchError(w, r))

	sequence.Do("loading attendee informaton", func() error {
		attendee, err = meetup.DecodeToken(r)
		return err
	})

	sequence.Do("loading feedback information", func() error {
		response.Feedbacks, err = store.ListFeedbacks(attendee.ID)
		return err
	})

	var j []byte
	sequence.Do("writing response", func() error {
		j, err = json.Marshal(response)
		return err
	})
}

func areResponsesInvalid(f store.Feedback) bool {
	switch "" {
	case f.Reaction1,
		f.Reaction2,
		f.Reaction3,
		f.Reaction4:
		return true
	}

	switch f.Reaction1 {
	case "-2", "-1", "1", "2", "3":
	default:
		return true
	}

	switch f.Reaction2 {
	case "-2", "-1", "1", "2", "3":
	default:
		return true
	}

	switch f.Reaction3 {
	case "no", "yes":
	default:
		return true
	}

	return false
}
