package people

import (
	"github.com/seadiaz/adoption/client/utils"
)

func (c *CommandHandler) load(filename string) {
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
