package entities

import "github.com/golang/glog"

// Person ...
type Person struct {
	ID         *ID
	Email      *Email
	Name       string
	Adoptables []*Adoptable
}

// CreatePersonWithNameAndEmail ...
func CreatePersonWithNameAndEmail(name string, email string) *Person {
	return &Person{
		ID:         generateID(),
		Email:      BuildEmail(email),
		Name:       name,
		Adoptables: make([]*Adoptable, 0),
	}
}

// AdoptAdoptable ...
func (p *Person) AdoptAdoptable(adoptable *Adoptable) error {
	glog.Info(p)
	glog.Info(adoptable)
	p.Adoptables = append(p.Adoptables, adoptable)
	return nil
}

// HasAdoptedAdoptable ...
func (p *Person) HasAdoptedAdoptable(adoptable *Adoptable) bool {
	for _, item := range p.Adoptables {
		if item.ID.value == adoptable.ID.value {
			return true
		}
	}

	return false
}
