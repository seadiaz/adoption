package people

import (
	"github.com/seadiaz/adoption/client/utils"
)

func load(r *Repository, filename string) {
	var people []Person
	utils.ReadCsvFile(filename, &people)
	channel := make(chan string)
	for _, item := range people {
		go postPerson(r, item, channel)
	}

	utils.ReceiveResponses(channel, len(people))
}

func postPerson(r *Repository, person Person, channel chan string) {
	r.Client.DoPostRequest(Path, person)
	channel <- person.Name
}
