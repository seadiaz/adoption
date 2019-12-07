package entities

import uuid "github.com/satori/go.uuid"

// ID ...
type ID struct {
	value string
}

func generateID() *ID {
	return &ID{
		value: uuid.NewV4().String(),
	}
}

// BuildID ...
func BuildID(value string) *ID {
	return &ID{
		value: value,
	}
}

func (id *ID) String() string {
	return id.value
}
