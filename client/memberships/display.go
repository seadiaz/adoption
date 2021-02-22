package memberships

import (
	"github.com/seadiaz/adoption/client/global"
	"github.com/seadiaz/adoption/client/people"
	"github.com/seadiaz/adoption/client/teams"
)

func display(c *global.CommandHandler) {
	people := people.GetPeople(c.BaseURL+people.Path, c.APIKey)
	teams := teams.GetTeams(c.BaseURL+teams.Path, c.APIKey)

}

func fulfillMembershipTeamFromTeamName(Memberships []*Membership, teams []*teams.Team, people []*people.Person) []*Membership {
	output := make([]*Membership, 0, 0)
	for _, item := range Memberships {
		team := filterTeamByName(teams, item.TeamName)
		person := filterPersonByEmail(people, item.PersonEmail)
		if team != nil {
			item.Team = team
			item.Person = person
		}
		output = append(output, item)
	}
	return output
}

func filterTeamByName(teams []*teams.Team, name string) *teams.Team {
	for _, v := range teams {
		if v.Name == name {
			return v
		}
	}

	return nil
}

func filterPersonByEmail(teams []*people.Person, email string) *people.Person {
	for _, v := range teams {
		if v.Email == email {
			return v
		}
	}

	return nil
}
