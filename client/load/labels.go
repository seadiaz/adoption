package load

// import "github.com/golang/glog"

// var labelsPath = "/labels"

// type label struct {
// 	AdoptableName string
// 	Adoptable     *Adoptable
// 	Kind          string `json:"kind"`
// 	Value         string `json:"value"`
// }

// // LoadLabels ...
// func (c *client) LoadLabels() {
// 	rawData := readCsvFile(c.filename)
// 	parsedData := mapArrayToLabels(rawData)
// 	Adoptables := c.getAdoptables()
// 	parsedData = fulfillLabelAdoptableFromAdoptableName(parsedData, Adoptables)
// 	c.postLabels(parsedData)
// }

// func mapArrayToLabels(array [][]string) []*label {
// 	output := make([]*label, 0, 0)
// 	for _, item := range array {
// 		output = append(output, &label{
// 			AdoptableName: item[0],
// 			Kind:          item[1],
// 			Value:         item[2],
// 		})
// 	}
// 	return output
// }

// func fulfillLabelAdoptableFromAdoptableName(labels []*label, Adoptables []*Adoptable) []*label {
// 	output := make([]*label, 0, 0)
// 	for _, item := range labels {
// 		Adoptable := findAdoptableByName(Adoptables, item.AdoptableName)
// 		if Adoptable != nil {
// 			item.Adoptable = Adoptable
// 			output = append(output, item)
// 		}
// 	}
// 	return output
// }

// func (c *client) postLabels(labels []*label) {
// 	for _, item := range labels {
// 		c.postLabel(item)
// 	}
// }

// func (c *client) postLabel(label *label) {
// 	body := label
// 	err := doPostRequest(body, c.url+adoptablesPath+"/"+label.Adoptable.ID+labelsPath, c.apiKey)
// 	if err != nil {
// 		glog.Errorf("fail adding label %s=%s to %s: %s", label.Kind, label.Value, label.Adoptable.Name, err.Error())
// 	} else {
// 		glog.Infof("%s=%s added to %s", label.Kind, label.Value, label.Adoptable.Name)
// 	}
// }
