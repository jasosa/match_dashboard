package console

import (
	"fmt"
	"github.com/jasosa/football_scoring_dashboard/pkg/dashboard"
	"sort"
	"strconv"
	"strings"
)

//Adapter Adapter between a MatchDashboard and the console
type Adapter struct {
	match   dashboard.ScoringMatch
	Message chan string
}

//NewAdapter New creates a new instance of Adapter
func NewAdapter(match dashboard.ScoringMatch) *Adapter {
	ca := Adapter{
		match:   match,
		Message: make(chan string),
	}

	return &ca
}

//Execute executes a console command and send a message in the channel
//with the result based on the value messageBack
func (c *Adapter) Execute(command string, messageBack bool) {

	cmd, err := ParseCommand(command)
	if err == nil {
		switch cmd.Name {
		case "Start":
			c.processStartCommand(cmd, messageBack)
		case "End":
			c.processEndCommand(cmd, messageBack)
		case "Add":
			c.processAddCommand(cmd, messageBack)
		case "Print":
			c.processPrintCommand(cmd, messageBack)
		}
	} else {
		if c.match.IsStarted() {
			sendMessage(c.Message, MsgNotRecognizedCommandInMatch)
		} else {
			sendMessage(c.Message, MsgNotRecognizedCommandNotInMatch)
		}
	}
}

//ExecuteWithMessage executes a console command and send a message in the channel
//with the result
func (c *Adapter) ExecuteWithMessage(command string) {
	c.Execute(command, true)
}

func (c *Adapter) print() {
	sb := strings.Builder{}
	sb.WriteString(c.match.GetHomeTeam())
	sb.WriteString(" ")
	sb.WriteString(strconv.Itoa(c.match.GetHomeTeamScore()))
	if len(c.match.GetHomeGoals()) > 0 {
		sb.WriteString(" (")
		sb.WriteString(printGoals(c.match.GetHomeGoals()))
		sb.WriteString(")")
	}
	sb.WriteString(" vs. ")
	sb.WriteString(c.match.GetAwayTeam())
	sb.WriteString(" ")
	sb.WriteString(strconv.Itoa(c.match.GetAwayTeamScore()))
	if len(c.match.GetAwayGoals()) > 0 {
		sb.WriteString(" (")
		sb.WriteString(printGoals(c.match.GetAwayGoals()))
		sb.WriteString(")")
	}
	sendMessage(c.Message, sb.String())
}

func (c *Adapter) processStartCommand(cmd Command, messageBack bool) {
	err := c.match.Start(cmd.Args[0], cmd.Args[1])
	if err == nil {
		if messageBack {
			sendMessage(c.Message, "Match started succesfully")
		}
	} else if err == dashboard.ErrMatchAlreadyStarted && messageBack {
		sendMessage(c.Message, MsgGameInProgress)
	} else if messageBack {
		sendMessage(c.Message, err.Error())
	}
}

func (c *Adapter) processEndCommand(cmd Command, messageBack bool) {

	err := c.match.End()
	if err == nil && messageBack {
		sendMessage(c.Message, "Match ended succesfully")
	} else if err == dashboard.ErrMatchNotStarted && messageBack {
		sendMessage(c.Message, MsNoGameInProgress)
	} else if messageBack {
		sendMessage(c.Message, err.Error())
	}
}

func (c *Adapter) processAddCommand(cmd Command, messageBack bool) {
	minute, _ := strconv.Atoi(cmd.Args[0])

	err := c.match.AddGoal(minute, cmd.Args[1], cmd.Args[2])
	if err == nil && messageBack {
		sendMessage(c.Message, "Goal added succesfully")
	} else if err == dashboard.ErrMatchNotStarted && messageBack {
		sendMessage(c.Message, MsNoGameInProgress)
	} else if messageBack {
		sendMessage(c.Message, err.Error())
	}
}

func (c *Adapter) processPrintCommand(cmd Command, messageBack bool) {
	if c.match.IsStarted() {
		c.print()
	} else {
		if messageBack {
			sendMessage(c.Message, MsNoGameInProgress)
		}
	}
}

func printGoals(goals []dashboard.Goal) string {
	sb := strings.Builder{}
	sort.Sort(byPlayer(goals))
	currPlayer := ""
	for _, goal := range goals {
		if currPlayer != goal.Player {
			sb.WriteString(goal.Player)
			sb.WriteString(" ")
		}

		sb.WriteString(fmt.Sprintf("%d'", goal.Minute))
		sb.WriteString(" ")
		currPlayer = goal.Player
	}

	return strings.TrimSpace(sb.String())
}

func sendMessage(ch chan string, message string) {
	go func() {
		ch <- message
	}()
}
