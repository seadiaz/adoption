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
	print(people)
}

func print(people []*people.Person) {
	table := tm.NewTable(0, 10, 5, ' ', 0)
	fmt.Fprintf(table, "ID\tName\tEmail\n")
	for _, p := range people {
		fmt.Fprintf(table, "%s\t%s\t%s\n", p.ID, p.Name, p.Email)
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
