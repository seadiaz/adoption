package entities

import uuid "github.com/satori/go.uuid"

// ID ...
type ID struct {
	Value string
}

func newID() *ID {
	return &ID{
		Value: uuid.NewV4().String(),
	}
}

func (id *ID) String() string {
	return id.Value
}
