// It can hold mixed data-types.

package main

import (
	"fmt"
	// "strings"
)

var bookings = make([]UserData, 0)  // Initiates empty list of userData struct

type UserData struct {   //type keyword creates a new type, with the name you specify
	firstName string 
	lastName string
	email string 
	numberOfTickets uint
}



func main()  {
	var userData = UserData {
		firstName: "Abhijeet",
		lastName: "Dandekar",
		email: "ab@gmail.com",
		numberOfTickets: 10,
	}

	bookings = append(bookings,userData)
	fmt.Println(userData.firstName)
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	fmt.Println(firstNames)

}



