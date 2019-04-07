package store

import (
	"time"
)

//Feedback for a session by an attendee
type Feedback struct {
	ID         string    `json:"id"`
	SessionID  string    `json:"session_id"`
	AttendeeID string    `json:"visitor_id"`
	Reaction1  string    `json:"reaction_1"`
	Reaction2  string    `json:"reaction_2"`
	Reaction3  string    `json:"reaction_3"`
	Reaction4  string    `json:"reaction_4"`
	CreatedAt  time.Time `json:"created_at"`
}
