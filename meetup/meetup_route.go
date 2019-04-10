package meetup

import (
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/fluxynet/sequitur"
	"github.com/mscraftsman/devcon-feedback/config"
	"github.com/mscraftsman/devcon-feedback/store"
	"github.com/mscraftsman/devcon-feedback/token"
)

// Login redirects to meetup url for user authentication
func Login(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, `https://secure.meetup.com/oauth2/authorize?`+url.Values{
		"client_id":     {config.MeetupKey},
		"response_type": {"code"},
		"redirect_uri":  {config.BaseURL + "/meetup"},
	}.Encode(), http.StatusFound)
}

// LoginCallback is an http endpoint for meetup.com auth flow
func LoginCallback(w http.ResponseWriter, r *http.Request) {
	var (
		sequence    sequitur.Sequence
		profile     *Profile
		attendee    store.Attendee
		tokenString string
	)

	code := r.URL.Query().Get("code")

	sequence.Do("Retrieving Profile", func() error {
		var err error
		profile, err = retrieveProfile(code)
		attendee.ID = strconv.Itoa(profile.ID)
		attendee.Name = profile.Name
		attendee.PhotoLink = profile.Photo.PhotoLink
		attendee.Status = true

		return err
	})

	sequence.Do("Saving attendee details", func() error {
		return store.DB.SetAttendee(attendee)
	})

	sequence.Do("Creating JWT", func() error {
		var err error
		tokenString, err = token.New(attendee)
		return err
	})

	sequence.Then(func() {
		cookie := http.Cookie{Name: config.CookieName, Value: tokenString, Expires: time.Now().Add(168 * time.Hour), HttpOnly: true, Path: "/"}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, config.FrontURL+"/", http.StatusFound)
	})

	sequence.Catch(func(act string, err error) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error": "` + act + `"}`))
	})
}

//Logout sends a stale cookie to clear the log in cookie
func Logout(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{Name: config.CookieName, Value: "", Expires: time.Unix(0, 0), HttpOnly: true, Path: "/"}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, config.FrontURL+"/", http.StatusFound)
}
