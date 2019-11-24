package client

// Adoption ...
type Adoption struct {
	PersonEmail string
	ToolName    string
	Tool        *Tool
}

// LoadAdoptions ...
func (c *client) LoadAdoptions() {
	rawData := readCsvFile(c.filename)
	parsedData := mapArrayToAdoptions(rawData)
	tools := c.getTools()
	parsedData = fulfillAdoptionToolIDFromTools(parsedData, tools)
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

func fulfillAdoptionToolIDFromTools(adoptions []*Adoption, tools []*Tool) []*Adoption {
	output := make([]*Adoption, 0, 0)
	for _, item := range adoptions {
		tool := findToolByName(tools, item.ToolName)
		if tool != nil {
			item.Tool = tool
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
	doPostRequest(adoption.Tool, c.url+peoplePath+"/"+adoption.PersonEmail+toolsPath, c.apiKey)
	channel <- adoption.PersonEmail + " " + toolsPath
}
