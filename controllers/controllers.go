package controllers

import (
	"log"
	"net/http"

	"github.com/mscraftsman/devcon-feedback/config"

	"github.com/fluxynet/sequitur"

	"github.com/mscraftsman/devcon-feedback/meetup"
)

func catchError(w http.ResponseWriter, r *http.Request) sequitur.Consequence {
	return func(name string, err error) {
		var msg string

		switch err {
		default:
			msg = "Error when " + name
		case meetup.ErrorInvalidToken:
			msg = "Invalid Token"
		case meetup.ErrorNoToken:
			msg = "You must login to proceed"
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(`{"error":"` + msg + `"}`))

		if config.Env == config.EnvironmentDev {
			log.Printf("[%s] %v\n", name, err)
		}
	}
}
