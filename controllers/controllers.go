package controllers

import (
	"log"
	"net/http"
	"strings"

	"github.com/mscraftsman/devcon-feedback/config"

	"github.com/fluxynet/sequitur"

	"github.com/mscraftsman/devcon-feedback/meetup"
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

		if config.Env == config.EnvironmentDev {
			log.Printf("[%s] %v\n", name, err)
		}
	}
}
