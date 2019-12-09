package main

import (
	"flag"
	"os"
	"testing"

	"github.com/seadiaz/adoption/bdd/drivers"
	"github.com/seadiaz/adoption/bdd/stepdefinitions"

	"github.com/golang/glog"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/colors"
)

var opt = godog.Options{Output: colors.Colored(os.Stdout)}

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &opt)
}

func TestMain(m *testing.M) {
	flag.Parse()
	opt.Paths = flag.Args()

	status := godog.RunWithOptions("adoption", contextInitializer, opt)

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}

func contextInitializer(s *godog.Suite) {
	featureContext(s)
}

func featureContext(s *godog.Suite) {
	world := stepdefinitions.CreateWorld()

	s.Step(`^there is a adoptable named (\w+)$`, world.ThereIsAToolNamed)
	s.Step(`^there is a person named (\w+) which have adopted adoptable (\w+)$`, world.ThereIsAPersonNamedWhichHaveAdoptedTool)
	s.Step(`^we ask for the level of adoption of the adoptable (\w+)$`, world.WeAskForTheLevelOfAdoptionOfTheTool)
	s.Step(`^we ask for the list of managed adoptables$`, world.WeAskForTheListOfManagedTools)
	s.Step(`^we try to create a adoptable named (\w+)$`, world.WeTryToCreateAToolNamed)
	s.Step(`^the adoption level of the adoptable (\w+) should be (\d+) percent$`, world.TheAdoptionLevelOfTheToolShouldBePercent)
	s.Step(`^the list of adopters of the adoptable (\w+) should contain to (\w+)$`, world.TheListOfAdoptersOfTheToolShouldContainTo)
	s.Step(`^the list of adopters of the adoptable (\w+) should not contain to (\w+)$`, world.TheListOfAdoptersOfTheToolShouldNotContainTo)
	s.Step(`^the list of absentees of the adoptable (\w+) should contain to (\w+)$`, world.TheListOfAbsenteesOfTheToolShouldContainTo)
	s.Step(`^the list of absentees of the adoptable (\w+) should not contain to (\w+)$`, world.TheListOfAbsenteesOfTheToolShouldNotContainTo)
	s.Step(`^the list of team adopters of the adoptable (\w+) should contain to (\w+)$`, world.TheListOfTeamAdoptersOfTheToolShouldContainTo)
	s.Step(`^the list of team absentees of the adoptable (\w+) should contain to (\w+)$`, world.TheListOfTeamAbsenteesOfTheToolShouldContainTo)
	s.Step(`^the list of the adoptable should have the length of (\d+)$`, world.TheListOfTheToolShouldHaveTheLengthOf)
	s.Step(`^there is a person named (\w+)$`, world.ThereIsAPersonNamed)
	s.Step(`^we try to create a person named (\w+)$`, world.WeTryToCreateAPersonNamed)
	s.Step(`^we ask for the list of people$`, world.WeAskForTheListOfPeople)
	s.Step(`^the list of the people should have the length of (\d+)$`, world.TheListOfThePeopleShouldHaveTheLengthOf)
	s.Step(`^there is a team named (\w+)$`, world.ThereIsATeamNamed)
	s.Step(`^the person (\w+) is member of the team (\w+)$`, world.ThePersonIsMemberOfTheTeam)
	s.Step(`^we ask for the members of team (\w+)$`, world.WeAskForTheMembersOfTeam)
	s.Step(`^the list of the members should have the length of (\d+)$`, world.TheListOfTheMembersShouldHaveTheLengthOf)
	s.Step(`^the list of the members should contains to (\w+)$`, world.TheListOfTheMembersShouldContainsTo)
	s.Step(`^the team adoption level of the adoptable (\w+) should be (\d+) percent$`, world.TheTeamAdoptionLevelOfTheToolShouldBePercent)
	s.Step(`^the adoptable (\w+) is marked with (\w+) label as (\w+)$`, world.TheToolIsMarkedWithLabelAs)
	s.Step(`^we ask for the adoptable (\w+)$`, world.WeAskForTheTool)
	s.Step(`^the list of the labels should have the length of (\d+)$`, world.TheListOfTheLabelsShouldHaveTheLengthOf)
	s.Step(`^the list of the labels should contains to (\w+)=(\w+)$`, world.TheListOfTheLabelsShouldContainsTo)
	s.Step(`^we ask for the list of managed adoptables filter by label (\w+) equals (\w+)$`, world.WeAskForTheListOfManagedToolsFilterByLabelTeamEquals)
	s.Step(`^the list of the adoptable should contains to (\w+)$`, world.TheListOfTheToolShouldContainsTo)

	var app *drivers.App
	s.BeforeScenario(func(interface{}) {
		glog.Info("starting app...")
		app = drivers.CreateApp()
		app.StartApp()
		world.Clear()
	})

	s.AfterScenario(func(interface{}, error) {
		glog.Info("stopping app...")
		app.StopApp()
	})
}
