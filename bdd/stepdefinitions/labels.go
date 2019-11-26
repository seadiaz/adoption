package stepdefinitions

import (
	"fmt"

	"github.com/seadiaz/adoption/bdd/drivers"
)

// TheToolIsMarkedWithLabelAs ...
func (w *World) TheToolIsMarkedWithLabelAs(toolName string, labelKind string, labelValue string) error {
	tool := w.Tools[toolName].(map[string]interface{})
	_, err := drivers.AddLabelToTool(labelKind, labelValue, tool["id"].(string))
	return err
}

// TheListOfTheLabelsShouldHaveTheLengthOf ...
func (w *World) TheListOfTheLabelsShouldHaveTheLengthOf(quantity int) error {
	labelList := w.LabelList
	if len(labelList) != quantity {
		return fmt.Errorf("expected length of %d is different than %d", quantity, len(labelList))
	}

	return nil
}

// TheListOfTheLabelsShouldContainsTo ...
func (w *World) TheListOfTheLabelsShouldContainsTo(kind string, value string) error {
	labelList := w.LabelList
	for _, item := range labelList {
		actualKind := item.(map[string]interface{})["kind"]
		actualValue := item.(map[string]interface{})["value"]
		if actualKind == kind && actualValue == value {
			return nil
		}
	}

	return fmt.Errorf("labels not contains label %s=%s", kind, value)
}
