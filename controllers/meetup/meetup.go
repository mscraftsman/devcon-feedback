package meetup

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/mscraftsman/devcon-feedback/app"
	"github.com/mscraftsman/devcon-feedback/models/visitor"
)

var (
	cookieName = "devcon"
	// ErrorNoToken indicates token not present in request
	ErrorNoToken = errors.New("token not found")
	// ErrorInvalidToken indicates an invalid token provided in request
	ErrorInvalidToken = errors.New("token is invalid")
)

var _key, _secret, _jwt string

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

// Init is used to initialize dependencies
func Init(meetupKey, meetupSecret, jwtSecret string) {
	_key = meetupKey
	_secret = meetupSecret
	_jwt = jwtSecret
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

		err error
	)

	err = func() error {
		client := &http.Client{Timeout: time.Second * 30}

		rs, e := client.PostForm("https://secure.meetup.com/oauth2/access",
			url.Values{
				"client_id":     {_key},
				"client_secret": {_secret},
				"grant_type":    {"authorization_code"},
				"redirect_uri":  {app.BaseURL + "/b/meetup"},
				"code":          {code},
			},
		)

		if e != nil {
			return e
		}

		resp, e := ioutil.ReadAll(rs.Body)
		if e != nil {
			return e
		}

		e = json.Unmarshal(resp, &meetupAccess)
		if e != nil {
			return e
		}

		if meetupAccess.Error != "" {
			return errors.New("meetup login error: " + meetupAccess.Error)
		}

		return nil
	}()

	if err != nil {
		return nil, err
	}

	err = func() error {
		client := &http.Client{Timeout: time.Second * 30}

		rs, e := client.Get("https://api.meetup.com/2/member/self?" +
			url.Values{
				"access_token": {meetupAccess.AccessToken},
			}.Encode(),
		)

		if e != nil {
			return e
		}

		resp, e := ioutil.ReadAll(rs.Body)
		if e != nil {
			return e
		}

		e = json.Unmarshal(resp, &meetupProfile)
		if e != nil {
			return e
		}

		return nil
	}()

	if err != nil {
		return nil, err
	}

	return &meetupProfile, nil
}

// DecodeToken returns a Profile from a request containing jwt token
func DecodeToken(r *http.Request) (*visitor.Visitor, error) {
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
		return []byte(_jwt), nil
	})

	if token == nil || err != nil {
		return nil, ErrorInvalidToken
	}

	visitor, ok := func() (*visitor.Visitor, bool) {
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

		return &visitor.Visitor{
			ID: &id, Name: &name, PhotoLink: &photoLink,
		}, idOk && nameOk && photoLinkOk
	}()

	if !ok {
		return nil, ErrorInvalidToken
	}

	return visitor, nil
}
