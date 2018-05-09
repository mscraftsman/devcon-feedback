package session

import "time"

// Session represents a conference session
type Session struct {
	ID              *int64     `json:"id"`
	Title           *string    `json:"title"`
	Description     *string    `json:"description"`
	Level           *string    `json:"level"`
	Language        *string    `json:"language"`
	Format          *string    `json:"format"`
	Room            *string    `json:"room"`
	Speakers        *string    `json:"speakers"`
	RatingsCount    *int       `json:"ratings_count"`
	Score           *int       `json:"score"`
	ReactionSummary *string    `json:"reaction_summary"`
	StartAt         *time.Time `json:"start_at"`
	EndAt           *time.Time `json:"end_at"`
	Status          *bool      `json:"status"`
}

// New returns an instance of Session
func New() *Session {
	return &Session{
		ID:              new(int64),
		Title:           new(string),
		Description:     new(string),
		Level:           new(string),
		Language:        new(string),
		Format:          new(string),
		Room:            new(string),
		Speakers:        new(string),
		RatingsCount:    new(int),
		Score:           new(int),
		ReactionSummary: new(string),
		StartAt:         new(time.Time),
		EndAt:           new(time.Time),
		Status:          new(bool),
	}
}
