package stepdefinitions

import (
	"github.com/seadiaz/adoption/bdd/drivers"
)

var people map[string]interface{} = make(map[string]interface{})

// APersonNamedWhichHaveAdoptedTool ...
func APersonNamedWhichHaveAdoptedTool(personName string, toolName string) error {
	res, err := drivers.CreatePersonWithName(personName)
	if err != nil {
		return err
	}
	people[personName] = res

	tool := tools[toolName].(map[string]interface{})
	_, err = drivers.AdoptToolByPerson(tool["ID"].(string), res["Email"].(string))
	return err
}

// APersonNamed ...
func APersonNamed(name string) error {
	_, err := drivers.CreatePersonWithName(name)
	return err
}
