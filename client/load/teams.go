package load

// import "github.com/golang/glog"

// const teamsPath = "/teams"

// // Team ...
// type Team struct {
// 	ID   string `json:"id,omitempty"`
// 	Name string `json:"name"`
// }

// // LoadTeams ...
// func (c *client) LoadTeams() {
// 	rawData := readCsvFile(c.filename)
// 	parsedData := mapArrayToTeams(rawData)
// 	c.postTeams(parsedData)
// }

// func mapArrayToTeams(array [][]string) []*Team {
// 	output := make([]*Team, 0, 0)
// 	for _, item := range array {
// 		output = append(output, &Team{
// 			Name: item[0],
// 		})
// 	}
// 	return output
// }

// func (c *client) postTeams(teams []*Team) {
// 	channel := make(chan string)
// 	for _, item := range teams {
// 		go c.postTeam(item, channel)
// 	}

// 	receiveResponses(channel, len(teams))
// }

// func (c *client) postTeam(team *Team, channel chan string) {
// 	doPostRequest(team, c.url+teamsPath, c.apiKey)
// 	channel <- team.Name
// }

// func findTeamByName(teams []*Team, name string) *Team {
// 	for _, item := range teams {
// 		if item.Name == name {
// 			return item
// 		}
// 	}

// 	return nil
// }

// func (c *client) getTeams() []*Team {
// 	res, err := doGetRequest(c.url+teamsPath, c.apiKey)
// 	output := make([]*Team, 0, 0)
// 	if err != nil {
// 		glog.Error(err)
// 		return output
// 	}

// 	for _, item := range res {
// 		output = append(output, &Team{
// 			ID:   item["id"].(string),
// 			Name: item["name"].(string),
// 		})
// 	}
// 	return output
// }
