package load

import (
	"github.com/seadiaz/adoption/client/utils"
)

// const peoplePath = "/people"

// // Person ...
// type Person struct {
// 	ID    string `json:"id"`
// 	Name  string `json:"name"`
// 	Email string `json:"email"`
// }

func (c *client) loadPeople() {
	rawData := utils.ReadCsvFile(c.filename)
	parsedData := mapArrayToPeople(rawData)
	c.postPeople(parsedData)
}

// func mapArrayToPeople(array [][]string) []*Person {
// 	output := make([]*Person, 0, 0)
// 	for _, item := range array {
// 		output = append(output, &Person{
// 			Name:  item[0],
// 			Email: item[1],
// 		})
// 	}
// 	return output
// }

// func (c *client) postPeople(people []*Person) {
// 	channel := make(chan string)
// 	for _, item := range people {
// 		go c.postPerson(item, channel)
// 	}

// 	receiveResponses(channel, len(people))
// }

// func (c *client) postPerson(person *Person, channel chan string) {
// 	doPostRequest(person, c.url+peoplePath, c.apiKey)
// 	channel <- person.Name
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
