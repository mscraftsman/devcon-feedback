package main

import (
	"log"
	"net/http"
	"time"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
	"github.com/mscraftsman/devcon-feedback/app"
	"github.com/mscraftsman/devcon-feedback/controllers/meetup"
	"github.com/mscraftsman/devcon-feedback/models/feedback"
	"github.com/mscraftsman/devcon-feedback/models/rating"
	"github.com/mscraftsman/devcon-feedback/models/visitor"
	"github.com/mscraftsman/devcon-feedback/sessionize"
)

//go:generate rice embed-go

func main() {
	config := app.Bootstrap()

	inject(config)

	assetsHandler := http.FileServer(rice.MustFindBox("web/webapp-client/dist").HTTPBox())

	if err := sessionize.Sync(); err != nil {
		log.Fatalf("Failed to load sessions information: %s", err)
	} else {
		log.Println("Sessionize data loaded")
	}

	go func() {
		time.Sleep(time.Minute * 5)
		sessionize.Sync()
	}()

	app.ServeHTTP(":"+config.HTTPPort, func(router *mux.Router) error {
		router.HandleFunc("/b/login", meetup.Login).Methods(http.MethodGet)
		router.HandleFunc("/b/meetup", meetup.LoginCallback).Methods(http.MethodGet)
		router.HandleFunc("/b/me", meetup.Me).Methods(http.MethodGet)
		router.HandleFunc("/b/api/feedbacks/me", feedback.MyFeedback).Methods(http.MethodGet)
		router.HandleFunc("/b/api/feedbacks", feedback.RestCreate).Methods(http.MethodPost)
		router.HandleFunc("/b/api/speakers/{id}", sessionize.GetSpeaker).Methods(http.MethodGet)
		router.HandleFunc("/b/api/speakers", sessionize.GetSpeakers).Methods(http.MethodGet)
		router.HandleFunc("/b/api/sessions/{id}", sessionize.GetSession).Methods(http.MethodGet)
		router.HandleFunc("/b/api/sessions", sessionize.GetSessions).Methods(http.MethodGet)
		router.HandleFunc("/b/api/rooms/{id}", sessionize.GetRoom).Methods(http.MethodGet)
		router.HandleFunc("/b/api/rooms", sessionize.GetRooms).Methods(http.MethodGet)
		router.PathPrefix("/").Handler(assetsHandler)
		return nil
	})
}

func inject(config *app.Config) {
	meetup.Init(config.MeetupKey, config.MeetupSecret, config.JWTSecret)
	feedback.Inject(config.DB)
	rating.Inject(config.DB)
	visitor.Inject(config.DB)
}
