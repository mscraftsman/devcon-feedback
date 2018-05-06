package visitor

// Visitor represents a conference visitor
type Visitor struct {
	ID *int64 `json:"id"`

	Name *string `json:"code"`

	Name *string `json:"name"`

	Affiliation *string `json:"affiliation"`

	Email *string `json:"email"`

	Level *String `json:"level"`
}

// New returns an instance of Visitor
func New() *Visitor {
	entity := new(Visitor)
	entity.ID = new(int64)

	entity.Name = new(string)

	entity.Name = new(string)

	entity.Affiliation = new(string)

	entity.Email = new(string)

	entity.Level = new(String)

	return entity
}
