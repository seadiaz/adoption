package client

const peoplePath = "/people"

// Person ...
type Person struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// LoadPeople ...
func (c *client) LoadPeople() {
	rawData := readCsvFile(c.filename)
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

func (c *client) postPeople(people []*Person) {
	for _, item := range people {
		c.postPerson(item)
	}
}

func (c *client) postPerson(person *Person) {
	doPostRequest(person, c.url+peoplePath, c.apiKey)
}
