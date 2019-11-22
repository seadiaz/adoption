package entities

import uuid "github.com/satori/go.uuid"

// Tool ...
type Tool struct {
	ID   string
	Name string
}

// CreateToolWithName ...
func CreateToolWithName(name string) *Tool {
	return &Tool{
		Name: name,
		ID:   uuid.Must(uuid.NewV4()).String(),
	}
}
