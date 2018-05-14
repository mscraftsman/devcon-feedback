// generated by gocipe 8711bda2250e6c55fba92576b5a3047c42d034e6742f4bff5ae1f9e546907463; DO NOT EDIT

package feedback

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type responseSingle struct {
	Status   bool      `json:"status"`
	Messages []message `json:"messages"`
	Entity   *Feedback `json:"entity"`
}

type responseList struct {
	Status   bool        `json:"status"`
	Messages []message   `json:"messages"`
	Entities []*Feedback `json:"entities"`
}

type message struct {
	Type    rune   `json:"type"`
	Message string `json:"message"`
}

//RegisterRoutes registers routes with a mux Router
func RegisterRoutes(router *mux.Router) {

	router.HandleFunc("/feedbacks", RestCreate).Methods("POST")

}

//RestCreate is a REST endpoint for POST /feedbacks
func RestCreate(w http.ResponseWriter, r *http.Request) {
	var (
		err      error
		rawbody  []byte
		response responseSingle
		tx       *sql.Tx
	)

	rawbody, err = ioutil.ReadAll(r.Body)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"status": false, "messages": [{"type": "error", "text": "Failed to read body"}]}`)
		return
	}

	response.Entity = New()
	err = json.Unmarshal(rawbody, response.Entity)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"status": false, "messages": [{"type": "error", "text": "Failed to decode body"}]}`)
		return
	}
	response.Entity.ID = nil

	tx, err = db.Begin()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"status": false, "messages": [{"type": "error", "text": "Failed to process"}]}`)
		return
	}

	err = response.Entity.Save(tx, false)
	if err != nil {
		tx.Rollback()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"status": false, "messages": [{"type": "error", "text": "Save failed"}]}`)
		return
	}

	if err = tx.Commit(); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"status": false, "messages": [{"type": "E", "message": "RestCreate could not commit transaction"}]}`)
		return
	}

	output, err := json.Marshal(response)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"status": false, "messages": [{"type": "error", "text": "JSON encoding failed"}]}`)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(output))
}