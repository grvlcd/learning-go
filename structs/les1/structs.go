package main

import (
	"example.com/les1/user"
	"fmt"
)

func main() {
	firstName := getUserData("Please enter first name: ")
	lastName := getUserData("Please enter last name: ")
	birthdate := getUserData("Please enter birth date (MM/DD/YYYY): ")

	var userDetails *user.UserDetails

	userDetails, err := user.NewUserDetails(
		firstName,
		lastName,
		birthdate,
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	// userDetails.firstName = firstName
	// userDetails.lastName = lastName
	// userDetails.birthdate = birthdate
	// userDetails.createdAt = time.Now()

	userDetails.OutputUserDetails()
	userDetails.ClearUserName()
	userDetails.OutputUserDetails()
}

func getUserData(label string) string {
	fmt.Print(label)
	var value string
	fmt.Scanln(&value)
	return value
}
