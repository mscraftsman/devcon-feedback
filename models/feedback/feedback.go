package feedback

import "time"

// Feedback represents feedback from a visitor on a session
type Feedback struct {
	ID *int64 `json:"id"`

	VisitorId *int64 `json:"visitor_id"`

	Reaction1 *string `json:"reaction_1"`

	Reaction2 *string `json:"reaction_2"`

	Reaction3 *string `json:"reaction_3"`

	Reaction4 *string `json:"reaction_4"`

	CreatedAt *time.Time `json:"created_at"`
}

// New returns an instance of Feedback
func New() *Feedback {
	entity := new(Feedback)
	entity.ID = new(int64)

	entity.VisitorId = new(int64)

	entity.Reaction1 = new(string)

	entity.Reaction2 = new(string)

	entity.Reaction3 = new(string)

	entity.Reaction4 = new(string)

	entity.CreatedAt = new(time.Time)

	return entity
}
