package client

import "github.com/golang/glog"

// Adoption ...
type Adoption struct {
	PersonEmail   string
	AdoptableName string
	Adoptable     *Adoptable
	Person        *Person
}

// LoadAdoptions ...
func (c *client) LoadAdoptions() {
	rawData := readCsvFile(c.filename)
	parsedData := mapArrayToAdoptions(rawData)
	adoptables := c.getAdoptables()
	people := c.getPeople()
	parsedData = fulfillAdoptionAdoptableIDFromAdoptables(parsedData, adoptables, people)
	c.postAdoptions(parsedData)
}

func mapArrayToAdoptions(array [][]string) []*Adoption {
	output := make([]*Adoption, 0, 0)
	for _, item := range array {
		output = append(output, &Adoption{
			PersonEmail:   item[0],
			AdoptableName: item[1],
		})
	}
	return output
}

func fulfillAdoptionAdoptableIDFromAdoptables(adoptions []*Adoption, adoptables []*Adoptable, people []*Person) []*Adoption {
	output := make([]*Adoption, 0, 0)
	for _, item := range adoptions {
		adoptable := findAdoptableByName(adoptables, item.AdoptableName)
		person := findPersonByEmail(people, item.PersonEmail)
		if adoptable != nil {
			item.Adoptable = adoptable
			item.Person = person
		}
		output = append(output, item)
	}
	return output
}

func (c *client) postAdoptions(adoptions []*Adoption) {
	for _, item := range adoptions {
		c.postAdoption(item)
	}
}

func (c *client) postAdoption(adoption *Adoption) {
	err := doPostRequest(adoption.Adoptable, c.url+peoplePath+"/"+adoption.Person.ID+adoptablesPath, c.apiKey)
	if err != nil {
		glog.Errorf("fail adding adoption %s by %s: %s", adoption.AdoptableName, adoption.PersonEmail, err.Error())
	} else {
		glog.Infof("adoptable %s adopted by %s", adoption.AdoptableName, adoption.PersonEmail)
	}
}
