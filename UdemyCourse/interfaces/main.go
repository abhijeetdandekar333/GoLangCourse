// Interfaces
// Used to reuse the code inside a function 
// So if we have two functions which contains same logic then we can reuse only one function using interfaces. 
// We can reduce two functions into one in below example
package main

import (
	"fmt"
)

type bot interface {
// This interface expects to see any other type inside code that implements function getGreeting() and return a 
// string and if any type does then that type is considere to be type bot
	getGreeting() string   
}
// Used interfaces to define a method for a function
// Our program now has a new type called 'bot'
// If you are a type in this program with a function called 'getGreeting' and you return a string then
// you are now an honorary member of type 'bot'

// Now that you are an honorary member of type 'bot', you can now call this function called 'printGreeting'


type englishBot struct {}
type spanishBot struct {}

func main()  {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb) 

}

func printGreeting(b bot)  {
	fmt.Println(b.getGreeting())
}

// Merged below two functions using interface 
// func printGreeting(eb englishBot)  {
// 	fmt.Println(eb.getGreeting())
// }  

// func printGreeting(sb spanishBot)  { //Will give error as we cannot use same name for two functions
// 	fmt.Println(sb.getGreeting())
// }

func (eb englishBot) getGreeting()string  { //If you are a type in this program with a function called 'getGreeting' and you return a string then
// you are now an honorary member of type 'bot'
	return "Hi There!"
}


func (sb spanishBot) getGreeting()string  {
	return "Hola!"
}



