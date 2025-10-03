package main

import (
	"first_day/models"
	"first_day/utils"
	"fmt"
)

func main() {

	// * TRIAL DATA
	tries := 2
	spoiled_votes := 0
	exit := false

	cand_1 := models.Candidate{Name: "modi", Party: "BJP", Votes: 0}
	cand_2 := models.Candidate{Name: "micks", Party: "GTR", Votes: 0}

	candidates := []models.Candidate{cand_1, cand_2}
	voters := []models.Person{}

	for !exit {
		registered_voter := utils.RegisterVoter(voters)

		// * If statemement

		utils.IFHelper(registered_voter.Age, candidates, tries, spoiled_votes, registered_voter)

		fmt.Println("Do you want to exit? (yes/no)")
		var response string
		fmt.Scan(&response)

		if response == "yes" {
			exit = true

			//  Final results
			fmt.Println("\n======= FINAL VOTE COUNT =======")

			for _, candidate := range candidates {
				fmt.Printf("\nName: %s | Party: %s | Votes: %d\n", candidate.Name, candidate.Party, candidate.Votes)

				if len(candidate.Voters) == 0 {
					fmt.Println("  No voters recorded for this candidate.")
				} else {
					fmt.Println("  Voters:")
					for _, voter := range candidate.Voters {
						fmt.Printf("   - %s (Age: %d, ID: %d)\n", voter.Name, voter.Age, voter.IDNumber)
					}
				}
			}

			fmt.Println("\nTotal Spoiled Votes:", spoiled_votes)
		} else {
			exit = false
		}
	}
}
