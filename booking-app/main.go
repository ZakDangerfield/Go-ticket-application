package main

import (
	"booking-app/validation"
	"fmt"
	"sync"
	"time"
)

const confTickets = 50

var confName = "My Go Learning"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName   string
	secondName  string
	email       string
	userTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	// Ticket booking application:
	// - Takes Booking information from user
	// - Subtracts selected amount of tickets from remainingTickets variable
	// - Stores Booking information into into expandable
	greetUsers()

	//Collecting booking information

	firstName, secondName, email, userTickets := GetUserInput()
	isValidName, isValidEmail, isValidTickets := validation.ValidateUserInput(firstName, secondName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTickets {

		bookTicket(userTickets, firstName, secondName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, secondName, email)

		firstNames := GetFirstNames()
		fmt.Println("The first names of the bookings are", firstNames)

		if remainingTickets == 0 {
			// end program
			fmt.Println("All tickets sold out, come back next year")

		}
	} else {
		if !isValidName {
			fmt.Println("First Name or last Name was too short")
		}
		if !isValidEmail {
			fmt.Println("Your email address is invalid")
		}
		if !isValidTickets {
			fmt.Println("The number of tickets you entered is invalid, please try again")
		}
	}
	wg.Wait()
}

func greetUsers() {
	// Greeting user message.
	fmt.Println("Welcome to my", confName, "booking application")
	fmt.Println("we have a total of", confTickets, "tickets. There are", remainingTickets, "tickets remaining")
	fmt.Println("Get tickets here")
}

func GetFirstNames() []string {
	// Gets the first name of each booking.
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func GetUserInput() (string, string, string, uint) {
	var firstName string
	var secondName string
	var email string
	var userTickets uint
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name:")
	fmt.Scan(&secondName)
	fmt.Println("Enter your Email address:")
	fmt.Scan(&email)
	fmt.Println("how many tickets would you like to book:")
	fmt.Scan(&userTickets)

	return firstName, secondName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, secondName string, email string) {
	remainingTickets = remainingTickets - userTickets
	//creating a map for user
	var userData = UserData{
		firstName:   firstName,
		secondName:  secondName,
		email:       email,
		userTickets: userTickets,
	}
	bookings = append(bookings, userData)
	fmt.Println("list of bookings is", bookings)

	fmt.Println("Thank you", firstName, secondName, "for booking", userTickets, "tickets. An e-receipt will be sent to", email, "shortly")
	fmt.Println("tickets remaining: ", remainingTickets)
}

func sendTicket(userTickets uint, firstName string, secondName string, email string) {
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, secondName)
	time.Sleep(10 * time.Second)
	fmt.Println("#####################")
	fmt.Println("Sending ticket\n", ticket, "\nto email address:\n", email)
	fmt.Println("#####################")
	wg.Done()
}
