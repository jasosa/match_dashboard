package dashboard

import (
	"testing"
)

func TestStartsSuccesfully(t *testing.T) {
	scoringMatch := New()
	scoringMatch.Start("home", "away")
	if !scoringMatch.IsStarted() {
		t.Error("Scoring Match should be started [ERROR]")
	} else {
		t.Log("Scoring Match should be started [OK]")
	}
}

func TestStartsCreateTeamsSuccesfully(t *testing.T) {
	scoringMatch := New()
	scoringMatch.Start("Italy", "Spain")
	homeTeam := scoringMatch.GetHomeTeam()
	awayTeam := scoringMatch.GetAwayTeam()
	if homeTeam != "Italy" || awayTeam != "Spain" {
		t.Errorf("Expected home and away teams were %s and %s but returned are %s and %s [ERROR]",
			"Italy", "Spain", homeTeam, awayTeam)
	} else {
		t.Logf("Expected home and away teams were %s and %s [OK]", homeTeam, awayTeam)
	}
}

func TestAddGoalSetInfoSuccesfully(t *testing.T) {
	scoringMatch := New()
	scoringMatch.Start("Italy", "Spain")
	scoringMatch.AddGoal(11, "Spain", "Iniesta")
	homeScore := scoringMatch.GetHomeTeamScore()
	awayScore := scoringMatch.GetAwayTeamScore()
	goals := scoringMatch.GetAwayGoals()
	if homeScore != 0 || awayScore != 1 {
		t.Errorf("Expected score was  %d-%d but returned is %d-%d [ERROR]",
			0, 1, homeScore, awayScore)
	} else if len(goals) != 1 {
		t.Errorf("Expected number of goals was %d but returned is %d [ERROR]",
			1, len(goals))
	} else if goals[0].Player != "Iniesta" || goals[0].Minute != 11 || goals[0].Team != "Spain" {
		t.Errorf("Expected goals info was %s, %d, %s but returned is %s, %d, %s [ERROR]",
			"Iniesta", 11, "Spain", goals[0].Player, goals[0].Minute, goals[0].Team)
	} else {
		t.Logf("Expected goals info was %s, %d, %s [OK] ", "Iniesta", 11, "Spain")
	}
}
