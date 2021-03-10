package memberships

import (
	"github.com/golang/glog"
	"github.com/seadiaz/adoption/client/global"
	"github.com/seadiaz/adoption/client/people"
	"github.com/seadiaz/adoption/client/teams"
	"github.com/seadiaz/adoption/client/utils"
)

func load(c *global.CommandHandler, teamName, filename string) {
	team := teams.FindTeamByName(c.BaseURL+teams.Path, c.APIKey, teamName)
	glog.Info(team)
	var items []MembershipInput
	utils.ReadCsvFile(filename, &items)
	channel := make(chan string)
	for _, item := range items {
		person := people.FindPersonByEmail(c.BaseURL+people.Path, c.APIKey, item.PersonEmail)
		go postMembership(c.BaseURL+teams.Path+"/"+team.ID+people.Path, c.APIKey, *person, channel)
	}

	utils.ReceiveResponses(channel, len(items))
}

func postMembership(url, apiKey string, person people.Person, channel chan string) {
	utils.DoPostRequest(url, apiKey, person)
	channel <- person.Name
}
