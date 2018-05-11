// generated by gocipe ce0bfb5c18c811c9714539fbe97b8d7b2dff895a3a2978a083ddc4f23c0a35ab; DO NOT EDIT

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