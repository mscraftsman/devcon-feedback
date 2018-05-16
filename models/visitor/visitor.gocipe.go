// generated by gocipe 89440e4dbce197f5b46d60d778de3e709012a9e0bf60a5510bf683dbf5304465; DO NOT EDIT

package visitor

// Visitor represents a conference visitor
type Visitor struct {
	ID        *string `json:"id"`
	Name      *string `json:"name"`
	PhotoLink *string `json:"photo_link"`
	Status    *bool   `json:"status"`
}

// New returns an instance of Visitor
func New() *Visitor {
	return &Visitor{
		ID:        new(string),
		Name:      new(string),
		PhotoLink: new(string),
		Status:    new(bool),
	}
}
