package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mscraftsman/devcon-feedback/app"
	"github.com/mscraftsman/devcon-feedback/controllers/meetup"
	"github.com/mscraftsman/devcon-feedback/models/feedback"
	"github.com/mscraftsman/devcon-feedback/models/rating"
	"github.com/mscraftsman/devcon-feedback/models/visitor"
	"github.com/mscraftsman/devcon-feedback/sessionize"
	"github.com/mscraftsman/devcon-feedback/stats"
)

func main() {
	config := app.Bootstrap()

	inject(config)

	if err := sessionize.Sync(); err != nil {
		log.Fatalf("Failed to load sessions information: %s", err)
	} else {
		log.Println("Sessionize data loaded")
	}

	go func() {
		for {
			time.Sleep(time.Minute * 5)
			sessionize.Sync()
		}
	}()

	app.ServeHTTP(":"+config.HTTPPort, func(router *mux.Router) error {
		router.Use(corsMiddleware)
		router.HandleFunc("/b/login", meetup.Login).Methods(http.MethodGet)
		router.HandleFunc("/b/meetup", meetup.LoginCallback).Methods(http.MethodGet)
		router.HandleFunc("/b/api/me", meetup.Me).Methods(http.MethodGet)
		router.HandleFunc("/b/api/feedbacks/me", feedback.MyFeedback).Methods(http.MethodGet)
		router.HandleFunc("/b/api/feedbacks", feedback.RestCreate).Methods(http.MethodPost)
		router.HandleFunc("/b/api/speakers/{id}", sessionize.GetSpeaker).Methods(http.MethodGet)
		router.HandleFunc("/b/api/speakers", sessionize.GetSpeakers).Methods(http.MethodGet)
		router.HandleFunc("/b/api/sessions/{id}", sessionize.GetSession).Methods(http.MethodGet)
		router.HandleFunc("/b/api/sessions", sessionize.GetSessions).Methods(http.MethodGet)
		router.HandleFunc("/b/api/rooms/{id}", sessionize.GetRoom).Methods(http.MethodGet)
		router.HandleFunc("/b/api/rooms", sessionize.GetRooms).Methods(http.MethodGet)
		router.HandleFunc("/b/api/stats", stats.Summary).Methods(http.MethodGet)
		return nil
	})
}

func inject(config *app.Config) {
	meetup.Init(config.MeetupKey, config.MeetupSecret, config.JWTSecret)
	feedback.Inject(config.DB)
	rating.Inject(config.DB)
	visitor.Inject(config.DB)
	stats.Inject(config.DB)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		w.Header().Set("Access-Control-Allow-Origin", app.BaseURL)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	})
}
