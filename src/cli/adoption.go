package cli

// Adoption ...
type Adoption struct {
	PersonEmail string
	ToolName    string
	Tool        *Tool
}

// LoadAdoptions ...
func (c *Client) LoadAdoptions() {
	rawData := readCsvFile(c.Filename)
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

func (c *Client) postAdoptions(adoptions []*Adoption) {
	for _, item := range adoptions {
		c.postAdoption(item)
	}
}

func (c *Client) postAdoption(adoption *Adoption) {
	doPostRequest(adoption.Tool, c.URL+peoplePath+"/"+adoption.PersonEmail+toolsPath, c.APIKey)
}
