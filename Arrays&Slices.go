package main

import (
	"fmt"
	"strings"
	"GoLangOp/helper"
)
func main() {
	fmt.Println("Arrays and Slices");

	var bookings = [50]string{}  //50 is size of the array and its an empty array
	// or 
	// var bookings [50]string 


	var firstName string
	var lastName string

	fmt.Scan(&firstName)
	fmt.Scan(&lastName)

	bookings[0] = firstName + " " + lastName 
	// bookings[1] = "Nicole"
	// bookings[2] = "Peter"
	fmt.Println(bookings)
	fmt.Println(bookings[0])
	fmt.Println(len(bookings)) //size of an array
	fmt.Printf("Array is of type: %T\n", bookings) //Type of an array

	// Problem is we have to specify the size of array while decalring it.
	// Use a list that is more Dynamic in Size where we dont need to specify the size at the beginning
	//We will use Slices it is an abstraction of array
	var bookings1 []string   // Slice
	var remainingTickets int = 50
	var firstName1 string
	var lastName1 string
	var email string 
	var userTickets int

	fmt.Scan(&firstName1)
	fmt.Scan(&lastName1)
	fmt.Scan(&email)
	fmt.Scan(&userTickets)

	isValidName := len(firstName1) >= 2 && len(lastName1) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidUserTickets := userTickets > 0 && userTickets <= remainingTickets
	// isInValidCity := city != "singapore" || city != "London" //If city is not singapore and London the true.

	
	if isValidName && isValidEmail && isValidUserTickets == true {

	bookings1 = append(bookings1, firstName1 + " " + lastName1) //Append adds the new element at the end of list

	fmt.Println(bookings1)
	fmt.Println(bookings1[0])
	fmt.Println(len(bookings1)) //size of an Slice
	fmt.Printf("Slice is of type: %T\n", bookings1) //Type of the Slice

	// // For Loop -> Infinite Loop
	// for {
	// 	var firstName2 string
	// 	var lastName2 string
	
	// 	fmt.Scan(&firstName2)
	// 	fmt.Scan(&lastName2)
	
	// 	var bookings2 []string 
	// 	bookings2 = append(bookings2, firstName2 + " " + lastName2) //Append adds the new element at the end of list
	
	// 	fmt.Println(bookings2)
	// 	fmt.Println(bookings2[0])
	// 	fmt.Println(len(bookings2)) //size of an Slice
	// 	fmt.Printf("Slice is of type: %T\n", bookings2) //Type of the Slice
	// }

	// for-each loop -> Iterating over a list
		// Assume that from each slice we want only to print the firstname and not the lastname 

	firstNames := []string{}
	for _ , booking := range bookings1 {// _ (underscore) is a blank identifier used to ignore a variable we dont want to use.
		var names = strings.Fields(booking)
		var firstName3 = names[0]
		// var lastName3 = names[1]
		firstNames = append(firstNames, firstName3)

	}
	fmt.Printf("The first names of bookings are: %v\n", firstNames)
	fmt.Println(bookings1)

// Range iterates over elements for different data structures, for arrays and slices range provides the index and
// value for each element.
// strings.Fields() splits the string with white spaces as seperator.
// and returns a slice with the split elements ["Abhijeet dandekar"] => ["Abhijeet","dandekar"]

	} else {
		fmt.Println("Input is Invalid!")
	}

// Functions
	message := helper.HelloWorld("Hello this is function Hello World!")

	fmt.Println(message);

	var arr = []string{"abhijet dandekar", "sarvesh dandekar"}
	arrayMsg := ArrayFunc(arr);
	fmt.Println(arrayMsg)


}


// func HelloWorld(message string) string {
// 	return message;
// }

func ArrayFunc(firstNamesOnly []string) []string {   // parameter is the inout from function call and []string is 
	// return type 
	return firstNamesOnly
}