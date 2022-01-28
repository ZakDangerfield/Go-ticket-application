package validation

import "strings"

func ValidateUserInput(firstName string, secondName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	//Checking Validity of booking information.
	isValidName := len(firstName) >= 2 && len(secondName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickets := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTickets
}
