package stepdefinitions

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/seadiaz/adoption/bdd/drivers"
)

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

// WeTryToCreateAPersonNamed ...
func (w *World) WeTryToCreateAPersonNamed(name string) error {
	_, err := drivers.CreatePersonWithName(name)
	if err != nil {
		glog.Error(err)
	}
	return nil
}

// WeAskForTheListOfPeople ...
func (w *World) WeAskForTheListOfPeople() error {
	res, err := drivers.GetAllPeople()
	w.PeopleList = res
	return err
}

// TheListOfThePeopleShouldHaveTheLengthOf ...
func (w *World) TheListOfThePeopleShouldHaveTheLengthOf(length int) error {
	actual := len(w.PeopleList)
	if actual != length {
		return fmt.Errorf("expected length of %d is different than %d", length, actual)
	}

	return nil
}
