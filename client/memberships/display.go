package memberships

import (
	"fmt"

	"github.com/seadiaz/adoption/client/people"

	tm "github.com/buger/goterm"
)

func display(r *Repository, teamName string) {
	members := r.GetMembersByTeam(teamName)
	print(members)
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
