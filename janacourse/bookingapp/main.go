package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	fmt.Println("Hello World")
	fmt.Printf("We have a total of %v tickts and %v are still available", conferenceTickets, remainingTickets)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)

	//fmt.Println(conferenceName)
}
