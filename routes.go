package main

import (
	"net/http"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/mscraftsman/devcon-feedback/config"
	"github.com/mscraftsman/devcon-feedback/controllers"
	"github.com/mscraftsman/devcon-feedback/meetup"
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
	{"/api/me", http.MethodGet, controllers.Me, nil},
	{"/api/bookmarks/{id}", http.MethodPut, controllers.AddBookmark, nil},
	{"/api/bookmarks", http.MethodGet, controllers.ListBookmarks, nil},
	{"/api/bookmarks/{id}", http.MethodDelete, controllers.RemoveBookmark, nil},
	{"/api/feedbacks", http.MethodPost, controllers.AddFeedback, nil},
	{"/api/feedbacks", http.MethodGet, controllers.ListAllFeedback, nil},
	{"/api/feedbacks/me", http.MethodGet, controllers.ListOwnFeedback, nil},
	{"/api/ratings", http.MethodGet, controllers.ListRatings, nil},
	{"/api/sessions", http.MethodGet, controllers.ListSessions, nil},
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
		return strings.HasPrefix(orig, config.FrontURL)
	})

	return handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods, allowCredentials)(r)
}
