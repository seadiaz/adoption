package client

import "fmt"

// Membership ...
type Membership struct {
	PersonEmail string
	TeamName    string
	Team        *Team
	Person      *Person
}

// LoadMemberships ...
func (c *client) LoadMemberships() {
	rawData := readCsvFile(c.filename)
	parsedData := mapArrayToMemberships(rawData)
	teams := c.getTeams()
	people := c.getPeople()
	parsedData = fulfillMembershipTeamFromTeamName(parsedData, teams, people)
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

func fulfillMembershipTeamFromTeamName(Memberships []*Membership, teams []*Team, people []*Person) []*Membership {
	output := make([]*Membership, 0, 0)
	for _, item := range Memberships {
		team := findTeamByName(teams, item.TeamName)
		person := findPersonByEmail(people, item.PersonEmail)
		if team != nil {
			item.Team = team
			item.Person = person
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
	body := &Person{ID: membership.Person.ID}
	err := doPostRequest(body, c.url+teamsPath+"/"+membership.Team.ID+peoplePath, c.apiKey)
	if err != nil {
		channel <- fmt.Sprintf("fail adding %s to %s: %s", membership.PersonEmail, membership.TeamName, err.Error())
	} else {
		channel <- fmt.Sprintf("%s added to %s", membership.PersonEmail, membership.TeamName)
	}
}
