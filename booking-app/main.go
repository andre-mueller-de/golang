package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
	"sync"
	"time"
)

var conferenceName = "Go Conference"
var city = "Berlin"

const conferenceTickets = 50

var remainingTickets uint = 50

//var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	for remainingTickets > 0 && len(bookings) < 50 {

		//var firstName, lastName, email string
		//var userTickets int
		//var city string = "Berlin"
		firstName, lastName, email, userTickets := getUserInput()
		// ask user for nane

		isValidName, isValidEmail, isValidTicketNumber := helper.VaildateUsers(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTickets(userTickets, firstName, lastName, email)

			wg.Add(1)
			go sendTicket(uint(userTickets), firstName, lastName, email)
			noTicketsAvailable := remainingTickets == 0
			if noTicketsAvailable {
				//end program
				fmt.Println("Our conference is booked out now ...")
				break
			}
		} else {
			if !(len(firstName) >= 2 && len(lastName) >= 2) {
				fmt.Println("First or Last Name is too short")
			}
			if !(strings.Contains(email, "@")) {
				fmt.Println("Email address does not contain @ sign")
			}
			continue

		}

	}
	wg.Wait()

}

func greetUsers() {
	fmt.Printf("Welcome to our %v in %v\n", conferenceName, city)
	fmt.Printf(" We have a total of %v tickets and %v are available \n", conferenceTickets, remainingTickets)
	fmt.Println("You can buy your ticket here")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {

		firstNames = append(firstNames, booking.firstName)
	}
	fmt.Println("<<< return from getFirstBanes()")
	return firstNames

}

func getUserInput() (string, string, string, int) {
	var firstName, lastName, email string
	var userTickets int
	fmt.Println("Enter you first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter you last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter you email:")
	fmt.Scan(&email)

	fmt.Println("Number of tickets:")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets int, firstName, lastName, email string) {
	remainingTickets = remainingTickets - uint(userTickets)

	//create a map for a user
	//var userData = make(map[string]string)
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: uint(userTickets),
	}
	//	userData["firstName"] = firstName
	//	userData["lastName"] = lastName
	//	userData["email"] = email
	//	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)

	fmt.Printf("List of bookings: %v\n", bookings)
	//	fmt.Printf("The first value: %v\n", bookings[0])
	//	fmt.Printf("The slice type: %T\n", bookings)
	//	fmt.Printf("The slice length: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for buying %v tickets.  You'll receive a confirmation at: %v \n", firstName, lastName, userTickets, email)

	fmt.Printf("The firstt names  of our bookings: %v\n", getFirstNames())

	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(60 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###############")
	fmt.Printf("Sending ticket\n %v \n to email address: %v\n", ticket, email)
	fmt.Println("###############")
	wg.Done()

}
