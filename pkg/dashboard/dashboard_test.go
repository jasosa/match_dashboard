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

func TestStartsAlreadyStartedMatch(t *testing.T) {
	scoringMatch := New()
	scoringMatch.Start("home", "away")
	err := scoringMatch.Start("Juve", "Madrid")
	if err == nil {
		t.Error("Start an already started match should return an error [ERROR]")
	} else if err != ErrMatchAlreadyStarted {
		t.Error("Start an already started match should return an error of type 'ErrMatchAlreadyStarted' [ERROR]")
	} else {
		t.Log("Start an already started match should return an error [OK]")
	}
}

func TestEndsSuccesfully(t *testing.T) {
	scoringMatch := New()
	scoringMatch.Start("home", "away")
	err := scoringMatch.End()
	if err != nil {
		t.Errorf("Ending an already started match should end the game succesfully  but returned error '%s' [Error]",
			err.Error())
	} else if scoringMatch.GetHomeTeam() != "" ||
		scoringMatch.GetAwayTeam() != "" ||
		len(scoringMatch.GetHomeGoals()) != 0 ||
		len(scoringMatch.GetAwayGoals()) != 0 {
		t.Error("Ending an started match should end the game succesfully and reset all the fields [Error]")
	} else {
		t.Log("Ending an started match should end the game succesfully [OK]")
	}
}

func TestEndsANonStartedMatch(t *testing.T) {
	scoringMatch := New()
	err := scoringMatch.End()
	if err == nil {
		t.Error("Ending an non-started match should return an error [ERROR]")
	} else if err != ErrMatchNotStarted {
		t.Error("Ending an non-started match should return an error of type 'ErrMatchNotStarted' [ERROR]")
	} else {
		t.Log("Ending an non-started match should return an error [OK]")
	}
}

func TestAddGoalInANonStartedMatch(t *testing.T) {
	scoringMatch := New()
	err := scoringMatch.AddGoal(35, "Juventus", "Dybala")
	if err == nil {
		t.Error("Adding a goal in an non-started match should return an error [ERROR]")
	} else if err != ErrMatchNotStarted {
		t.Error("Adding a goal in an non-started match should return an error of type 'ErrMatchNotStarted' [ERROR]")
	} else {
		t.Log("Adding a goal in an non-started match should return an error [OK]")
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
	scoringMatch.AddGoal(13, "Italy", "Vieri")
	homeScore := scoringMatch.GetHomeTeamScore()
	awayScore := scoringMatch.GetAwayTeamScore()
	homeGoals := scoringMatch.GetHomeGoals()
	awayGoals := scoringMatch.GetAwayGoals()
	if homeScore != 1 || awayScore != 1 {
		t.Errorf("Expected score was  %d-%d but returned is %d-%d [ERROR]",
			1, 1, homeScore, awayScore)
	} else if len(homeGoals) != 1 || len(awayGoals) != 1 {
		t.Errorf("Expected number of goals was %d and %d but returned is %d and %d [ERROR]",
			1, 1, len(homeGoals), len(awayGoals))
	} else if awayGoals[0].Player != "Iniesta" || awayGoals[0].Minute != 11 || awayGoals[0].Team != "Spain" {
		t.Errorf("Expected away goals info was %s, %d, %s but returned is %s, %d, %s [ERROR]",
			"Iniesta", 11, "Spain", awayGoals[0].Player, awayGoals[0].Minute, awayGoals[0].Team)
	} else {
		t.Logf("Expected away goals info was %s, %d, %s [OK] ", "Iniesta", 11, "Spain")
	}
}
