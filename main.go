package main

import "fmt" //this package provided input output processing functionality

func main() {

	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50 // Const keyword mean this value is not allowed to change
	var remainingTickets int = 50

	fmt.Printf("conferenceName is %T,conferenceTickets is %T and remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)
	fmt.Printf("Welcome to %v Booking Application\n", conferenceName)
	fmt.Printf("We have to tatal of %v tickets and %v are still available \n", conferenceTickets, remainingTickets)
	fmt.Println("Get Your Tickets Here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets int
	//ask user for their name

	fmt.Println("Enter your First Name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your Last Name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your Email Address")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation at %v\n", firstName, lastName, userTickets, email)

}
