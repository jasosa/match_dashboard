package console

import (
	"errors"
	"regexp"
	"strings"
)

const (
	//StartCommand command to start the match
	StartCommand uint8 = iota
)

//ParseCommand returns a Command based in the given command string
func ParseCommand(command string) (Command, error) {
	if command == "Print" {
		return parsePrintCommand()
	}
	if command == "End" {
		return parseEndCommand()
	}

	if strings.HasPrefix(command, "Start:") {
		return parseStartCommand(command)
	}
	return parseAddGoalCommand(command)
}

//ParsePrintCommand returns Print command information
func parsePrintCommand() (Command, error) {
	return Command{
		Name: "Print",
		Args: []string{},
	}, nil
}

//ParseEndCommand returns end command information
func parseEndCommand() (Command, error) {
	return Command{
		Name: "End",
		Args: []string{},
	}, nil
}

//ParseStartCommand returns StartCommand information
func parseStartCommand(command string) (Command, error) {

	exp, err := regexp.Compile("^Start:'(.+)' vs. '(.+)'$")

	if err != nil {
		return Command{}, err
	}

	subm := exp.FindAllStringSubmatch(command, -1)
	if subm != nil {
		cmd := Command{
			Name: "Start",
			Args: []string{
				subm[0][1],
				subm[0][2]},
		}
		return cmd, nil
	}

	return Command{}, errors.New("Command does not match")
}

//ParseAddGoalCommand returns addGoal command information
func parseAddGoalCommand(command string) (Command, error) {

	exp, err := regexp.Compile("^([0-9]{1,3}) '(.+)' (.+)$")

	if err != nil {
		return Command{}, err
	}

	subm := exp.FindAllStringSubmatch(command, -1)
	if subm != nil {
		cmd := Command{
			Name: "Add",
			Args: []string{
				subm[0][1],
				subm[0][2],
				subm[0][3]},
		}
		return cmd, nil
	}

	return Command{}, errors.New("Command does not match")
}

//Command represents a command struct
type Command struct {
	Name string
	Args []string
}
