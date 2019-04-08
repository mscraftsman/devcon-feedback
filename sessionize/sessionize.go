package sessionize

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/mscraftsman/devcon-feedback/config"

	"github.com/fluxynet/sequitur"
)

//Sessions contains sessionID -> endsAt from sessionize; must only be read
var Sessions map[string]string

type sessionize struct {
	Sessions []struct {
		ID     string `json:"id"`
		EndsAt string `json:"endsAt"`
	}
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
		m := make(map[string]string)
		for i := range sess.Sessions {
			m[sess.Sessions[i].ID] = sess.Sessions[i].EndsAt
		}
		Sessions = m
	})
}
