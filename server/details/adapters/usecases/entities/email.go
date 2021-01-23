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

// IsValid ...
func (v *Email) IsValid() bool {
	return v.value != ""
}

// IsEqual ...
func (v *Email) IsEqual(other *Email) bool {
	return v.value == other.value
}

func (v *Email) String() string {
	return v.value
}
