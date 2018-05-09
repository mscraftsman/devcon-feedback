package visitor

// Visitor represents a conference visitor
type Visitor struct {
	ID *int64 `json:"id"`

	Name *string `json:"name"`

	MeetupID *int `json:"meetup_id"`

	PhotoLink *string `json:"photo_link"`
}

// New returns an instance of Visitor
func New() *Visitor {
	entity := new(Visitor)
	entity.ID = new(int64)

	entity.Name = new(string)

	entity.MeetupID = new(int)

	entity.PhotoLink = new(string)

	return entity
}
