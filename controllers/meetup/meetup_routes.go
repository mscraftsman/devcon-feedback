package meetup

import (
	"net/http"
	"net/url"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/mscraftsman/devcon-feedback/app"
	"github.com/mscraftsman/devcon-feedback/models/visitor"
	"github.com/mscraftsman/devcon-feedback/util"
)

// Login redirects to meetup url for user authentication
func Login(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, `https://secure.meetup.com/oauth2/authorize?`+url.Values{
		"client_id":     {_key},
		"response_type": {"code"},
		"redirect_uri":  {app.BaseURL + "/b/meetup"},
	}.Encode(), http.StatusFound)
}

// LoginCallback is an http endpoint for meetup.com auth flow
func LoginCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	profile, err := retrieveProfile(code)

	if err != nil {
		http.Redirect(w, r, app.BaseURL+`/#/loginerror`, http.StatusFound)
		return
	}

	visitor := visitor.New()
	*visitor.ID = strconv.Itoa(profile.ID)
	*visitor.Name = profile.Name
	*visitor.PhotoLink = profile.Photo.PhotoLink
	*visitor.Status = true

	if err = visitor.Merge(nil, true); err != nil {
		http.Redirect(w, r, app.BaseURL+`/#/loginerror`, http.StatusFound)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         *visitor.ID,
		"name":       *visitor.Name,
		"photo_link": *visitor.PhotoLink,
	})
	tokenString, err := token.SignedString([]byte(_jwt))

	if err != nil {
		http.Redirect(w, r, app.BaseURL+`/#/loginerror`, http.StatusFound)
		return
	}

	cookie := http.Cookie{Name: cookieName, Value: tokenString, Expires: time.Now().Add(72 * time.Hour), HttpOnly: true, Path: "/"}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, app.BaseURL+"/", http.StatusFound)
}

// Me is an http endpoint to get logged in user details
func Me(w http.ResponseWriter, r *http.Request) {
	profile, err := DecodeToken(r)

	if err != nil {
		var msg string

		switch err {
		default:
			msg = err.Error()
		case ErrorInvalidToken:
			msg = "Invalid Token"
		case ErrorNoToken:
			msg = "You must login to proceed"
		}

		util.JSONOutputError(w, http.StatusInternalServerError, msg)
		return
	}

	util.JSONOutputResponse(w, profile)
}

// Logout clears JWT cookie
func Logout(w http.ResponseWriter, r *http.Request) {
	c := http.Cookie{
		Name:   cookieName,
		MaxAge: -1}
	http.SetCookie(w, &c)

	util.JSONOutputResponse(w, nil)
}
