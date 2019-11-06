package entities

import uuid "github.com/satori/go.uuid"

// Person ...
type Person struct {
	ID    string
	Email string
	Name  string
	Tools []*Tool
}

// CreatePersonWithNameAndEmail ...
func CreatePersonWithNameAndEmail(name string, email string) *Person {
	return &Person{
		ID:    uuid.NewV4().String(),
		Email: email,
		Name:  name,
		Tools: make([]*Tool, 0),
	}
}

// AdoptTool ...
func (p *Person) AdoptTool(tool *Tool) error {
	p.Tools = append(p.Tools, tool)
	return nil
}

// HasAdoptedTool ...
func (p *Person) HasAdoptedTool(tool *Tool) bool {
	for _, item := range p.Tools {
		if item.ID == tool.ID {
			return true
		}
	}

	return false
}
