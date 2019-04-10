package meetup

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/fluxynet/sequitur"
	"github.com/mscraftsman/devcon-feedback/config"
)

// Profile represents meetup.com user profile
type Profile struct {
	ID    int          `json:"id"`
	Name  string       `json:"name"`
	Photo ProfilePhoto `json:"photo"`
}

// ProfilePhoto contains profile photo information
type ProfilePhoto struct {
	// PhotoLink represents link to the profile photo
	PhotoLink string `json:"photo_link"`
}

// retrieveProfile retrieves visitor profile information given code (automatically exchanges code for accessToken)
func retrieveProfile(code string) (*Profile, error) {
	var (
		meetupAccess struct {
			AccessToken string `json:"access_token"`
			TokenType   string `json:"token_type"`
			ExpiresIn   int    `json:"expires_in"`
			Error       string `json:"error"`
		}

		meetupProfile Profile

		sequence sequitur.Sequence
		err      error
	)

	var rs *http.Response
	sequence.Do("Authorizing with meetup.com", func() error {
		client := &http.Client{Timeout: time.Second * 30}

		rs, err = client.PostForm("https://secure.meetup.com/oauth2/access",
			url.Values{
				"client_id":     {config.MeetupKey},
				"client_secret": {config.MeetupSecret},
				"grant_type":    {"authorization_code"},
				"redirect_uri":  {config.BaseURL + "/meetup"},
				"code":          {code},
			},
		)

		return err
	})

	var resp []byte
	sequence.Do("Reading response from meetup.com during authorization", func() error {
		resp, err = ioutil.ReadAll(rs.Body)
		return err
	})

	sequence.Do("Decoding response from meetup.com during authorization", func() error {
		return json.Unmarshal(resp, &meetupAccess)
	})

	sequence.Do("Signing in with meetup.com", func() error {
		if meetupAccess.Error == "" {
			return nil
		}
		return errors.New("meetup login error: " + meetupAccess.Error)
	})

	sequence.Do("Loading profile from meetup.com", func() error {
		client := &http.Client{Timeout: time.Second * 30}
		query := url.Values{
			"access_token": {meetupAccess.AccessToken},
		}.Encode()

		rs, err = client.Get("https://api.meetup.com/2/member/self?" + query)

		return err
	})

	sequence.Do("Reading response from meetup.com during profile loading", func() error {
		resp, err = ioutil.ReadAll(rs.Body)
		return err
	})

	sequence.Do("Decoding response from meetup.com during profile loading", func() error {
		return json.Unmarshal(resp, &meetupProfile)
	})

	sequence.Catch(func(name string, e error) {
		err = e
	})

	sequence.Then(func() {
		err = nil
	})

	return &meetupProfile, err
}
