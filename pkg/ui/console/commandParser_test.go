package console_test

import (
	"github.com/jasosa/football_scoring_dashboard/pkg/ui/console"
	"testing"
)

func TestParseStartCommand(t *testing.T) {
	cmd, _ := console.ParseCommand("Start:'England' vs. 'West Germany'")
	if cmd.Name == "Start" && cmd.Args[0] == "England" && cmd.Args[1] == "West Germany" {
		t.Log("Start command parsed succesfully")
	} else {
		t.Errorf(
			"Error parsing start command. Expected values were: Name='%s', Args={%s, %s} but returned values were: Name='%s', Args={%s, %s}",
			"Start", "England", "West Germany", cmd.Name, cmd.Args[0], cmd.Args[1])
	}
}

func TestParseWrongStartCommand(t *testing.T) {
	_, err := console.ParseCommand("StartX:'England' vs. 'West Germany'")
	if err != nil {
		t.Log("Error parsing start command", err.Error())
	} else {
		t.Error("Parse a wrong error command should give an error")
	}
}

func TestParseAddGoalCommand(t *testing.T) {
	cmd, _ := console.ParseCommand("11 'West Germany' Haller")
	if cmd.Name == "Add" && cmd.Args[0] == "11" && cmd.Args[1] == "West Germany" && cmd.Args[2] == "Haller" {
		t.Log("Add goal command parsed succesfully")
	} else {
		t.Errorf(
			"Error parsing add goal command. Expected values were: Name='%s', Args={%s, %s, %s} but returned values were: Name='%s', Args={%s, %s, %s}",
			"Add", "11", "West Germany", "Haller", cmd.Name, cmd.Args[0], cmd.Args[1], cmd.Args[2])
	}
}

func TestParsePrintCommand(t *testing.T) {
	cmd, _ := console.ParseCommand("print")
	if cmd.Name != "Print" {
		t.Errorf("Error parsing print command. Expected command name was %s but is %s", "Print", cmd.Name)
	} else if len(cmd.Args) > 0 {
		t.Errorf("Error parsing print command. Expected number of args was %d but is %d", 0, len(cmd.Args))
	}
}
