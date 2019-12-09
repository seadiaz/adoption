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
		ID:   generateID(),
		Name: name,
	}
}

// AddPerson ...
func (t *Team) AddPerson(person *Person) {
	person.Adoptables = nil
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

// HasTeamAdoptedAdoptable ...
func (t *Team) HasTeamAdoptedAdoptable(adoptable *Adoptable) bool {
	for _, item := range t.People {
		if item.HasAdoptedAdoptable(adoptable) {
			return true
		}
	}

	return false
}
