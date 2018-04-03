package console

import (
	"github.com/jasosa/football_scoring_dashboard/pkg/dashboard"
)

type sortByPlayer []dashboard.Goal

func (a sortByPlayer) Len() int      { return len(a) }
func (a sortByPlayer) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a sortByPlayer) Less(i, j int) bool {
	if a[i].Player == a[j].Player {
		return a[i].Minute < a[j].Minute
	}

	//check from i to 0 if there was another a[i].Player
	for x := j; x > 0; x-- {
		if a[i].Player == a[x-1].Player {
			return true
		}
	}

	return false
}
