package entities

// Adoption ...
type Adoption struct {
	People []*Person
	Tool   Tool
}

// CreateAdoption ...
func CreateAdoption() *Adoption {
	return &Adoption{
		People: make([]*Person, 0),
	}
}

// IncludePerson ...
func (a *Adoption) IncludePerson(person *Person) error {
	a.People = append(a.People, person)
	return nil
}

// CalculateForTool ...
func (a *Adoption) CalculateForTool(tool *Tool) int {
	total := len(a.People)
	if total == 0 {
		return 0
	}

	counter := 0
	for _, person := range a.People {
		if person.HasAdoptedTool(tool) {
			counter++
		}
	}
	return 100 * counter / total
}
