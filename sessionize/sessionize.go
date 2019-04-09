package sessionize

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/mscraftsman/devcon-feedback/config"

	"github.com/fluxynet/sequitur"
)

var (
	// Sessions contains sessionID -> Session from sessionize; must only be read
	Sessions    map[string]Session
	rawSessions []byte
)

//Session represents a session from sessionize
type Session struct {
	ID               string   `json:"id"`
	StartsAt         string   `json:"startsAt"`
	Speakers         []string `json:"speakers"`
	IsServiceSession bool     `json:"isServiceSession"`
}

type sessionize struct {
	Sessions []Session
}

//LoadSessions from sessionize
func LoadSessions() {
	var (
		sequence sequitur.Sequence
		rs       *http.Response
		err      error
		resp     []byte
		sess     sessionize
	)

	defer func() {
		if Sessions == nil {
			Sessions = make(map[string]Session)
		}
	}()

	sequence.Do("Loading sessions from sessionize", func() error {
		client := &http.Client{Timeout: time.Second * 30}
		req, err := http.NewRequest(http.MethodGet, config.SessionizeURL, nil)
		if err != nil {
			return err
		}

		req.Header.Set("Content-type", "application/json")
		rs, err = client.Do(req)

		return err
	})

	sequence.Do("Reading response from sessionize.com during sessions loading", func() error {
		resp, err = ioutil.ReadAll(rs.Body)
		return err
	})

	sequence.Do("Decoding response from sessionize.com during sessions loading", func() error {
		return json.Unmarshal(resp, &sess)
	})

	sequence.Then(func() {
		m := make(map[string]Session)
		for i := range sess.Sessions {
			if !sess.Sessions[i].IsServiceSession {
				m[sess.Sessions[i].ID] = sess.Sessions[i]
			}
		}
		Sessions = m
		rawSessions = resp
	})
}
