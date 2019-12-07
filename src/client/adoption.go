package client

// Adoption ...
type Adoption struct {
	PersonEmail string
	ToolName    string
	Tool        *Tool
	Person      *Person
}

// LoadAdoptions ...
func (c *client) LoadAdoptions() {
	rawData := readCsvFile(c.filename)
	parsedData := mapArrayToAdoptions(rawData)
	tools := c.getTools()
	people := c.getPeople()
	parsedData = fulfillAdoptionToolIDFromTools(parsedData, tools, people)
	c.postAdoptions(parsedData)
}

func mapArrayToAdoptions(array [][]string) []*Adoption {
	output := make([]*Adoption, 0, 0)
	for _, item := range array {
		output = append(output, &Adoption{
			PersonEmail: item[0],
			ToolName:    item[1],
		})
	}
	return output
}

func fulfillAdoptionToolIDFromTools(adoptions []*Adoption, tools []*Tool, people []*Person) []*Adoption {
	output := make([]*Adoption, 0, 0)
	for _, item := range adoptions {
		tool := findToolByName(tools, item.ToolName)
		person := findPersonByEmail(people, item.PersonEmail)
		if tool != nil {
			item.Tool = tool
			item.Person = person
		}
		output = append(output, item)
	}
	return output
}

func (c *client) postAdoptions(adoptions []*Adoption) {
	channel := make(chan string)
	for _, item := range adoptions {
		go c.postAdoption(item, channel)
	}

	receiveResponses(channel, len(adoptions))
}

func (c *client) postAdoption(adoption *Adoption, channel chan string) {
	doPostRequest(adoption.Tool, c.url+peoplePath+"/"+adoption.Person.ID+toolsPath, c.apiKey)
	channel <- adoption.PersonEmail + " " + adoption.ToolName
}
