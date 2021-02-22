package people

import (
	"github.com/mitchellh/mapstructure"
	"github.com/seadiaz/adoption/client/utils"
)

// GetPeople ...
func GetPeople(url, apiKey string) []*Person {
	res, _ := utils.DoGetRequest(url, apiKey)
	output := make([]*Person, 0, 0)
	for _, item := range res {
		var person Person
		mapstructure.Decode(item, &person)
		output = append(output, &person)
	}
	return output
}
