package entities

// Email ...
type Email struct {
	value string
}

func (id *Email) String() string {
	return id.value
}
