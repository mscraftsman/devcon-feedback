package session

import "time"

// Session represents a conference session
type Session struct {
	ID *int64 `json:"id"`

	Title *string `json:"title"`

	Description *string `json:"description"`

	Level *string `json:"level"`

	Language *string `json:"language"`

	Format *string `json:"format"`

	Room *string `json:"room"`

	Speakers *string `json:"speakers"`

	RatingsCount *int `json:"ratings_count"`

	Score *int `json:"score"`

	ReactionSummary *string `json:"reaction_summary"`

	StartAt *time.Time `json:"start_at"`

	EndAt *time.Time `json:"end_at"`

	Status *bool `json:"status"`
}

// New returns an instance of Session
func New() *Session {
	entity := new(Session)
	entity.ID = new(int64)

	entity.Title = new(string)

	entity.Description = new(string)

	entity.Level = new(string)

	entity.Language = new(string)

	entity.Format = new(string)

	entity.Room = new(string)

	entity.Speakers = new(string)

	entity.RatingsCount = new(int)

	entity.Score = new(int)

	entity.ReactionSummary = new(string)

	entity.StartAt = new(time.Time)

	entity.EndAt = new(time.Time)

	entity.Status = new(bool)

	return entity
}
