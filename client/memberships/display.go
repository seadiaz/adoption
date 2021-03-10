package memberships

import (
	"fmt"

	"github.com/seadiaz/adoption/client/global"
	"github.com/seadiaz/adoption/client/people"
	"github.com/seadiaz/adoption/client/teams"

	tm "github.com/buger/goterm"
)

func display(c *global.CommandHandler, teamName string) {
	team := teams.FindTeamByName(c.BaseURL+teams.Path, c.APIKey, teamName)
	people := teams.GetMembersByTeam(c.BaseURL+teams.Path+"/"+team.ID+people.Path, c.APIKey)
	output := prepareToDisplay(team, people)
	print(output)
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

func print(items []*MembershipOutput) {
	table := tm.NewTable(0, 10, 5, ' ', 0)
	fmt.Fprintf(table, "Team\tPerson\n")
	for _, v := range items {
		fmt.Fprintf(table, "%s\t%s\n", v.TeamName, v.PersonName)
	}
	tm.Println(table)
	tm.Flush()
}

func filterPersonByEmail(teams []*people.Person, email string) *people.Person {
	for _, v := range teams {
		if v.Email == email {
			return v
		}
	}

	return nil
}
