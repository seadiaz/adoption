package teams

import (
	"github.com/mitchellh/mapstructure"
	"github.com/seadiaz/adoption/client/utils"
)

// GetTeams ...
func GetTeams(url, apiKey string) []*Team {
	res, _ := utils.DoGetRequest(url, apiKey)
	output := make([]*Team, 0, 0)
	for _, item := range res {
		var person Team
		mapstructure.Decode(item, &person)
		output = append(output, &person)
	}
	return output
}
