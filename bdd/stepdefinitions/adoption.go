package stepdefinitions

import (
	"fmt"

	"github.com/seadiaz/adoption/bdd/drivers"
)

// WeAskForTheLevelOfAdoptionOfTheAdoptable ...
func (w *World) WeAskForTheLevelOfAdoptionOfTheAdoptable(adoptableName string) error {
	adoptable := w.Adoptables[adoptableName].(map[string]interface{})
	adoptableID := adoptable["id"].(string)
	res, err := drivers.CalculateAdoptionForAdoptable(adoptableID)
	w.Adoptions[adoptableName] = res
	return err
}

// TheAdoptionLevelOfTheAdoptableShouldBePercent ...
func (w *World) TheAdoptionLevelOfTheAdoptableShouldBePercent(adoptableName string, percent float64) error {
	adoption := w.Adoptions[adoptableName].(map[string]interface{})
	actual := adoption["adoption"].(float64)
	if actual != percent {
		return fmt.Errorf("expected percent %f is different than %f", percent, actual)
	}

	return nil
}

// TheTeamAdoptionLevelOfTheAdoptableShouldBePercent ...
func (w *World) TheTeamAdoptionLevelOfTheAdoptableShouldBePercent(adoptableName string, percent float64) error {
	adoption := w.Adoptions[adoptableName].(map[string]interface{})
	actual := adoption["team_adoption"].(float64)
	if actual != percent {
		return fmt.Errorf("expected percent %f is different than %f", percent, actual)
	}

	return nil
}

// TheListOfAdoptersOfTheAdoptableShouldContainTo ...
func (w *World) TheListOfAdoptersOfTheAdoptableShouldContainTo(adoptableName string, personName string) error {
	adoption := w.Adoptions[adoptableName].(map[string]interface{})
	adopters := adoption["adopters"].([]interface{})
	for _, item := range adopters {
		if item.(map[string]interface{})["name"] == personName {
			return nil
		}
	}

	return fmt.Errorf("person %s not found", personName)
}

// TheListOfAdoptersOfTheAdoptableShouldNotContainTo ...
func (w *World) TheListOfAdoptersOfTheAdoptableShouldNotContainTo(adoptableName string, personName string) error {
	adoption := w.Adoptions[adoptableName].(map[string]interface{})
	adopters := adoption["adopters"].([]interface{})
	for _, item := range adopters {
		if item.(map[string]interface{})["name"] == personName {
			return fmt.Errorf("person %s found", personName)
		}
	}

	return nil
}

// TheListOfAbsenteesOfTheAdoptableShouldContainTo ...
func (w *World) TheListOfAbsenteesOfTheAdoptableShouldContainTo(adoptableName string, personName string) error {
	adoption := w.Adoptions[adoptableName].(map[string]interface{})
	absentees := adoption["absentees"].([]interface{})
	for _, item := range absentees {
		if item.(map[string]interface{})["name"] == personName {
			return nil
		}
	}

	return fmt.Errorf("person %s not found", personName)
}

// TheListOfAbsenteesOfTheAdoptableShouldNotContainTo ...
func (w *World) TheListOfAbsenteesOfTheAdoptableShouldNotContainTo(adoptableName string, personName string) error {
	adoption := w.Adoptions[adoptableName].(map[string]interface{})
	absentees := adoption["absentees"].([]interface{})
	for _, item := range absentees {
		if item.(map[string]interface{})["name"] == personName {
			return fmt.Errorf("person %s found", personName)
		}
	}

	return nil
}

// TheListOfTeamAdoptersOfTheAdoptableShouldContainTo ...
func (w *World) TheListOfTeamAdoptersOfTheAdoptableShouldContainTo(adoptableName string, teamName string) error {
	adoption := w.Adoptions[adoptableName].(map[string]interface{})
	adopters := adoption["team_adopters"].([]interface{})
	for _, item := range adopters {
		if item.(map[string]interface{})["name"] == teamName {
			return nil
		}
	}

	return fmt.Errorf("team %s not found", teamName)
}

// TheListOfTeamAbsenteesOfTheAdoptableShouldContainTo ...
func (w *World) TheListOfTeamAbsenteesOfTheAdoptableShouldContainTo(adoptableName string, teamName string) error {
	adoption := w.Adoptions[adoptableName].(map[string]interface{})
	adopters := adoption["team_absentees"].([]interface{})
	for _, item := range adopters {
		if item.(map[string]interface{})["name"] == teamName {
			return nil
		}
	}

	return fmt.Errorf("team %s not found", teamName)
}
