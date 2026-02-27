package user

import (
	"errors"
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
func (userDetails UserDetails) OutputUserDetails() {
	fmt.Printf("%v %v, %v - %v\n", userDetails.firstName, userDetails.lastName, userDetails.birthdate, userDetails.createdAt)
}

func (userDetails *UserDetails) ClearUserName() {
	userDetails.firstName = ""
	userDetails.lastName = ""
}

func New(firstName, lastName, birthdate string) (*UserDetails, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("First name, last name or birthdate is required!")
	}

	return &UserDetails{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}
