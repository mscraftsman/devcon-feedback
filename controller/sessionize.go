package controller

import (
	"encoding/json"
	"net/http"

	"github.com/fluxynet/sequitur"
	"github.com/mscraftsman/devcon-feedback/sessionize"
)

//ListSessions returns all sessions cached
func ListSessions(w http.ResponseWriter, r *http.Request) {
	var (
		sequence sequitur.Sequence
		j        []byte
		err      error
	)

	defer sequence.Catch(catchError(w, r))

	sequence.Do("encoding response", func() error {
		j, err = json.Marshal(sessionize.Sessions)
		return err
	})

	sequence.Then(func() {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	})
}
