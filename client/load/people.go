package load

import (
	"github.com/seadiaz/adoption/client/utils"
)

const peoplePath = "/people"

// Person ...
type Person struct {
	ID    string `json:"id"`
	Name  string `json:"name" csv:"Name"`
	Email string `json:"email" csv:"Email"`
}

func (c *CommandHandler) loadPeople(filename string) {
	var people []Person
	utils.ReadCsvFile(filename, &people)
	channel := make(chan string)
	for _, item := range people {
		go postPerson(c.baseURL+peoplePath, c.apiKey, item, channel)
	}

	utils.ReceiveResponses(channel, len(people))
}

func postPerson(url, apiKey string, person Person, channel chan string) {
	utils.DoPostRequest(url, apiKey, person)
	channel <- person.Name
}

// func (c *client) loadPeople() {
// rawData := utils.ReadCsvFile(c.filename)
// parsedData := mapArrayToPeople(rawData)
// c.postPeople(parsedData)
// }

// func (c *client) postPeople(people []*Person) {
// 	channel := make(chan string)
// 	for _, item := range people {
// 		go c.postPerson(item, channel)
// 	}

// 	receiveResponses(channel, len(people))
// }

// func (c *client) getPeople() []*Person {
// 	res, _ := doGetRequest(c.url+peoplePath, c.apiKey)
// 	output := make([]*Person, 0, 0)
// 	for _, item := range res {
// 		output = append(output, &Person{
// 			ID:    item["id"].(string),
// 			Name:  item["name"].(string),
// 			Email: item["email"].(string),
// 		})
// 	}
// 	return output
// }

// func findPersonByEmail(people []*Person, email string) *Person {
// 	for _, item := range people {
// 		if item.Email == email {
// 			return item
// 		}
// 	}

// 	return nil
// }
