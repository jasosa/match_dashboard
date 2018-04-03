package console

import (
	"github.com/jasosa/football_scoring_dashboard/pkg/dashboard"
	"testing"
	"time"
)

func TestMatchDashboardStartsSuccesfully(t *testing.T) {
	consoleCmd := "Start:'England' vs. 'West Germany'"
	t.Log("Given the command: ", consoleCmd)
	{
		t.Log("When the command is executed into the dashboard")
		{
			matchDashboard := dashboard.New()
			adp := NewAdapter(matchDashboard)
			adp.Execute(consoleCmd)
			if matchDashboard.IsStarted() {
				t.Log("The match should be started [OK]")
			} else {
				t.Error("The match should be started [ERROR]")
			}

		}

	}
}

func TestWhenMatchStartsTeamsShouldBeSetSuccesfully(t *testing.T) {
	consoleCmd := "Start:'England' vs. 'West Germany'"
	t.Log("Given the command: ", consoleCmd)
	{
		t.Log("When the command is executed into the dashboard")
		{
			matchDashboard := dashboard.New()
			adp := NewAdapter(matchDashboard)
			adp.Execute(consoleCmd)
			if matchDashboard.GetHomeTeam() == "England" && matchDashboard.GetAwayTeam() == "West Germany" {
				t.Log("Home Team should be 'England' and Away Team should be 'West Germany' [OK]")
			} else {
				t.Errorf("Home Team should be 'England' and Away Team should be 'West Germany but they are %s and %s'[ERROR]",
					matchDashboard.GetHomeTeam(),
					matchDashboard.GetAwayTeam())
			}
		}
	}
}

func TestWhenMatchIsStartedAGoalCanBeAdded(t *testing.T) {
	consoleStartCmd := "Start:'England' vs. 'West Germany'"
	consoleAddGoalCmd := "11 'West Germany' Haller"
	t.Log("Given the command: ", consoleAddGoalCmd)
	{
		t.Log("When the command is executed into a dashboard with an already started match")
		{
			matchDashboard := dashboard.New()
			adp := NewAdapter(matchDashboard)
			adp.Execute(consoleStartCmd)
			adp.Execute(consoleAddGoalCmd)
			if matchDashboard.GetHomeTeamScore() == 0 && matchDashboard.GetAwayTeamScore() == 1 {
				t.Log("Home Team score should be '0' and Away Team score should be '1' [OK]")
			} else {
				t.Errorf("Home Team score should be '0' and Away Team score should be '1' but they are %d and %d'[ERROR]",
					matchDashboard.GetHomeTeamScore(),
					matchDashboard.GetAwayTeamScore())
			}
		}
	}
}

func TestWhenGoalIsScoredInfoIsSetSuccesfully(t *testing.T) {
	consoleStartCmd := "Start:'England' vs. 'West Germany'"
	consoleAddGoalCmd := "11 'West Germany' Haller"
	t.Log("Given the command: ", consoleAddGoalCmd)
	{
		t.Log("When the command is executed into a dashboard with an already started match")
		{
			matchDashboard := dashboard.New()
			adp := NewAdapter(matchDashboard)
			adp.Execute(consoleStartCmd)
			adp.Execute(consoleAddGoalCmd)
			goals := matchDashboard.GetAwayGoals()
			expectedPlayer := "Haller"
			expectedMinute := 11
			expectedTeam := "West Germany"
			if len(goals) != 1 {
				//error
				t.Errorf("Number of goals should be %d but is %d [ERROR]", 1, len(goals))
			} else {
				if goals[0].Player != expectedPlayer ||
					goals[0].Minute != expectedMinute ||
					goals[0].Team != expectedTeam {
					//error
					t.Errorf("Goal information should be {%s, %d, %s} but is {{%s, %d, %s} [ERROR]",
						expectedPlayer, expectedMinute, expectedTeam,
						goals[0].Player, goals[0].Minute, goals[0].Team)
				} else {
					t.Logf("Goal information should be {%s, %d, %s} [OK]", expectedPlayer, expectedMinute, expectedTeam)
				}
			}
		}
	}
}

func TestWhenTwoGoalsAreScoredInfoIsSetSuccesfully(t *testing.T) {
	consoleStartCmd := "Start:'England' vs. 'West Germany'"
	consoleAddFirstGoalCmd := "11 'West Germany' Haller"
	consoleAddSecondGoalCmd := "13 'England' Charlton"
	t.Log("Given the commands: ", consoleAddFirstGoalCmd, "and", consoleAddSecondGoalCmd)
	{
		t.Log("When the commands are executed into a dashboard with an already started match")
		{
			matchDashboard := dashboard.New()
			adp := NewAdapter(matchDashboard)
			adp.Execute(consoleStartCmd)
			adp.Execute(consoleAddFirstGoalCmd)
			adp.Execute(consoleAddSecondGoalCmd)
			homeGoals := matchDashboard.GetHomeGoals()
			awayGoals := matchDashboard.GetAwayGoals()
			homeScore := matchDashboard.GetHomeTeamScore()
			awayScore := matchDashboard.GetAwayTeamScore()
			if homeScore != 1 || awayScore != 1 || len(homeGoals) != 1 || len(awayGoals) != 1 {
				t.Errorf("Match score should be %d-%d and number of goals should be %d but result is %d-%d and number of goals is %d [ERROR]",
					1, 1, 2, homeScore, awayScore, len(homeGoals)+len(awayGoals))
			} else {
				t.Logf("Match score should be %d-%d and number of goals should be %d [OK]",
					1, 1, 2)
			}
		}
	}
}

func TestPrintSucesfullyPrintsMatchInformation(t *testing.T) {
	consoleStartCmd := "Start:'England' vs. 'West Germany'"
	consoleAddFirstGoalCmd := "11 'West Germany' Haller"
	consoleAddSecondGoalCmd := "13 'England' Charlton"
	consoleAddThirdGoalCmd := "87 'England' Max"
	consolePrintCommand := "print"
	t.Log("Given the command: ", consolePrintCommand)
	{
		t.Log("When the command is executed into a dashboard with an already started match")
		{
			matchDashboard := dashboard.New()
			adp := NewAdapter(matchDashboard)
			adp.Execute(consoleStartCmd)
			adp.Execute(consoleAddFirstGoalCmd)
			adp.Execute(consoleAddSecondGoalCmd)
			adp.Execute(consoleAddThirdGoalCmd)
			adp.Execute(consolePrintCommand)
			timeout := time.After(500 * time.Millisecond) //give some time to the goroutine to send the info into the channel
			select {
			case message := <-adp.Message: //Print command should put a message into adp.Message channel
				if message == "England 2 (Charlton 13' Max 87') vs. West Germany 1 (Haller 11')" {
					t.Logf("Score information should be printed successfully [OK]")
				} else {
					t.Errorf("Score information should be printed successfully, but was %s [Error]", message)
				}
			case <-timeout:
				t.Errorf("Score information should be printed before given time [Error]")
			}
		}
	}
}

// Add a goal with wrong teams
// Print command where the same player scores more than one goal
