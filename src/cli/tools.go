package cli

const toolsPath = "/tools"

// Tool ...
type Tool struct {
	Name string `json:"name"`
}

// LoadTools ...
func (c *Client) LoadTools() {
	rawData := readCsvFile(c.Filename)
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

func (c *Client) postTools(tools []*Tool) {
	for _, item := range tools {
		c.postTool(item)
	}
}

func (c *Client) postTool(tool *Tool) {
	doPostRequest(tool, c.URL+toolsPath)
}
