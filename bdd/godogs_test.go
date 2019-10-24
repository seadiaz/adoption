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

	status := godog.RunWithOptions("godogs", func(s *godog.Suite) {
		FeatureContext(s)
	}, opt)

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^there is a tool named (\w+)$`, stepdefinitions.ThereIsAToolNamed)
	s.Step(`^a person named (\w+) which have adopted tool (\w+)$`, stepdefinitions.APersonNamedWhichHaveAdoptedTool)
	s.Step(`^a person named (\w+)$`, stepdefinitions.APersonNamed)
	s.Step(`^we ask for the level of adoption of the tool (\w+)$`, stepdefinitions.WeAskForTheLevelOfAdoptionOfTheTool)
	s.Step(`^the adoption level of tool (\w+) should be (\d+) percent$`, stepdefinitions.TheAdoptionLevelOfToolShouldBePercent)

	s.BeforeScenario(func(interface{}) {
	})
	s.BeforeSuite(func() {
		glog.Infoln("starting app...")
		drivers.StartApp()
	})
}
