package cli

const teamsPath = "/teams"

// Team ...
type Team struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
}

// LoadTeams ...
func (c *Client) LoadTeams() {
	rawData := readCsvFile(c.Filename)
	parsedData := mapArrayToTeams(rawData)
	c.postTeams(parsedData)
}

func mapArrayToTeams(array [][]string) []*Team {
	output := make([]*Team, 0, 0)
	for _, item := range array {
		output = append(output, &Team{
			Name: item[0],
		})
	}
	return output
}

func (c *Client) postTeams(teams []*Team) {
	for _, item := range teams {
		c.postTeam(item)
	}
}

func (c *Client) postTeam(team *Team) {
	doPostRequest(team, c.URL+teamsPath, c.APIKey)
}

func findTeamByName(teams []*Team, name string) *Team {
	for _, item := range teams {
		if item.Name == name {
			return item
		}
	}

	return nil
}

func (c *Client) getTeams() []*Team {
	res := doGetRequest(c.URL+teamsPath, c.APIKey)
	output := make([]*Team, 0, 0)
	for _, item := range res {
		output = append(output, &Team{
			ID:   item["id"].(string),
			Name: item["name"].(string),
		})
	}
	return output
}
