package main

import (
	// "errors"
	"fmt"
	// "time" // work with date and time value
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user

	appUser, err := newUser(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	// appUser = user{ // instantiating
	// 	firstName: userFirstName,
	// 	lastName:  userLastName,
	// 	birthDate: userBirthdate,
	// 	createdAt: time.Now(),
	// }

	// ... do something awesome with that gathered data!

	// outputUserDetails(&appUser) // passing pointer

	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()
}

// func outputUserDetails(u *user) { // dereference

// 	fmt.Println(u.firstName, u.lastName, u.birthDate)
// }

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
