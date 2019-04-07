package sessionize

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

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
		rs, err = client.Get("https://sessionize.com/api/v2/i2m69kbq/view/All")

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
