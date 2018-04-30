package session

// Session represents a conference session
type Session struct {
	ID        *int64  `json:"id"`
	
	Id *int64 `json:"id"`
	
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