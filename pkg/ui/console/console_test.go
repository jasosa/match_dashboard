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
			adp.Execute(consoleCmd, true)
			<-adp.Message
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
			adp.Execute(consoleCmd, true)
			<-adp.Message
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
			adp.Execute(consoleStartCmd, false)
			adp.Execute(consoleAddGoalCmd, false)
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
			adp.Execute(consoleStartCmd, false)
			adp.Execute(consoleAddGoalCmd, false)
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
			adp.Execute(consoleStartCmd, false)
			adp.Execute(consoleAddFirstGoalCmd, false)
			adp.Execute(consoleAddSecondGoalCmd, false)
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
	consoleAddFourthGoalCmd := "91 'England' Charlton"
	consolePrintCommand := "Print"
	t.Log("Given the command: ", consolePrintCommand)
	{
		t.Log("When the command is executed into a dashboard with an already started match")
		{
			matchDashboard := dashboard.New()
			adp := NewAdapter(matchDashboard)
			adp.Execute(consoleStartCmd, false)
			adp.Execute(consoleAddFirstGoalCmd, false)
			adp.Execute(consoleAddSecondGoalCmd, false)
			adp.Execute(consoleAddThirdGoalCmd, false)
			adp.Execute(consoleAddFourthGoalCmd, false)
			adp.Execute(consolePrintCommand, true)
			timeout := time.After(500 * time.Millisecond) //give some time to the goroutine to send the info into the channel
			select {
			case message := <-adp.Message: //Print command should put a message into adp.Message channel
				if message == "England 3 (Charlton 13' 91' Max 87') vs. West Germany 1 (Haller 11')" {
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

func TestCommandWhileMatchNotInProgress(t *testing.T) {
	expectedMessage := "No game currently in progress"
	consoleAddFirstGoalCmd := "11 'West Germany' Haller"
	t.Log("Given the command: ", consoleAddFirstGoalCmd)
	{
		t.Log("When the command is executed into a dashboard with NO already started match")
		{
			matchDashboard := dashboard.New()
			adp := NewAdapter(matchDashboard)
			adp.Execute(consoleAddFirstGoalCmd, true)
			timeout := time.After(500 * time.Millisecond)
			select {
			case message := <-adp.Message:
				if message != expectedMessage {
					t.Errorf("The message '%s' should be sent but the message '%s' is received [Error]", expectedMessage, message)
				} else {
					t.Logf("The message '%s' should be received [OK]", expectedMessage)
				}
			case <-timeout:
				t.Errorf("The message '%s' should be sent but any message is received [Error]", expectedMessage)
			}
		}
	}
}

func TestStartAGameWhileCurrentOneIsStarted(t *testing.T) {
	expectedMessage := "Game currently in progress - to start a new game finish the current one  through typing 'End'"
	consoleStartFirstCmd := "Start:'England' vs. 'West Germany'"
	consoleStartSecondCmd := "Start:'Italy' vs. 'Japan'"
	t.Log("Given the command: ", consoleStartSecondCmd)
	{
		t.Log("When the command is executed into a dashboard with an already started match")
		{
			matchDashboard := dashboard.New()
			adp := NewAdapter(matchDashboard)
			adp.Execute(consoleStartFirstCmd, false)
			adp.Execute(consoleStartSecondCmd, true)
			timeout := time.After(500 * time.Millisecond)
			select {
			case message := <-adp.Message:
				if message != expectedMessage {
					t.Errorf("The message '%s' should be sent but the message '%s' is received [Error]", expectedMessage, message)
				} else {
					t.Logf("The message '%s' should be received [OK]", expectedMessage)
				}
			case <-timeout:
				t.Errorf("The message '%s' should be sent but any message is received [Error]", expectedMessage)
			}
		}
	}
}

func TestEndMatchSuccesfully(t *testing.T) {
	consoleStartCmd := "Start:'England' vs. 'West Germany'"
	consoleEndCmd := "End"
	t.Log("Given the command: ", consoleEndCmd)
	{
		t.Log("When the command is executed into a dashboard with an already started match")
		{
			matchDashboard := dashboard.New()
			adp := NewAdapter(matchDashboard)
			adp.Execute(consoleStartCmd, false)
			adp.Execute(consoleEndCmd, false)
			if matchDashboard.IsStarted() {
				t.Error("Match shouldn't be started anymore [Error]")
			} else {
				t.Log("Match shouldn't be started anymore [OK]")
			}
		}
	}
}

func TestEndMatchSuccesfullySendsRightMessage(t *testing.T) {
	expectedMessage := "Match ended succesfully"
	consoleStartCmd := "Start:'England' vs. 'West Germany'"
	consoleEndCmd := "End"
	t.Log("Given the command: ", consoleEndCmd)
	{
		t.Log("When the command is executed into a dashboard with an already started match")
		{
			matchDashboard := dashboard.New()
			adp := NewAdapter(matchDashboard)
			adp.Execute(consoleStartCmd, false)
			adp.Execute(consoleEndCmd, true)
			timeout := time.After(500 * time.Millisecond)
			select {
			case message := <-adp.Message:
				if message != expectedMessage {
					t.Errorf("The message '%s' should be sent but the message '%s' is received [Error]", expectedMessage, message)
				} else {
					t.Logf("The message '%s' should be received [OK]", expectedMessage)
				}
			case <-timeout:
				t.Errorf("The message '%s' should be sent but any message is received [Error]", expectedMessage)
			}
		}
	}
}

// End command
// Errors
// Add a goal with wrong teams
// Non recognized command
// Commands in wrong order
// Before start
