package entities

import "fmt"

// Adoption ...
type Adoption struct {
	People []Person
}

// BuildAdoption ...
func BuildAdoption() *Adoption {
	return &Adoption{
		People: make([]Person, 0),
	}
}

// IncludePerson ...
func (a *Adoption) IncludePerson(person Person) error {
	a.People = append(a.People, person)
	return nil
}

// CalculateForTool ...
func (a *Adoption) CalculateForTool(tool Tool) (int, error) {
	for i := 0; i < len(a.People); i++ {
		fmt.Println(a.People[i].Name)
	}
	return 0, nil
}
