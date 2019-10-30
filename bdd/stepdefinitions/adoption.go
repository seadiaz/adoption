package stepdefinitions

import (
	"fmt"

	"github.com/seadiaz/adoption/bdd/drivers"
)

var adoptions map[string]interface{} = make(map[string]interface{})

// WeAskForTheLevelOfAdoptionOfTheTool ...
func WeAskForTheLevelOfAdoptionOfTheTool(toolName string) error {
	tool := tools[toolName].(map[string]interface{})
	toolID := tool["id"].(string)
	res, err := drivers.CalculateAdoptionForTool(toolID)
	adoptions[toolName] = res
	return err
}

// TheAdoptionLevelOfTheToolShouldBePercent ...
func TheAdoptionLevelOfTheToolShouldBePercent(toolName string, percent float64) error {
	adoption := adoptions[toolName].(map[string]interface{})
	actual := adoption["adoption"].(float64)
	if actual != percent {
		return fmt.Errorf("expected percent %f is different than %f", percent, actual)
	}

	return nil
}

// TheListOfAdoptersOfTheToolShouldContainTo ...
func TheListOfAdoptersOfTheToolShouldContainTo(toolName string, personName string) error {
	adoption := adoptions[toolName].(map[string]interface{})
	adopters := adoption["adopters"].([]interface{})
	for _, item := range adopters {
		if item.(map[string]interface{})["name"] == personName {
			return nil
		}
	}

	return fmt.Errorf("person %s not found", personName)
}

// TheListOfAdoptersOfTheToolShouldNotContainTo ...
func TheListOfAdoptersOfTheToolShouldNotContainTo(toolName string, personName string) error {
	adoption := adoptions[toolName].(map[string]interface{})
	adopters := adoption["adopters"].([]interface{})
	for _, item := range adopters {
		if item.(map[string]interface{})["name"] == personName {
			return fmt.Errorf("person %s found", personName)
		}
	}

	return nil
}

// TheListOfAbsenteesOfTheToolShouldContainTo ...
func TheListOfAbsenteesOfTheToolShouldContainTo(toolName string, personName string) error {
	adoption := adoptions[toolName].(map[string]interface{})
	absentees := adoption["absentees"].([]interface{})
	for _, item := range absentees {
		if item.(map[string]interface{})["name"] == personName {
			return nil
		}
	}

	return fmt.Errorf("person %s not found", personName)
}

// TheListOfAbsenteesOfTheToolShouldNotContainTo ...
func TheListOfAbsenteesOfTheToolShouldNotContainTo(toolName string, personName string) error {
	adoption := adoptions[toolName].(map[string]interface{})
	absentees := adoption["absentees"].([]interface{})
	for _, item := range absentees {
		if item.(map[string]interface{})["name"] == personName {
			return fmt.Errorf("person %s found", personName)
		}
	}

	return nil
}
