package stepdefinitions

import (
	"github.com/DATA-DOG/godog"
	"github.com/seadiaz/adoption/bdd/drivers"
)

// TheToolIsMarkedWithLabel ...
func (w *World) TheToolIsMarkedWithLabel(toolName string, label string) error {
	tool := w.Tools[toolName].(map[string]interface{})
	_, err := drivers.AddLabelToTool(label, tool["id"].(string))
	return err
}

// TheListOfTheLabelsShouldHaveTheLengthOf ...
func (w *World) TheListOfTheLabelsShouldHaveTheLengthOf(quantity int) error {
	return godog.ErrPending
}

// TheListOfTheLabelsShouldContainsTo ...
func (w *World) TheListOfTheLabelsShouldContainsTo(name string) error {
	return godog.ErrPending
}
