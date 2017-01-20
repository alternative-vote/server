package altVote

import "errors"
import "fmt"

type Results struct {
	Winner string
	Rounds []map[string]int64
}

var NoVotes = errors.New("No one votes for any of the candidates")

func GetResults(candidates []string, ballots [][]string) (Results, error) {
	fmt.Println("getting results using the coppied library...")
	//set up the bins
	bins := map[string]int64{}
	for _, candidate := range candidates {
		bins[candidate] = 0
	}

	//results will be mutated as runRound gets called
	results := new(Results)
	err := runRound(candidates, bins, ballots, results)
	return *results, err
}

func runRound(candidates []string, bins map[string]int64, ballots [][]string, results *Results) error {
	//get the totals for this round
	for _, ballot := range ballots {
		for _, vote := range ballot {
			_, ok := bins[vote]
			if ok {
				bins[vote]++
				break
			}
		}
	}

	results.Rounds = append(results.Rounds, copyBins(bins))

	//check to see if there's a winner yet
	winner, err := getWinner(bins)
	if err != nil {
		return err
	}

	if winner != "" {
		results.Winner = winner
		return nil
	}

	//looks like we don't have a winner - drop the person in last and re-runRound
	lastPlaceCandidate := getLast(candidates, bins)
	delete(bins, lastPlaceCandidate)
	//reset bins
	for i := range bins {
		bins[i] = 0
	}
	return runRound(candidates, bins, ballots, results)

}

func getWinner(bins map[string]int64) (string, error) {
	var totalBallots float64
	for _, votes := range bins {
		totalBallots += float64(votes)
	}

	if totalBallots == 0 {
		return "", NoVotes
	}

	for candidate, votes := range bins {
		if float64(votes)/totalBallots > .5 {
			return candidate, nil
		}
	}

	return "", nil
}

//need to pass candidates here so we can deterministically get a loser in the event of a tie - lower index loses
func getLast(candidates []string, bins map[string]int64) string {
	var lastCandidate string
	var lastCandidateVotes int64
	for _, candidate := range candidates {
		if _, ok := bins[candidate]; ok { //only bother with this candidate if they are still in the running
			votes := bins[candidate]
			if lastCandidate == "" || votes < lastCandidateVotes {
				lastCandidate = candidate
				lastCandidateVotes = votes
			}
		}
	}
	return lastCandidate
}

func copyBins(bins map[string]int64) map[string]int64 {
	ret := map[string]int64{}
	for k, v := range bins {
		ret[k] = v
	}
	return ret
}
