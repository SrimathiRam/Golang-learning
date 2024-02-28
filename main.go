package main

import "fmt"

func main() {

	var conferenceName = "Go conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	var booking[] string

	fmt.Printf(" Welcome to our %v booking application !!!\n", conferenceName)
	fmt.Printf("  we have total of %v tickets and %v  are still available ....\n", conferenceTickets, remainingTickets)
	fmt.Println("   Get your tickets here to attend....")

	var firstName string
	var lastName string
	var email string
	var userTicket int

	fmt.Println("Enter your First Name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your Last Name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter Email ID:")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets :")
	fmt.Scan(&userTicket)

	remainingTickets = remainingTickets - userTicket
        booking=append(booking,firstName+" "+lastName)


	fmt.Printf("The Whole slices:%v \n",booking)
	fmt.Printf("The First value:%v \n",booking[0])
	fmt.Printf("Slices type:%T \n",booking)
	fmt.Printf("Slices length:%v \n",len(booking))


	fmt.Printf(" User %v%v boooked %v tickets...\n", firstName, lastName, userTicket)

	fmt.Printf("       Thank you  for booking....\nYou will receive a confirmation email at %v", email)
	fmt.Printf(" %v tickets remaing for %v\n", remainingTickets, conferenceName)
}
