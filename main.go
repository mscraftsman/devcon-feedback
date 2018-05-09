package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mscraftsman/devcon-feedback/app"
	"github.com/mscraftsman/devcon-feedback/controllers/meetup"
)

func main() {
	config := app.Bootstrap()

	meetup.Init(config.MeetupKey, config.MeetupSecret, config.JWTSecret)

	app.ServeHttp(config.HttpPort, func(router *mux.Router) error {
		router.HandleFunc("/b/meetup", meetup.LoginCallback).Methods(http.MethodGet)
		return nil
	})
}
