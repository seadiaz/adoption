package client

const peoplePath = "/people"

// Person ...
type Person struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// LoadPeople ...
func (c *CommandDispatcher) LoadPeople(filename string) {
	rawData := readCsvFile(filename)
	parsedData := mapArrayToPeople(rawData)
	c.postPeople(parsedData)
}

func mapArrayToPeople(array [][]string) []*Person {
	output := make([]*Person, 0, 0)
	for _, item := range array {
		output = append(output, &Person{
			Name:  item[0],
			Email: item[1],
		})
	}
	return output
}

func (c *CommandDispatcher) postPeople(people []*Person) {
	channel := make(chan string)
	for _, item := range people {
		go c.postPerson(item, channel)
	}

	receiveResponses(channel, len(people))
}

func (c *CommandDispatcher) postPerson(person *Person, channel chan string) {
	c.apiClient.doPostRequest(peoplePath, person)
	channel <- person.Name
}

func (c *CommandDispatcher) getPeople() []*Person {
	res, _ := c.apiClient.doGetRequest(peoplePath)
	output := make([]*Person, 0, 0)
	for _, item := range res {
		output = append(output, &Person{
			ID:    item["id"].(string),
			Name:  item["name"].(string),
			Email: item["email"].(string),
		})
	}
	return output
}

func findPersonByEmail(people []*Person, email string) *Person {
	for _, item := range people {
		if item.Email == email {
			return item
		}
	}

	return nil
}
