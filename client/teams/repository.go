package teams

import (
	"github.com/mitchellh/mapstructure"
	"github.com/seadiaz/adoption/client/people"
	"github.com/seadiaz/adoption/client/utils"
)

// GetTeams ...
func GetTeams(url, apiKey string) []*Team {
	res, _ := utils.DoGetRequest(url, apiKey)
	output := make([]*Team, 0, 0)
	for _, item := range res {
		var person Team
		mapstructure.Decode(item, &person)
		output = append(output, &person)
	}
	return output
}

// FindTeamByName ...
func FindTeamByName(url, apiKey, teamName string) *Team {
	teams := GetTeams(url, apiKey)
	team := filterTeamByName(teams, teamName)
	return team
}

func filterTeamByName(teams []*Team, name string) *Team {
	for _, v := range teams {
		if v.Name == name {
			return v
		}
	}

	return nil
}

// GetMembersByTeam ...
func GetMembersByTeam(url, apiKey string) []*people.Person {
	res, _ := utils.DoGetRequest(url, apiKey)
	output := make([]*people.Person, 0, 0)
	for _, item := range res {
		var person people.Person
		mapstructure.Decode(item, &person)
		output = append(output, &person)
	}
	return output
}
