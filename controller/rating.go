package controller

import (
	"encoding/json"
	"net/http"

	"github.com/fluxynet/sequitur"
	"github.com/mscraftsman/devcon-feedback/store"
	log "github.com/sirupsen/logrus"
)

//ListRatings returns all ratings
func ListRatings(w http.ResponseWriter, r *http.Request) {
	var (
		j        []byte
		err      error
		sequence sequitur.Sequence
		ratings  []store.Rating
	)

	sequence.Catch(catchError(w, r))

	sequence.Do("loading ratings data", func() error {
		ratings = store.DB.ListRatings()
		log.WithField("ratings", ratings).Debug("ratings:list")
		return nil
	})

	sequence.Do("encoding response", func() error {
		j, err = json.Marshal(ratings)
		return err
	})

	sequence.Then(func() {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	})
}

//Leaderboards returns leaderboard
func Leaderboards(w http.ResponseWriter, r *http.Request) {
	var (
		sequence sequitur.Sequence
		j        []byte
		err      error
	)

	defer sequence.Catch(catchError(w, r))

	sequence.Do("encoding response", func() error {
		j, err = json.Marshal(store.Leaderboard)
		return err
	})

	sequence.Then(func() {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	})
}
