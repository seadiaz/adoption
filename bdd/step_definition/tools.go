package stepdefinition

import "fmt"

var Godogs int

// ThereAreGodogs ...
func ThereAreGodogs(available int) error {
	Godogs = available
	return nil
}

// IEat ...
func IEat(num int) error {
	if Godogs < num {
		return fmt.Errorf("you cannot eat %d godogs, there are %d available", num, Godogs)
	}
	Godogs -= num
	return nil
}

// ThereShouldBeRemaining ...
func ThereShouldBeRemaining(remaining int) error {
	if Godogs != remaining {
		return fmt.Errorf("expected %d godogs to be remaining, but there is %d", remaining, Godogs)
	}
	return nil
}
