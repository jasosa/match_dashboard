/* The app should print out a scoring dashboard as text during
a football match.
Ex: Match between England and West Germany in the 80th minute
in world cup final in 1966:
"England 2 (Hurst 18' Peters 78') vs West Germany 1 (Haller 12')"
*/

package main

import (
	"bufio"
	"fmt"
	"github.com/jasosa/football_scoring_dashboard/pkg/dashboard"
	"github.com/jasosa/football_scoring_dashboard/pkg/ui/console"
	"os"
)

func main() {
	scoringMatch := dashboard.New()
	adp := console.NewAdapter(scoringMatch)

	fmt.Println("Welcome to Match Dashboard. We are ready for a match")
	fmt.Print(">")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		cmd := scanner.Text()
		adp.Execute(cmd, true)
		fmt.Println(<-adp.Message)
		fmt.Print(">")
	}
}
