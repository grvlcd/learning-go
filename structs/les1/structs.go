package main

import (
	"example.com/les1/user"
	"fmt"
)

func main() {
	firstName := getUserData("Please enter first name: ")
	lastName := getUserData("Please enter last name: ")
	birthdate := getUserData("Please enter birth date (MM/DD/YYYY): ")

	var appUser *user.User

	appUser, err := user.New(
		firstName,
		lastName,
		birthdate,
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	adminUser := user.NewAdmin("super@admin.com", "password")

	adminUser.OutputUserDetails()
	adminUser.ClearUserName()
	adminUser.OutputUserDetails()

	// userDetails.firstName = firstName
	// userDetails.lastName = lastName
	// userDetails.birthdate = birthdate
	// userDetails.createdAt = time.Now()

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(label string) string {
	fmt.Print(label)
	var value string
	fmt.Scanln(&value)
	return value
}
