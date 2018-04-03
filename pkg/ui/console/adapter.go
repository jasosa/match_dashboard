package console

import (
	"fmt"
	"github.com/jasosa/football_scoring_dashboard/pkg/dashboard"
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

//Execute executes a console command
func (c *Adapter) Execute(command string) {
	cmd, err := ParseCommand(command)
	if err == nil {
		switch cmd.Name {
		case "Start":
			c.match.Start(cmd.Args[0], cmd.Args[1])
		case "Add":
			minute, _ := strconv.Atoi(cmd.Args[0])
			c.match.AddGoal(minute, cmd.Args[1], cmd.Args[2])
		case "Print":
			c.print()
		}
	}
}

func (c *Adapter) print() {
	go func() {
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
		c.Message <- sb.String()
	}()
}

func printGoals(goals []dashboard.Goal) string {
	sb := strings.Builder{}
	for _, goal := range goals {
		sb.WriteString(goal.Player)
		sb.WriteString(" ")
		sb.WriteString(fmt.Sprintf("%d'", goal.Minute))
		sb.WriteString(" ")
	}

	return strings.TrimSpace(sb.String())
}
