package main

import (
	"flag"
	"os"
	"testing"

	"github.com/golang/glog"
	stepdefinition "github.com/seadiaz/adoption/bdd/step_definition"

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

	status := godog.RunWithOptions("godogs", func(s *godog.Suite) {
		FeatureContext(s)
	}, opt)

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^there are (\d+) godogs$`, stepdefinition.ThereAreGodogs)
	s.Step(`^I eat (\d+)$`, stepdefinition.IEat)
	s.Step(`^there should be (\d+) remaining$`, stepdefinition.ThereShouldBeRemaining)

	s.BeforeScenario(func(interface{}) {
		stepdefinition.Godogs = 0 // clean the state before every scenario
	})
	s.BeforeSuite(func() {
		glog.Infoln("Before suite")
	})
}
