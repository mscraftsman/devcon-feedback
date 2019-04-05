package sessionize

//import (
//	"net/http"
//
//	"github.com/gorilla/mux"
//	"github.com/mscraftsman/devcon-feedback/util"
//)
//
//// GetSpeaker is an API endpoint to retrieve a specific Route
//func GetSpeaker(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//
//	id, ok := vars["id"]
//	if !ok {
//		util.JSONOutputError(w, http.StatusBadRequest, "id not specified")
//		return
//	}
//
//	_, ok = _speakers[id]
//	if !ok {
//		util.JSONOutputError(w, http.StatusNotFound, "speaker not found")
//		return
//	}
//
//	util.JSONOutputResponse(w, _speakers[id])
//}
//
//// GetSpeakers is an API endpoint to retrieve the list of Routes
//func GetSpeakers(w http.ResponseWriter, r *http.Request) {
//	var speakers []Speaker
//	for _, speaker := range _speakers {
//		speakers = append(speakers, speaker)
//	}
//	util.JSONOutputResponse(w, speakers)
//}
//
//// GetSession is an API endpoint to retrieve a specific Session
//func GetSession(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//
//	id, ok := vars["id"]
//	if !ok {
//		util.JSONOutputError(w, http.StatusBadRequest, "id not specified")
//		return
//	}
//
//	_, ok = _sessions[id]
//	if !ok {
//		util.JSONOutputError(w, http.StatusNotFound, "session not found")
//		return
//	}
//
//	util.JSONOutputResponse(w, _sessions[id])
//}
//
//// GetSessions is an API endpoint to retrieve the list of Sessions
//func GetSessions(w http.ResponseWriter, r *http.Request) {
//	var sessions []Session
//	for _, session := range _sessions {
//		sessions = append(sessions, session)
//	}
//	util.JSONOutputResponse(w, sessions)
//}
//
//// GetRoom is an API endpoint to retrieve a specific Room
//func GetRoom(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//
//	id, ok := vars["id"]
//	if !ok {
//		util.JSONOutputError(w, http.StatusBadRequest, "id not specified")
//		return
//	}
//
//	_, ok = _rooms[id]
//	if !ok {
//		util.JSONOutputError(w, http.StatusNotFound, "room not found")
//		return
//	}
//
//	util.JSONOutputResponse(w, _rooms[id])
//}
//
//// GetRooms is an API endpoint to retrieve the list of Rooms
//func GetRooms(w http.ResponseWriter, r *http.Request) {
//	var rooms []Room
//	for _, room := range _rooms {
//		rooms = append(rooms, room)
//	}
//	util.JSONOutputResponse(w, rooms)
//}
