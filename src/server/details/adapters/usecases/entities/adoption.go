package entities

// Adoption ...
type Adoption struct {
	People []*Person
	Teams  []*Team
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

// IncludeTeam ...
func (a *Adoption) IncludeTeam(team *Team) error {
	for i, item := range team.People {
		team.People[i] = a.findPersonByEmail(item.Email)
	}
	a.Teams = append(a.Teams, team)
	return nil
}

func (a *Adoption) findPersonByEmail(email string) *Person {
	for _, item := range a.People {
		if item.Email == email {
			return item
		}
	}

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

// CalculateTeamForTool ...
func (a *Adoption) CalculateTeamForTool(tool *Tool) int {
	total := len(a.Teams)
	if total == 0 {
		return 0
	}

	counter := 0
	for _, team := range a.Teams {
		if team.HasTeamAdoptedTool(tool) {
			counter++
		}
	}

	return 100 * counter / total
}

// FilterAdoptersForTool ...
func (a *Adoption) FilterAdoptersForTool(tool *Tool) []*Person {
	output := make([]*Person, 0, 0)
	if len(a.People) == 0 {
		return output
	}

	for _, person := range a.People {
		if person.HasAdoptedTool(tool) {
			output = append(output, person)
		}
	}
	return output
}

// FilterAbsenteesForTool ...
func (a *Adoption) FilterAbsenteesForTool(tool *Tool) []*Person {
	output := make([]*Person, 0, 0)
	if len(a.People) == 0 {
		return output
	}

	for _, item := range a.People {
		if !item.HasAdoptedTool(tool) {
			output = append(output, item)
		}
	}
	return output
}

// FilterTeamAdoptersForTool ...
func (a *Adoption) FilterTeamAdoptersForTool(tool *Tool) []*Team {
	output := make([]*Team, 0, 0)
	if len(a.Teams) == 0 {
		return output
	}

	for _, item := range a.Teams {
		if item.HasTeamAdoptedTool(tool) {
			output = append(output, item)
		}
	}
	return output
}

// FilterTeamAbsenteesForTool ...
func (a *Adoption) FilterTeamAbsenteesForTool(tool *Tool) []*Team {
	output := make([]*Team, 0, 0)
	if len(a.Teams) == 0 {
		return output
	}

	for _, item := range a.Teams {
		if !item.HasTeamAdoptedTool(tool) {
			output = append(output, item)
		}
	}
	return output
}
