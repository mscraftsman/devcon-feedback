package main

import (
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
	"github.com/mscraftsman/devcon-feedback/app"
	"github.com/mscraftsman/devcon-feedback/controllers/meetup"
	"github.com/mscraftsman/devcon-feedback/models/feedback"
	"github.com/mscraftsman/devcon-feedback/models/rating"
	"github.com/mscraftsman/devcon-feedback/models/session"
	"github.com/mscraftsman/devcon-feedback/models/visitor"
	"github.com/mscraftsman/devcon-feedback/sessionize"
)

//go:generate rice embed-go

func main() {
	config := app.Bootstrap()

	inject(config)

	// go func() {
	// 	sessionize.Sync()
	// 	time.Sleep(time.Minute * 5)
	// }()

	assetsHandler := http.FileServer(rice.MustFindBox("web/webapp-client/dist").HTTPBox())

	app.ServeHTTP(":"+config.HTTPPort, func(router *mux.Router) error {
		router.HandleFunc("/b/meetup", meetup.LoginCallback).Methods(http.MethodGet)
		router.HandleFunc("/b/me", meetup.Me).Methods(http.MethodGet)
		router.HandleFunc("/b/api/feedbacks", feedback.RestCreate).Methods("POST")
		router.PathPrefix("/").Handler(assetsHandler)
		return nil
	})
}

func inject(config *app.Config) {
	meetup.Init(config.MeetupKey, config.MeetupSecret, config.JWTSecret)
	sessionize.Init(config.DB)
	feedback.Inject(config.DB)
	rating.Inject(config.DB)
	session.Inject(config.DB)
	visitor.Inject(config.DB)
}
