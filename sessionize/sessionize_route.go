package sessionize

import (
	"net/http"
)

//SessionsCache returns a cached version of sessions information from sessionize
func SessionsCache(w http.ResponseWriter, r *http.Request) {
	if rawSessions == nil {
		w.WriteHeader(http.StatusServiceUnavailable)
	} else {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(rawSessions)
	}
}
