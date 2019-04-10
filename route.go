package main

import (
	"net/http"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/mscraftsman/devcon-feedback/config"
	"github.com/mscraftsman/devcon-feedback/controller"
	"github.com/mscraftsman/devcon-feedback/meetup"
	"github.com/mscraftsman/devcon-feedback/sessionize"
)

type route struct {
	Path        string
	Method      string
	HandlerFunc http.HandlerFunc
	Handler     http.Handler
}

var routes = []route{
	{"/login", http.MethodGet, meetup.Login, nil},
	{"/logout", http.MethodGet, meetup.Logout, nil},
	{"/meetup", http.MethodGet, meetup.LoginCallback, nil},
	{"/api/me", http.MethodGet, controller.Me, nil},
	{"/api/bookmarks/{id}", http.MethodPut, controller.AddBookmark, nil},
	{"/api/bookmarks", http.MethodGet, controller.ListBookmarks, nil},
	{"/api/bookmarks/{id}", http.MethodDelete, controller.RemoveBookmark, nil},
	{"/api/feedbacks", http.MethodPost, controller.AddFeedback, nil},
	{"/api/feedbacks/me", http.MethodGet, controller.ListOwnFeedback, nil},
	{"/api/sessions", http.MethodGet, sessionize.SessionsCache, nil},
	// {"/api/leaderboards", http.MethodGet, controller.Leaderboards, nil},
	// {"/api/feedbacks", http.MethodGet, controller.ListAllFeedback, nil},
	// {"/api/ratings", http.MethodGet, controller.ListRatings, nil},
	// {"/api/sessions", http.MethodGet, controller.ListSessions, nil},
}

func initRouter() http.Handler {
	r := mux.NewRouter()
	for i := range routes {
		t := r.Path(routes[i].Path).Methods(http.MethodOptions, routes[i].Method)
		if routes[i].HandlerFunc != nil {
			t.HandlerFunc(routes[i].HandlerFunc)
		} else if routes[i].Handler != nil {
			t.Handler(routes[i].Handler)
		}
	}

	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Cookie", "Origin", "Content-type"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})
	allowCredentials := handlers.AllowCredentials()
	allowedOrigins := handlers.AllowedOriginValidator(func(orig string) bool {
		if strings.HasPrefix(orig, config.FrontURL) {
			return true
		}

		for _, url := range config.AdditionalCORS {
			if strings.HasPrefix(orig, url) {
				return true
			}
		}

		return false
	})

	return handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods, allowCredentials)(r)
}
