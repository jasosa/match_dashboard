package dashboard

//ScoringMatch represents a match in the dashboard
type ScoringMatch interface {
	Start(home, away string)
	AddGoal(minute int, team, player string)
	IsStarted() bool
	GetHomeTeam() string
	GetAwayTeam() string
	GetHomeTeamScore() int
	GetAwayTeamScore() int
	GetHomeGoals() []Goal
	GetAwayGoals() []Goal
}

type scoringMatch struct {
	isStarted     bool
	homeTeam      string
	awayTeam      string
	homeTeamScore int
	awayTeamScore int
	homeGoals     []Goal
	awayGoals     []Goal
}

//Goal represents the information of a goal
type Goal struct {
	Minute int
	Player string
	Team   string
}

//New new scoring match
func New() ScoringMatch {
	return new(scoringMatch)
}

func (sm *scoringMatch) Start(home, away string) {
	sm.isStarted = true
	sm.homeTeam = home
	sm.awayTeam = away
}

func (sm *scoringMatch) AddGoal(minute int, team, player string) {
	if team == sm.homeTeam {
		sm.homeTeamScore++
		sm.homeGoals = append(sm.homeGoals, Goal{Minute: minute, Team: team, Player: player})
	} else if team == sm.awayTeam {
		sm.awayTeamScore++
		sm.awayGoals = append(sm.awayGoals, Goal{Minute: minute, Team: team, Player: player})
	}
}

func (sm *scoringMatch) IsStarted() bool {
	return sm.isStarted
}

func (sm *scoringMatch) GetHomeTeam() string {
	return sm.homeTeam
}

func (sm *scoringMatch) GetAwayTeam() string {
	return sm.awayTeam
}

func (sm *scoringMatch) GetHomeTeamScore() int {
	return sm.homeTeamScore
}

func (sm *scoringMatch) GetAwayTeamScore() int {
	return sm.awayTeamScore
}

func (sm *scoringMatch) GetHomeGoals() []Goal {
	return sm.homeGoals
}

func (sm *scoringMatch) GetAwayGoals() []Goal {
	return sm.awayGoals
}
