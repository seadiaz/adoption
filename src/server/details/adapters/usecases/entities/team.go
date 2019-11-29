package entities

// Team ...
type Team struct {
	ID     *ID
	Name   string
	People []*Person
}

// CreateTeamWithName ...
func CreateTeamWithName(name string) *Team {
	return &Team{
		ID:   newID(),
		Name: name,
	}
}

// AddPerson ...
func (t *Team) AddPerson(person *Person) {
	person.Tools = nil
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

// HasTeamAdoptedTool ...
func (t *Team) HasTeamAdoptedTool(tool *Tool) bool {
	for _, item := range t.People {
		if item.HasAdoptedTool(tool) {
			return true
		}
	}

	return false
}
