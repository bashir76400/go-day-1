package utils

import (
	"first_day/models"
	"fmt"
)

func RegisterVoter(voters []models.Person) models.Person {

	fmt.Println("---------------------------------------------------")
	fmt.Println("REGISTER FOR VOTING")

	var name string
	var idNo int
	var age int

	fmt.Println("Enter name: ")
	fmt.Scan(&name)

	fmt.Println("Enter ID NUMBER: ")
	fmt.Scan(&idNo)

	fmt.Println("Enter AGE: ")
	fmt.Scan(&age)

	voter := models.Person{Name: name, IDNumber: idNo, Age: age}
	voters = append(voters, voter)

	return voter
}
