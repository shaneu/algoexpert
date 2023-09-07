package tournamentwinner

const HOME_TEAM_WON = 1

func TournamentWinner(competitions [][]string, results []int) string {
	teamResults := make(map[string]int)
	curLeader := ""

	for i, comp := range competitions {
		result := results[i]
		homeTeam, awayTeam := comp[0], comp[1]
		var roundWinner string

		if result == HOME_TEAM_WON {
			roundWinner = homeTeam
		} else {
			roundWinner = awayTeam
		}

		_, ok := teamResults[roundWinner]
		if !ok {
			teamResults[roundWinner] = 0
		}

		teamResults[roundWinner] += 3

		if curLeader == "" {
			curLeader = roundWinner
		}

		if teamResults[roundWinner] > teamResults[curLeader] {
			curLeader = roundWinner
		}
	}

	return curLeader
}
