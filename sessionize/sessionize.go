package sessionize

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	_sessions map[string]Session
	_speakers map[string]Speaker
	_rooms    map[string]Room
	_timeZone *time.Location
)

// Time represents json time without timezone
type Time time.Time

func init() {
	var e error
	_timeZone, e = time.LoadLocation("Indian/Mauritius")
	if e != nil {
		fmt.Println(e)
	}
}

// UnmarshalJSON decodes from json string
func (m *Time) UnmarshalJSON(p []byte) error {
	t, err := time.Parse("2006-01-02T15:04:05", strings.Trim(string(p), `"`))

	if err != nil {
		return err
	}

	*m = Time(t)

	return nil
}

// APIResponse represents api response from sessionize
type APIResponse struct {
	Sessions []Session `json:"sessions"`
	Speakers []Speaker `json:"speakers"`
	Rooms    []Room    `json:"rooms"`
}

// Session represents a session on sessionize
type Session struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	StartsAt    Time     `json:"startsAt"`
	EndsAt      Time     `json:"endsAt"`
	RoomID      int      `json:"roomID"`
	Speakers    []string `json:"speakers"`
}

// Speaker represents a speaker profile on sessionize
type Speaker struct {
	ID             string `json:"id"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	Bio            string `json:"bio"`
	TagLine        string `json:"tagLine"`
	ProfilePicture string `json:"profilePicture"`
	Sessions       []int  `json:"sessions"`
}

// Room represents session venue on sessionize
type Room struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// ReadFromAPI loads data from sessionize API
func ReadFromAPI() (*APIResponse, error) {
	var err error

	var response APIResponse

	client := &http.Client{Timeout: time.Second * 30}

	rs, err := client.Get("https://sessionize.com/api/v2/351ijy5v/view/All")

	if err != nil {
		return nil, err
	}

	resp, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// IsVotableSession check if given session id is valid
// returns valid, open
func IsVotableSession(id string) (bool, bool) {
	sess, ok := _sessions[id]
	return ok, time.Now().In(_timeZone).After(time.Time(sess.StartsAt))
}

// Sync keeps Sessionize data up to date
func Sync() error {
	response, err := ReadFromAPI()

	if err != nil {
		return err
	}

	_sessions = make(map[string]Session)
	for _, session := range response.Sessions {
		_sessions[session.ID] = session
	}

	_speakers = make(map[string]Speaker)
	for _, speaker := range response.Speakers {
		_speakers[speaker.ID] = speaker
	}

	_rooms = make(map[string]Room)
	for _, room := range response.Rooms {
		_rooms[strconv.Itoa(room.ID)] = room
	}

	return nil
}

// Get gets a specific session
func Get(sessionID string) Session {
	var names []string
	s, ok := _sessions[sessionID]

	if ok {
		for _, i := range s.Speakers {
			if speaker, ok := _speakers[i]; ok {
				names = append(names, speaker.FirstName+" "+speaker.LastName)
			}
		}

		s.Speakers = names

		return s
	} else {
		return Session{}
	}
}
