package main

import (
	"fmt"
	"time"
)

type UserDetails struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

// Receiver argument
func (userDetails UserDetails) outputUserDetails() {
	fmt.Printf("%v %v, %v - %v\n", userDetails.firstName, userDetails.lastName, userDetails.birthdate, userDetails.createdAt)
}

func main() {
	firstName := getUserData("Please enter first name: ")
	lastName := getUserData("Please enter last name: ")
	birthdate := getUserData("Please enter birth date (MM/DD/YYYY): ")

	var userDetails UserDetails

	userDetails = UserDetails{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}

	// userDetails.firstName = firstName
	// userDetails.lastName = lastName
	// userDetails.birthdate = birthdate
	// userDetails.createdAt = time.Now()

	userDetails.outputUserDetails()
}

func getUserData(label string) string {
	fmt.Print(label)
	var value string
	fmt.Scan(&value)
	return value
}
