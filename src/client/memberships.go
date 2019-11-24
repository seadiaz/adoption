package client

// Membership ...
type Membership struct {
	PersonEmail string
	TeamName    string
	Team        *Team
}

// LoadMemberships ...
func (c *client) LoadMemberships() {
	rawData := readCsvFile(c.filename)
	parsedData := mapArrayToMemberships(rawData)
	teams := c.getTeams()
	parsedData = fulfillMembershipTeamFromTeamName(parsedData, teams)
	c.postMemberships(parsedData)
}

func mapArrayToMemberships(array [][]string) []*Membership {
	output := make([]*Membership, 0, 0)
	for _, item := range array {
		output = append(output, &Membership{
			TeamName:    item[0],
			PersonEmail: item[1],
		})
	}
	return output
}

func fulfillMembershipTeamFromTeamName(Memberships []*Membership, teams []*Team) []*Membership {
	output := make([]*Membership, 0, 0)
	for _, item := range Memberships {
		team := findTeamByName(teams, item.TeamName)
		if team != nil {
			item.Team = team
		}
		output = append(output, item)
	}
	return output
}

func (c *client) postMemberships(memberships []*Membership) {
	channel := make(chan string)
	for _, item := range memberships {
		go c.postMembership(item, channel)
	}

	receiveResponses(channel, len(memberships))
}

func (c *client) postMembership(membership *Membership, channel chan string) {
	body := &Person{Email: membership.PersonEmail}
	doPostRequest(body, c.url+teamsPath+"/"+membership.Team.ID+peoplePath, c.apiKey)
	channel <- membership.PersonEmail
}
