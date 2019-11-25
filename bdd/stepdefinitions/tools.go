package stepdefinitions

import (
	"fmt"

	"github.com/DATA-DOG/godog"
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

// TheListOfTheToolShouldHaveTheLengthOf ...
func (w *World) TheListOfTheToolShouldHaveTheLengthOf(length int) error {
	actual := len(w.ToolList)
	if actual != length {
		return fmt.Errorf("expected length of %d is different than %d", length, actual)
	}

	return nil
}

// WeAskForTheTool ...
func (w *World) WeAskForTheTool(name string) error {
	return godog.ErrPending
}
