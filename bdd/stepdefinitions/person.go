package stepdefinitions

import "github.com/seadiaz/adoption/bdd/drivers"

// APersonNamedWhichHaveAdoptedTool ...
func (w *World) APersonNamedWhichHaveAdoptedTool(personName string, toolName string) error {
	res, err := drivers.CreatePersonWithName(personName)
	if err != nil {
		return err
	}
	w.People[personName] = res

	tool := w.Tools[toolName].(map[string]interface{})
	_, err = drivers.AdoptToolByPerson(tool["id"].(string), res["email"].(string))
	return err
}

// APersonNamed ...
func (w *World) APersonNamed(name string) error {
	_, err := drivers.CreatePersonWithName(name)
	return err
}
