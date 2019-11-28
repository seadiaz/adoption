package stepdefinitions

import (
	"fmt"

	"github.com/seadiaz/adoption/bdd/drivers"
)

// ThereIsATeamNamed ...
func (w *World) ThereIsATeamNamed(name string) error {
	res, err := drivers.CreateTeamWithName(name)
	w.Teams[name] = res
	return err
}

// ThePersonIsMemberOfTheTeam ...
func (w *World) ThePersonIsMemberOfTheTeam(personName string, teamName string) error {
	team := w.Teams[teamName].(map[string]interface{})
	person := w.People[personName].(map[string]interface{})
	_, err := drivers.AddMemberToTeam(person["id"].(string), team["id"].(string))
	return err
}

// WeAskForTheMembersOfTeam ...
func (w *World) WeAskForTheMembersOfTeam(teamName string) error {
	team := w.Teams[teamName].(map[string]interface{})
	res, err := drivers.GetMembersFromTeam(team["id"].(string))
	w.PeopleList = res
	return err
}

// TheListOfTheMembersShouldHaveTheLengthOf ...
func (w *World) TheListOfTheMembersShouldHaveTheLengthOf(lengthOf int) error {
	actual := len(w.PeopleList)
	if actual != lengthOf {
		return fmt.Errorf("expected length of %d is different than %d", lengthOf, actual)
	}

	return nil
}

// TheListOfTheMembersShouldContainsTo ...
func (w *World) TheListOfTheMembersShouldContainsTo(personName string) error {
	people := w.PeopleList
	for _, item := range people {
		if item.(map[string]interface{})["name"] == personName {
			return nil
		}
	}

	return fmt.Errorf("person %s not found", personName)
}
