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
	var items []people.Person
	utils.ReadCsvFile(filename, &items)
	channel := make(chan string)
	for _, item := range items {
		p := people.FindPersonByEmail(c.BaseURL+people.Path, c.APIKey, item.Email)
		go postMembership(c.BaseURL+teams.Path+"/"+team.ID+people.Path, c.APIKey, *p, channel)
	}

	utils.ReceiveResponses(channel, len(items))
}

func postMembership(url, apiKey string, person people.Person, channel chan string) {
	utils.DoPostRequest(url, apiKey, person)
	channel <- person.Name
}
