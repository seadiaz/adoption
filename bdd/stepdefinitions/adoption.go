package stepdefinitions

import (
	"fmt"

	"github.com/seadiaz/adoption/bdd/drivers"
)

// WeAskForTheLevelOfAdoptionOfTheTool ...
func (w *World) WeAskForTheLevelOfAdoptionOfTheTool(toolName string) error {
	tool := w.Tools[toolName].(map[string]interface{})
	toolID := tool["id"].(string)
	res, err := drivers.CalculateAdoptionForTool(toolID)
	w.Adoptions[toolName] = res
	return err
}

// TheAdoptionLevelOfTheToolShouldBePercent ...
func (w *World) TheAdoptionLevelOfTheToolShouldBePercent(toolName string, percent float64) error {
	adoption := w.Adoptions[toolName].(map[string]interface{})
	actual := adoption["adoption"].(float64)
	if actual != percent {
		return fmt.Errorf("expected percent %f is different than %f", percent, actual)
	}

	return nil
}

// TheListOfAdoptersOfTheToolShouldContainTo ...
func (w *World) TheListOfAdoptersOfTheToolShouldContainTo(toolName string, personName string) error {
	adoption := w.Adoptions[toolName].(map[string]interface{})
	adopters := adoption["adopters"].([]interface{})
	for _, item := range adopters {
		if item.(map[string]interface{})["name"] == personName {
			return nil
		}
	}

	return fmt.Errorf("person %s not found", personName)
}

// TheListOfAdoptersOfTheToolShouldNotContainTo ...
func (w *World) TheListOfAdoptersOfTheToolShouldNotContainTo(toolName string, personName string) error {
	adoption := w.Adoptions[toolName].(map[string]interface{})
	adopters := adoption["adopters"].([]interface{})
	for _, item := range adopters {
		if item.(map[string]interface{})["name"] == personName {
			return fmt.Errorf("person %s found", personName)
		}
	}

	return nil
}

// TheListOfAbsenteesOfTheToolShouldContainTo ...
func (w *World) TheListOfAbsenteesOfTheToolShouldContainTo(toolName string, personName string) error {
	adoption := w.Adoptions[toolName].(map[string]interface{})
	absentees := adoption["absentees"].([]interface{})
	for _, item := range absentees {
		if item.(map[string]interface{})["name"] == personName {
			return nil
		}
	}

	return fmt.Errorf("person %s not found", personName)
}

// TheListOfAbsenteesOfTheToolShouldNotContainTo ...
func (w *World) TheListOfAbsenteesOfTheToolShouldNotContainTo(toolName string, personName string) error {
	adoption := w.Adoptions[toolName].(map[string]interface{})
	absentees := adoption["absentees"].([]interface{})
	for _, item := range absentees {
		if item.(map[string]interface{})["name"] == personName {
			return fmt.Errorf("person %s found", personName)
		}
	}

	return nil
}
