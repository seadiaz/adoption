package client

const toolsPath = "/tools"

// Tool ...
type Tool struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
}

// LoadTools ...
func (c *client) LoadTools() {
	rawData := readCsvFile(c.filename)
	parsedData := mapArrayToTools(rawData)
	c.postTools(parsedData)
}

func mapArrayToTools(array [][]string) []*Tool {
	output := make([]*Tool, 0, 0)
	for _, item := range array {
		output = append(output, &Tool{
			Name: item[0],
		})
	}
	return output
}

func (c *client) postTools(tools []*Tool) {
	channel := make(chan string)
	for _, item := range tools {
		go c.postTool(item, channel)
	}

	receiveResponses(channel, len(tools))
}

func (c *client) postTool(tool *Tool, channel chan string) {
	doPostRequest(tool, c.url+toolsPath, c.apiKey)
	channel <- toolsPath
}

func (c *client) getTools() []*Tool {
	res, _ := doGetRequest(c.url+toolsPath, c.apiKey)
	output := make([]*Tool, 0, 0)
	for _, item := range res {
		output = append(output, &Tool{
			ID:   item["id"].(string),
			Name: item["name"].(string),
		})
	}
	return output
}

func findToolByName(tools []*Tool, name string) *Tool {
	for _, item := range tools {
		if item.Name == name {
			return item
		}
	}

	return nil
}
