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
	"github.com/mscraftsman/devcon-feedback/store"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	cookieName = "devcon"
	// ErrorNoToken indicates token not present in request
	ErrorNoToken = errors.New("token not found")
	// ErrorInvalidToken indicates an invalid token provided in request
	ErrorInvalidToken = errors.New("token is invalid")
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
				"redirect_uri":  {config.BaseURL + "/b/meetup"},
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

// DecodeToken returns a Profile from a request containing jwt token
func DecodeToken(r *http.Request) (*store.Attendee, error) {
	var (
		tokenString string
		cookie      *http.Cookie
		err         error
	)

	if cookie, err = r.Cookie(cookieName); err != nil {
		return nil, ErrorNoToken
	}

	tokenString = cookie.Value

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWTSecret), nil
	})

	if token == nil || err != nil {
		return nil, ErrorInvalidToken
	}

	attendee, ok := func() (*store.Attendee, bool) {
		var (
			id          string
			idOk        bool
			name        string
			nameOk      bool
			photoLink   string
			photoLinkOk bool
		)
		claims, claimok := token.Claims.(jwt.MapClaims)

		if claimok && token.Valid {
			id, idOk = claims["id"].(string)
			name, nameOk = claims["name"].(string)
			photoLink, photoLinkOk = claims["photo_link"].(string)
		}

		return &store.Attendee{
			ID: id, Name: name, PhotoLink: photoLink,
		}, idOk && nameOk && photoLinkOk
	}()

	if !ok {
		return nil, ErrorInvalidToken
	}

	return attendee, nil
}
