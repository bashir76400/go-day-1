package utils

import (
	"first_day/models"
	"fmt"
	"golang.org/x/term"
	"os"
	"strings"
)

func Calling(greeting, name, course string, age int) {
	fmt.Println(greeting, "my name is", name, "and my age is", age, "i am learning", course)
}

func IFHelper(age int, candidates []models.Candidate, tries, spoiled_votes int, voter models.Person) {
	if age > 18 {
		fmt.Println("YOU ARE ELIGIBLE TO VOTE")
		fmt.Println("---------------------------------------------------")
		fmt.Println("WELCOME TO VOTING SYSTEM")
		fmt.Println("List of candidates:")

		for _, candidate := range candidates {
			fmt.Printf("\nName: %s, Party: %s", candidate.Name, candidate.Party)
		}

		fmt.Println("\n---------------------------------------------------")
		fmt.Println("Enter the name of the candidate you want to vote for:")

		// 🔹 Read hidden input
		byteChoice, _ := term.ReadPassword(int(os.Stdin.Fd()))
		choice := strings.TrimSpace(string(byteChoice)) // convert []byte -> string and trim spaces

		fmt.Println("\n(You typed something hidden)") // for UX, don’t reveal choice

		// 🔹 Check candidate existence
		var found bool
		for i := range candidates {
			if strings.EqualFold(candidates[i].Name, choice) { // case-insensitive match
				candidates[i].Votes++
				candidates[i].Voters = append(candidates[i].Voters, voter)
				found = true
				// fmt.Println("✅ Vote recorded for", candidates[i].Name)
				break
			}
		}

		if !found {
			tries--
			if tries > 0 {
				fmt.Println("Candidate not found. You have", tries, "tries left.")
				IFHelper(age, candidates, tries, spoiled_votes, voter) // recursive retry
			} else {
				fmt.Println("❌ You have exhausted all your tries. Spoiled Vote 😒")
				spoiled_votes++
			}
		}
	} else {
		fmt.Println("🚫 You are not eligible to vote")
	}
}
