package entities

// Email ...
type Email struct {
	value string
}

// BuildEmail ...
func BuildEmail(value string) *Email {
	return &Email{
		value: value,
	}
}

func (id *Email) String() string {
	return id.value
}
