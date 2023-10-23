// Stores data key value pairs

package main 

import (
	"fmt"
	// "strings"
	"strconv"
)

func main(){
	
	// var userData = map[string]string    //[string] -> Type of key, string-> type of value
	// Create an empty map
	var userTickets uint = 40
	var userData = make(map[string]string)
	userData["name"] = "Abhijeet"   
	userData["lastname"] = "Dandekar"   
	userData["email"] = "abhijett@gmail.com"  
	userData["NumberofTickets"] = strconv.FormatUint(uint64(userTickets),10)

	var bookings = make([]map[string]string, 0) //Empty slice using map i.e. the empty list of maps with initial sizeof 0

	bookings = append(bookings,userData)
	// fmt.Println(userData); 
	fmt.Println(bookings); 
	fmt.Printf("%v",bookings["name"]); 
}