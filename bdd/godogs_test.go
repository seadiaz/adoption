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
	flag.Set("logtostderr", "true")
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

	s.Step(`^there is a tool named (\w+)$`, world.ThereIsAToolNamed)
	s.Step(`^a person named (\w+) which have adopted tool (\w+)$`, world.APersonNamedWhichHaveAdoptedTool)
	s.Step(`^a person named (\w+)$`, world.APersonNamed)
	s.Step(`^we ask for the level of adoption of the tool (\w+)$`, world.WeAskForTheLevelOfAdoptionOfTheTool)
	s.Step(`^we ask for the list of managed tools$`, world.WeAskForTheListOfManagedTools)
	s.Step(`^we try to create a tool named (\w+)$`, world.WeTryToCreateAToolNamed)
	s.Step(`^the adoption level of the tool (\w+) should be (\d+) percent$`, world.TheAdoptionLevelOfTheToolShouldBePercent)
	s.Step(`^the list of adopters of the tool (\w+) should contain to (\w+)$`, world.TheListOfAdoptersOfTheToolShouldContainTo)
	s.Step(`^the list of adopters of the tool (\w+) should not contain to (\w+)$`, world.TheListOfAdoptersOfTheToolShouldNotContainTo)
	s.Step(`^the list of absentees of the tool (\w+) should contain to (\w+)$`, world.TheListOfAbsenteesOfTheToolShouldContainTo)
	s.Step(`^the list of absentees of the tool (\w+) should not contain to (\w+)$`, world.TheListOfAbsenteesOfTheToolShouldNotContainTo)
	s.Step(`^the list of the tool should have the length of (\d+)$`, world.TheListOfTheToolShouldHaveTheLengthOf)

	s.Step(`^there is a person named (\w+)$`, world.APersonNamed)
	s.Step(`^we try to create a person named (\w+)$`, world.WeTryToCreateAPersonNamed)
	s.Step(`^we ask for the list of people$`, world.WeAskForTheListOfPeople)
	s.Step(`^the list of the people should have the length of (\d+)$`, world.TheListOfThePeopleShouldHaveTheLengthOf)

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
