package memberships

import (
	"github.com/golang/glog"
	"github.com/seadiaz/adoption/client/people"
	"github.com/seadiaz/adoption/client/teams"
)

// Membership ...
type Membership struct {
	Person *people.Person
	Team   *teams.Team
}

// Repository ...
type Repository struct {
	PeopleRepository *people.Repository
	TeamsRepository  *teams.Repository
}

// CreateRepository ...
func CreateRepository(peopleRepository *people.Repository, teamsRepository *teams.Repository) *Repository {
	return &Repository{
		peopleRepository,
		teamsRepository,
	}
}

// GetMembersByTeam ...
func (r *Repository) GetMembersByTeam(teamName string) []*MembershipOutput {
	team := r.TeamsRepository.FindTeamByName(teamName)
	people := r.TeamsRepository.GetMembersByTeam(team)
	members := prepareToDisplay(team, people)
	return members
}

// AddMemberToTeam ...
func (r *Repository) AddMemberToTeam(p *people.Person, t *teams.Team) []*Membership {
	err := r.TeamsRepository.AddMemberToTeam(p, t)
	if err != nil {
		glog.Errorf("add member %s to team %s failed", p.Name, t.Name)
	}

	return nil
}

// FindTeamByName ...
func (r *Repository) FindTeamByName(teamName string) *teams.Team {
	team := r.TeamsRepository.FindTeamByName(teamName)
	return team
}

// FindPersonByEmail ...
func (r *Repository) FindPersonByEmail(email string) *people.Person {
	person := r.PeopleRepository.FindPersonByEmail(email)
	return person
}

func prepareToDisplay(team *teams.Team, people []*people.Person) []*MembershipOutput {
	output := make([]*MembershipOutput, len(people))
	for i, v := range people {
		output[i] = &MembershipOutput{
			TeamName:   team.Name,
			PersonName: v.Name,
		}
	}
	return output
}
