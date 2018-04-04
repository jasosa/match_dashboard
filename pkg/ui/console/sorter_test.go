package console

import (
	"github.com/jasosa/football_scoring_dashboard/pkg/dashboard"
	"sort"
	"testing"
)

func TestSorterByPlayerSortsSuccessfully(t *testing.T) {
	goals := []dashboard.Goal{
		dashboard.Goal{Player: "Iniesta", Minute: 11, Team: "Spain"},
		dashboard.Goal{Player: "Iniesta", Minute: 13, Team: "Spain"},
		dashboard.Goal{Player: "Costa", Minute: 75, Team: "Spain"},
		dashboard.Goal{Player: "Iniesta", Minute: 81, Team: "Spain"},
	}

	sort.Sort(byPlayer(goals))
	if goals[0].Player != "Iniesta" || goals[1].Player != "Iniesta" || goals[2].Player != "Iniesta" {
		t.Errorf("Sort was not working as expected")
	} else {
		t.Log("Sort by players working as expected")
	}
}
