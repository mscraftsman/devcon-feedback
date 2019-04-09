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
		_, _ = w.Write([]byte(`{"error":"` + msg + `"}`))

		log.WithField("name", name).WithField("error", err).Debug("controller:catchError")
	}
}

//IsVoteOpen checks if vote opened
func IsVoteOpen(startsAt string) bool {
	var err error
	now := new(time.Time)
	t := new(time.Time)

	if config.Now == nil {
		*now = time.Now().In(config.Tzone)
	} else {
		now = config.Now
	}
	log.WithField("now", now.Format(config.DatetimeLayout)).Debug("IsVoteOpen:now")

	if now.After(*config.VoteClosedAt) {
		log.WithField("VoteClosedAt", config.VoteClosedAt.Format(config.DatetimeLayout)).Debug("IsVoteOpen:voteclosed")
		return false
	}

	log.WithField("startsAt", startsAt).Debug("IsVoteOpen:startsAt")
	if *t, err = time.Parse(config.DatetimeLayout, startsAt); err != nil || t == nil {
		log.WithField("startsAt", startsAt).WithField("error", err).Debug("IsVoteOpen:startsAt:parse")
		return false
	}

	open := t.Add(time.Minute * 15)
	log.WithField("time", t.Format(config.DatetimeLayout)).Debug("IsVoteOpen:votetime")

	return now.After(open)
}
