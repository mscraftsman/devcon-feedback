package sessionize

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/mscraftsman/devcon-feedback/models/session"
)

var db *sql.DB

// Session represents a session on sessionize
type Session struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartsAt    time.Time `json:"starts_at"`
	EndsAt      time.Time `json:"ends_at"`
	Room        string    `json:"room"`
	Speakers    []struct {
		Name string `json:"name"`
	} `json:"name"`
}

// Init allows injection of services into the package
func Init(database *sql.DB) {
	db = database
}

func retrieveSessions() ([]Session, error) {
	var err error

	var response struct {
		groups []struct {
			Sessions []Session `json:"sessions"`
		}
	}

	client := &http.Client{Timeout: time.Second * 30}

	rs, err := client.Get("https://sessionize.com/api/v2/351ijy5v/view/Sessions")

	if err != nil {
		return nil, err
	}

	resp, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, err
	}

	if len(response.groups) == 0 {
		return nil, errors.New("no groups found from sessionize")
	}

	return response.groups[0].Sessions, nil
}

// Sync synchronizes session from sessionize api onto database
func Sync() {
	var err error
	sessions, err := retrieveSessions()
	if err != nil {
		return
	}

	tx, err := db.Begin()
	if err != nil {
		return
	}

	for _, sess := range sessions {
		var (
			s        session.Session
			speakers []string
		)

		*s.ID = sess.ID
		*s.Title = sess.Title
		*s.Description = sess.Description
		*s.Room = sess.Room

		for _, speaker := range sess.Speakers {
			speakers = append(speakers, speaker.Name)
		}

		*s.Speakers = strings.Join(speakers, ", ")
		*s.StartAt = sess.StartsAt
		*s.EndAt = sess.EndsAt
		*s.Status = true

		s.Merge(tx, false)
	}

	tx.Commit()
}
