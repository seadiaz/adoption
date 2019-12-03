package entities

// Email ...
type Email struct {
	Value string
}

func (id *Email) String() string {
	return id.Value
}
