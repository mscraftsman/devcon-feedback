package controller

import (
	"net/http"
	"strings"
	"time"

	"github.com/fluxynet/sequitur"
	"github.com/mscraftsman/devcon-feedback/config"
	"github.com/mscraftsman/devcon-feedback/meetup"
	log "github.com/sirupsen/logrus"
)

func catchError(w http.ResponseWriter, r *http.Request) sequitur.Consequence {
	return func(name string, err error) {
		var (
			msg    string
			status int
		)

		switch err {
		default:
			msg = "Error when " + name
			if strings.HasPrefix(err.Error(), "invalid") {
				status = http.StatusBadRequest
			} else {
				status = http.StatusInternalServerError
			}
		case meetup.ErrorInvalidToken:
			msg = "Invalid Token"
			status = http.StatusForbidden
		case meetup.ErrorNoToken:
			msg = "You must login to proceed"
			status = http.StatusForbidden
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		w.Write([]byte(`{"error":"` + msg + `"}`))

		log.WithField("name", name).WithField("error", err).Debug("controller:catchError")
	}
}

//IsBeforeNow checks if a t
func IsBeforeNow(t string) bool {
	var now string

	if config.Now == "_" {
		now = time.Now().In(config.Tzone).Format("2006-01-02T15:04:05")
	} else {
		now = config.Now
	}

	return t < now
}
