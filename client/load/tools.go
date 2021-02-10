package load

// const adoptablesPath = "/adoptables"

// // Adoptable ...
// type Adoptable struct {
// 	ID   string `json:"id,omitempty"`
// 	Name string `json:"name"`
// }

// // LoadAdoptables ...
// func (c *client) LoadAdoptables() {
// 	rawData := readCsvFile(c.filename)
// 	parsedData := mapArrayToAdoptables(rawData)
// 	c.postAdoptables(parsedData)
// }

// func mapArrayToAdoptables(array [][]string) []*Adoptable {
// 	output := make([]*Adoptable, 0, 0)
// 	for _, item := range array {
// 		output = append(output, &Adoptable{
// 			Name: item[0],
// 		})
// 	}
// 	return output
// }

// func (c *client) postAdoptables(adoptables []*Adoptable) {
// 	channel := make(chan string)
// 	for _, item := range adoptables {
// 		go c.postAdoptable(item, channel)
// 	}

// 	receiveResponses(channel, len(adoptables))
// }

// func (c *client) postAdoptable(adoptable *Adoptable, channel chan string) {
// 	doPostRequest(adoptable, c.url+adoptablesPath, c.apiKey)
// 	channel <- adoptable.Name
// }

// func (c *client) getAdoptables() []*Adoptable {
// 	res, _ := doGetRequest(c.url+adoptablesPath, c.apiKey)
// 	output := make([]*Adoptable, 0, 0)
// 	for _, item := range res {
// 		output = append(output, &Adoptable{
// 			ID:   item["id"].(string),
// 			Name: item["name"].(string),
// 		})
// 	}
// 	return output
// }

// func findAdoptableByName(adoptables []*Adoptable, name string) *Adoptable {
// 	for _, item := range adoptables {
// 		if item.Name == name {
// 			return item
// 		}
// 	}

// 	return nil
// }
