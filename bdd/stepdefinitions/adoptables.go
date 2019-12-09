package stepdefinitions

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/seadiaz/adoption/bdd/drivers"
)

// ThereIsAnAdoptableNamed ...
func (w *World) ThereIsAnAdoptableNamed(name string) error {
	res, err := drivers.CreateAdoptableWithName(name)
	w.Adoptables[name] = res
	return err
}

// WeTryToCreateAnAdoptableNamed ...
func (w *World) WeTryToCreateAnAdoptableNamed(name string) error {
	res, err := drivers.CreateAdoptableWithName(name)
	if err == nil {
		glog.Info(err)
		w.Adoptables[name] = res
	}
	return nil
}

// WeAskForTheListOfManagedAdoptables ...
func (w *World) WeAskForTheListOfManagedAdoptables() error {
	res, err := drivers.GetAllAdoptables()
	w.AdoptableList = res
	return err
}

// WeAskForTheListOfManagedAdoptablesFilterByLabelTeamEquals ...
func (w *World) WeAskForTheListOfManagedAdoptablesFilterByLabelTeamEquals(labelKind string, labelValue string) error {
	res, err := drivers.GetAllAdoptablesFilterByLabel(labelKind, labelValue)
	w.AdoptableList = res
	return err
}

// TheListOfTheAdoptableShouldHaveTheLengthOf ...
func (w *World) TheListOfTheAdoptableShouldHaveTheLengthOf(length int) error {
	actual := len(w.AdoptableList)
	if actual != length {
		return fmt.Errorf("expected length of %d is different than %d", length, actual)
	}

	return nil
}

// TheListOfTheAdoptableShouldContainsTo ...
func (w *World) TheListOfTheAdoptableShouldContainsTo(name string) error {
	list := w.AdoptableList
	for _, item := range list {
		if item.(map[string]interface{})["name"] == name {
			return nil
		}
	}

	return fmt.Errorf("adoptable %s not found", name)
}

// WeAskForTheAdoptable ...
func (w *World) WeAskForTheAdoptable(name string) error {
	res, err := drivers.GetAllAdoptables()
	for _, item := range res {
		if name == item.(map[string]interface{})["name"] {
			w.Adoptables[name] = item
			w.LabelList = item.(map[string]interface{})["labels"].([]interface{})
		}
	}
	return err
}
