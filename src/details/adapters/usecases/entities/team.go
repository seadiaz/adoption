package entities

import uuid "github.com/satori/go.uuid"

// Team ...
type Team struct {
	ID     string
	Name   string
	People []*Person
}

// CreateTeamWithName ...
func CreateTeamWithName(name string) *Team {
	return &Team{
		ID:   uuid.Must(uuid.NewV4()).String(),
		Name: name,
	}
}

// AddPerson ...
func (t *Team) AddPerson(person *Person) {
	t.People = append(t.People, person)
}

// RemovePerson ...
func (t *Team) RemovePerson(person *Person) {
	for i, item := range t.People {
		if item.Email == person.Email {
			t.People = append(t.People[:i], t.People[i+1:]...)
			return
		}
	}
}
