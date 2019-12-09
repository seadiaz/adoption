package entities

// Adoption ...
type Adoption struct {
	People    []*Person
	Teams     []*Team
	Adoptable Adoptable
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

// IncludeTeam ...
func (a *Adoption) IncludeTeam(team *Team) error {
	for i, item := range team.People {
		team.People[i] = a.findPersonByEmail(item.Email)
	}
	a.Teams = append(a.Teams, team)
	return nil
}

func (a *Adoption) findPersonByEmail(email *Email) *Person {
	for _, item := range a.People {
		if item.Email.IsEqual(email) {
			return item
		}
	}

	return nil
}

// CalculateForAdoptable ...
func (a *Adoption) CalculateForAdoptable(adoptable *Adoptable) int {
	total := len(a.People)
	if total == 0 {
		return 0
	}

	counter := 0
	for _, person := range a.People {
		if person.HasAdoptedAdoptable(adoptable) {
			counter++
		}
	}

	return 100 * counter / total
}

// CalculateTeamForAdoptable ...
func (a *Adoption) CalculateTeamForAdoptable(adoptable *Adoptable) int {
	total := len(a.Teams)
	if total == 0 {
		return 0
	}

	counter := 0
	for _, team := range a.Teams {
		if team.HasTeamAdoptedAdoptable(adoptable) {
			counter++
		}
	}

	return 100 * counter / total
}

// FilterAdoptersForAdoptable ...
func (a *Adoption) FilterAdoptersForAdoptable(adoptable *Adoptable) []*Person {
	output := make([]*Person, 0, 0)
	if len(a.People) == 0 {
		return output
	}

	for _, person := range a.People {
		if person.HasAdoptedAdoptable(adoptable) {
			output = append(output, person)
		}
	}
	return output
}

// FilterAbsenteesForAdoptable ...
func (a *Adoption) FilterAbsenteesForAdoptable(adoptable *Adoptable) []*Person {
	output := make([]*Person, 0, 0)
	if len(a.People) == 0 {
		return output
	}

	for _, item := range a.People {
		if !item.HasAdoptedAdoptable(adoptable) {
			output = append(output, item)
		}
	}
	return output
}

// FilterTeamAdoptersForAdoptable ...
func (a *Adoption) FilterTeamAdoptersForAdoptable(adoptable *Adoptable) []*Team {
	output := make([]*Team, 0, 0)
	if len(a.Teams) == 0 {
		return output
	}

	for _, item := range a.Teams {
		if item.HasTeamAdoptedAdoptable(adoptable) {
			output = append(output, item)
		}
	}
	return output
}

// FilterTeamAbsenteesForAdoptable ...
func (a *Adoption) FilterTeamAbsenteesForAdoptable(adoptable *Adoptable) []*Team {
	output := make([]*Team, 0, 0)
	if len(a.Teams) == 0 {
		return output
	}

	for _, item := range a.Teams {
		if !item.HasTeamAdoptedAdoptable(adoptable) {
			output = append(output, item)
		}
	}
	return output
}
