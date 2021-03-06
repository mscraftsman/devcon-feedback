package controller

import (
	"encoding/json"
	"net/http"

	"github.com/fluxynet/sequitur"
	"github.com/mscraftsman/devcon-feedback/store"
)

// Me is an http endpoint to get logged in user details
func Me(w http.ResponseWriter, r *http.Request) {
	var (
		attendee *store.Attendee
		err      error
		sequence sequitur.Sequence
	)

	defer sequence.Catch(catchError(w, r))

	sequence.Do("Loading attendee information", func() error {
		attendee, err = attendeeFromRequest(r)
		return err
	})

	var response []byte

	sequence.Do("Writing response", func() error {
		attendee.Bookmarks = ""
		attendee.Feedbacks = ""
		response, err = json.Marshal(attendee)
		return err
	})

	sequence.Then(func() {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(response)
	})
}
