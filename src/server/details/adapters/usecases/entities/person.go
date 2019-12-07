package entities

import "github.com/golang/glog"

// Person ...
type Person struct {
	ID    *ID
	Email *Email
	Name  string
	Tools []*Tool
}

// CreatePersonWithNameAndEmail ...
func CreatePersonWithNameAndEmail(name string, email string) *Person {
	return &Person{
		ID:    generateID(),
		Email: BuildEmail(email),
		Name:  name,
		Tools: make([]*Tool, 0),
	}
}

// AdoptTool ...
func (p *Person) AdoptTool(tool *Tool) error {
	glog.Info(p)
	glog.Info(tool)
	p.Tools = append(p.Tools, tool)
	return nil
}

// HasAdoptedTool ...
func (p *Person) HasAdoptedTool(tool *Tool) bool {
	for _, item := range p.Tools {
		if item.ID.value == tool.ID.value {
			return true
		}
	}

	return false
}
