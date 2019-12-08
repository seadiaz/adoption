package stepdefinitions

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/seadiaz/adoption/bdd/drivers"
)

// ThereIsAToolNamed ...
func (w *World) ThereIsAToolNamed(name string) error {
	res, err := drivers.CreateToolWithName(name)
	w.Tools[name] = res
	return err
}

// WeTryToCreateAToolNamed ...
func (w *World) WeTryToCreateAToolNamed(name string) error {
	res, err := drivers.CreateToolWithName(name)
	if err == nil {
		glog.Info(err)
		w.Tools[name] = res
	}
	return nil
}

// WeAskForTheListOfManagedTools ...
func (w *World) WeAskForTheListOfManagedTools() error {
	res, err := drivers.GetAllTools()
	w.ToolList = res
	return err
}

// WeAskForTheListOfManagedToolsFilterByLabelTeamEquals ...
func (w *World) WeAskForTheListOfManagedToolsFilterByLabelTeamEquals(labelKind string, labelValue string) error {
	res, err := drivers.GetAllToolsFilterByLabel(labelKind, labelValue)
	w.ToolList = res
	return err
}

// TheListOfTheToolShouldHaveTheLengthOf ...
func (w *World) TheListOfTheToolShouldHaveTheLengthOf(length int) error {
	actual := len(w.ToolList)
	if actual != length {
		return fmt.Errorf("expected length of %d is different than %d", length, actual)
	}

	return nil
}

// TheListOfTheToolShouldContainsTo ...
func (w *World) TheListOfTheToolShouldContainsTo(name string) error {
	list := w.ToolList
	for _, item := range list {
		if item.(map[string]interface{})["name"] == name {
			return nil
		}
	}

	return fmt.Errorf("tool %s not found", name)
}

// WeAskForTheTool ...
func (w *World) WeAskForTheTool(name string) error {
	res, err := drivers.GetAllTools()
	for _, item := range res {
		if name == item.(map[string]interface{})["name"] {
			w.Tools[name] = item
			w.LabelList = item.(map[string]interface{})["labels"].([]interface{})
		}
	}
	return err
}
