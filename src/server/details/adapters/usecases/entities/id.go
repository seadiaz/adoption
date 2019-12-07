package entities

import uuid "github.com/satori/go.uuid"

// ID ...
type ID struct {
	value string
}

func newID() *ID {
	return &ID{
		value: uuid.NewV4().String(),
	}
}

// RecoverID ...
func RecoverID(value string) *ID {
	return &ID{
		value: value,
	}
}

func (id *ID) String() string {
	return id.value
}
