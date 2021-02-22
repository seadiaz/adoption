package teams

import (
	"github.com/seadiaz/adoption/client/global"
	"github.com/seadiaz/adoption/client/utils"
)

func load(c *global.CommandHandler, filename string) {
	var teams []Team
	utils.ReadCsvFile(filename, &teams)
	channel := make(chan string)
	for _, item := range teams {
		go postTeam(c.BaseURL+Path, c.APIKey, item, channel)
	}

	utils.ReceiveResponses(channel, len(teams))
}

func postTeam(url, apiKey string, team Team, channel chan string) {
	utils.DoPostRequest(url, apiKey, team)
	channel <- team.Name
}
