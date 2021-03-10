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

// FindPersonByEmail ...
func FindPersonByEmail(url, apiKey, email string) *Person {
	items := GetPeople(url, apiKey)
	item := filterPersonByEmail(items, email)
	return item
}

func filterPersonByEmail(items []*Person, email string) *Person {
	for _, v := range items {
		if v.Email == email {
			return v
		}
	}

	return nil
}
