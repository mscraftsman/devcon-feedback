package meetup

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/mscraftsman/devcon-feedback/app"
	"github.com/mscraftsman/devcon-feedback/models/visitor"
)

// LoginCallback is an http endpoint for meetup.com auth flow
func LoginCallback(w http.ResponseWriter, r *http.Request) {
	var (
		visitor visitor.Visitor
		err     error
	)
	code := r.URL.Query().Get("code")
	profile, err := retrieveProfile(code)

	if err != nil {
		return
	}

	*visitor.ID = strconv.Itoa(profile.ID)
	*visitor.Name = profile.Name
	*visitor.PhotoLink = profile.Photo.PhotoLink

	if err = visitor.Merge(nil, true); err != nil {
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         profile.ID,
		"name":       profile.Name,
		"photo_link": profile.Photo.PhotoLink,
	})
	tokenString, err := token.SignedString([]byte(_jwt))

	if err != nil {
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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		switch err {
		default:
			msg = err.Error()
		case ErrorInvalidToken:
			msg = "Invalid Token"
		case ErrorNoToken:
			msg = "You must login to proceed"
		}

		fmt.Fprint(w, `{"status": false, "error": "`+msg+`"}`)
	}

	output, _ := json.Marshal(profile)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(output))
}
