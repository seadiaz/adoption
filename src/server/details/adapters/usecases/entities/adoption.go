package entities

import "github.com/thoas/go-funk"

// Adoption ...
type Adoption struct {
	People      []*Person
	TeamDetails []*AdoptionTeamDetail
	Adoptable   Adoptable
}

// AdoptionTeamDetail ...
type AdoptionTeamDetail struct {
	Team  *Team
	Level int
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
	a.TeamDetails = append(a.TeamDetails, &AdoptionTeamDetail{Team: team, Level: 0})
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

// CalculateOverPeopleForAdoptable ...
func (a *Adoption) CalculateOverPeopleForAdoptable(adoptable *Adoptable) int {
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

// CalculateOverTeamForAdoptable ...
func (a *Adoption) CalculateOverTeamForAdoptable(adoptable *Adoptable) int {
	total := len(a.TeamDetails)
	if total == 0 {
		return 0
	}

	counter := 0
	for _, team := range a.TeamDetails {
		if team.Team.HasTeamAdoptedAdoptable(adoptable) {
			counter++
		}
	}

	return 100 * counter / total
}

// FilterPeopleAdoptersForAdoptable ...
func (a *Adoption) FilterPeopleAdoptersForAdoptable(adoptable *Adoptable) []*Person {
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

// FilterPeopleAbsenteesForAdoptable ...
func (a *Adoption) FilterPeopleAbsenteesForAdoptable(adoptable *Adoptable) []*Person {
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
func (a *Adoption) FilterTeamAdoptersForAdoptable(adoptable *Adoptable) []*AdoptionTeamDetail {
	output := make([]*AdoptionTeamDetail, 0, 0)
	if len(a.TeamDetails) == 0 {
		return output
	}

	for _, item := range a.TeamDetails {
		if item.Team.HasTeamAdoptedAdoptable(adoptable) {
			item.Level = a.calculateOverPeopleForTeamAndAdoptable(item.Team, adoptable)
			output = append(output, item)
		}
	}
	return output
}

func (a *Adoption) calculateOverPeopleForTeamAndAdoptable(team *Team, adoptable *Adoptable) int {
	adoption := CreateAdoption()
	for _, item := range team.People {
		completePerson := funk.Find(a.People, func(p *Person) bool { return p.ID == item.ID })
		adoption.IncludePerson(completePerson.(*Person))
	}

	return adoption.CalculateOverPeopleForAdoptable(adoptable)
}

// FilterTeamAbsenteesForAdoptable ...
func (a *Adoption) FilterTeamAbsenteesForAdoptable(adoptable *Adoptable) []*AdoptionTeamDetail {
	output := make([]*AdoptionTeamDetail, 0, 0)
	if len(a.TeamDetails) == 0 {
		return output
	}

	for _, item := range a.TeamDetails {
		if !item.Team.HasTeamAdoptedAdoptable(adoptable) {
			output = append(output, item)
		}
	}
	return output
}
