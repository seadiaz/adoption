package stepdefinitions

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/seadiaz/adoption/bdd/drivers"
)

// ThereIsAPersonNamedWhichHaveAdoptedAdoptable ...
func (w *World) ThereIsAPersonNamedWhichHaveAdoptedAdoptable(personName string, adoptableName string) error {
	res, err := drivers.CreatePersonWithName(personName)
	if err != nil {
		return err
	}
	w.People[personName] = res

	adoptable := w.Adoptables[adoptableName].(map[string]interface{})
	_, err = drivers.AdoptAdoptableByPerson(adoptable["id"].(string), res["id"].(string))
	return err
}

// ThereIsAPersonNamed ...
func (w *World) ThereIsAPersonNamed(name string) error {
	res, err := drivers.CreatePersonWithName(name)
	w.People[name] = res
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
