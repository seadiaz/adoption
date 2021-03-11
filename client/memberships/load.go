package memberships

import (
	"github.com/golang/glog"
	"github.com/seadiaz/adoption/client/utils"
)

func load(r *Repository, teamName, filename string) {
	team := r.FindTeamByName(teamName)
	var items []MembershipInput
	utils.ReadCsvFile(filename, &items)
	for _, item := range items {
		if team.Name != item.TeamName {
			glog.Warningf("membership for team %s omitted", item.TeamName)
			continue
		}
		person := r.FindPersonByEmail(item.PersonEmail)
		r.AddMemberToTeam(person, team)
	}
}
