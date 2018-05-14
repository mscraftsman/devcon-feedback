package meetup

import (
	"net/http"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/mscraftsman/devcon-feedback/app"
	"github.com/mscraftsman/devcon-feedback/models/visitor"
	"github.com/mscraftsman/devcon-feedback/util"
)

// LoginCallback is an http endpoint for meetup.com auth flow
func LoginCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	profile, err := retrieveProfile(code)

	if err != nil {
		util.JSONOutputError(w, http.StatusUnauthorized, err.Error())
		return
	}

	visitor := visitor.New()
	*visitor.ID = strconv.Itoa(profile.ID)
	*visitor.Name = profile.Name
	*visitor.PhotoLink = profile.Photo.PhotoLink

	if err = visitor.Merge(nil, true); err != nil {
		util.JSONOutputError(w, http.StatusInternalServerError, err.Error())
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         *visitor.ID,
		"name":       *visitor.Name,
		"photo_link": *visitor.PhotoLink,
	})
	tokenString, err := token.SignedString([]byte(_jwt))

	if err != nil {
		util.JSONOutputError(w, http.StatusInternalServerError, err.Error())
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
