package meetup

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/mscraftsman/devcon-feedback/app"
	"github.com/mscraftsman/devcon-feedback/models/visitor"
)

var _key, _secret, _jwt string

// Init is used to initialize dependencies
func Init(meetupKey, meetupSecret, jwtSecret string) {
	_key = meetupKey
	_secret = meetupSecret
	_jwt = jwtSecret
}

// retrieveVisitor retrieves visitor profile information given code (automatically exchanges code for accessToken)
func retrieveVisitor(code string) (*visitor.Visitor, error) {
	var (
		meetupAccess struct {
			AccessToken string `json:"access_token"`
			TokenType   string `json:"token_type"`
			ExpiresIn   int    `json:"expires_in"`
		}

		meetupProfile struct {
			ID    int    `json:"id"`
			Name  string `json:"name"`
			Photo struct {
				PhotoLink string `json:"photo_link"`
			} `json:"photo"`
		}

		err error
	)

	client := &http.Client{Timeout: time.Second * 30}

	rs, err := client.PostForm("https://secure.meetup.com/oauth2/access",
		url.Values{
			"client_id":    {_key},
			"grant_type":   {"authorization_code"},
			"redirect_uri": {app.MeetupRedirectURL},
			"code":         {code},
		},
	)

	if err != nil {
		return nil, err
	}

	resp, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &meetupAccess)
	if err != nil {
		return nil, err
	}

	client.Get("https://api.meetup.com/2/member/self?" +
		url.Values{
			"access_token": {meetupAccess.AccessToken},
		}.Encode(),
	)

	if err != nil {
		return nil, err
	}

	resp, err = ioutil.ReadAll(rs.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &meetupProfile)
	if err != nil {
		return nil, err
	}

	v := visitor.New()
	*v.MeetupID = meetupProfile.ID
	*v.Name = meetupProfile.Name
	*v.PhotoLink = meetupProfile.Photo.PhotoLink

	err = v.Save(nil, true)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// LoginCallback is an http endpoint for meetup.com auth flow
func LoginCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	visitor, err := retrieveVisitor(code)

	if err != nil {
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         *visitor.ID,
		"name":       *visitor.Name,
		"meetup_id":  *visitor.MeetupID,
		"photo_link": *visitor.PhotoLink,
	})
	tokenString, err := token.SignedString([]byte(_jwt))

	if err != nil {
		return
	}

	cookie := http.Cookie{Name: "devcon", Value: tokenString, Expires: time.Now().Add(72 * time.Hour), HttpOnly: true, Path: "/"}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, app.BaseURL+"/", http.StatusFound)
}
