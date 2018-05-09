package visitor

// Visitor represents a conference visitor
type Visitor struct {
	ID        *int64  `json:"id"`
	Name      *string `json:"name"`
	MeetupID  *int    `json:"meetup_id"`
	PhotoLink *string `json:"photo_link"`
}

// New returns an instance of Visitor
func New() *Visitor {
	return &Visitor{
		ID:        new(int64),
		Name:      new(string),
		MeetupID:  new(int),
		PhotoLink: new(string),
	}
}
