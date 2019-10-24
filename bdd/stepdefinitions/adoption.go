package stepdefinitions

import (
	"fmt"

	"github.com/seadiaz/adoption/bdd/drivers"
)

var adoptions map[string]interface{} = make(map[string]interface{})

// WeAskForTheLevelOfAdoptionOfTheTool ...
func WeAskForTheLevelOfAdoptionOfTheTool(toolName string) error {
	tool := tools[toolName].(map[string]interface{})
	toolID := tool["ID"].(string)
	res, err := drivers.CalculateAdoptionForTool(toolID)
	adoptions[toolName] = res
	return err
}

// TheAdoptionLevelOfToolShouldBePercent ...
func TheAdoptionLevelOfToolShouldBePercent(toolName string, percent float64) error {
	adoption := adoptions[toolName].(map[string]interface{})
	actual := adoption["adoption"].(float64)
	if actual != percent {
		return fmt.Errorf("expected percent %f is different than %f", percent, actual)
	}

	return nil
}
