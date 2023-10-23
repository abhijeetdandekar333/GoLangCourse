package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"math/rand"
	"time"
)

type deck []string
// We can use this deck anywhere to initiate a slice

//Receiver function
func (d deck) print() {   //Receiver function, Here function name is "print()" and (d deck) is a receiver.
	for i, card := range d {
		fmt.Println(i, card)
	}
}
// func (d deck) print() -> Any variable of type deck now gets access to the print method.
// d-> actual copy of deck we are working with is available in function as a variable called d.
// Every variable of type deck can call this function on itself.

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades","Hearts","Diamond","Clubs"}
	cardValues := []string{"Ace","One","Two","Three","Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues{
			cards = append(cards,value+ " of " + suit + " ") 
		}
	}
	return cards

}

// Multiple Return Values- Returns two values of type deck
func deal(d deck, handSize int) (deck,deck) {
	return d[:handSize], d[handSize:]
}

// Saving file to local machine hard drive using ioutil package
// Byte Slices -> [72 105 32] -> find the character for each byte in table e.g. 72 is H, 105 is i and 32 is a space
// Its computer friendly way of reperesenting a string.
// Deck to String(Type Conversion)
// Deck is a slice of string so first convert it into string and then to []byte

// Join Concatenates the elements of a slice to create a single string. 
func (d deck) toString() string {
	return strings.Join([]string(d), ",") //Takes the slice of strings and joins it with comma in seperate string.	
}

// Creates a file on Local Machine
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)   // 0666 is permission to file for everyone to write
}


// Create a Deck from a file
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)  //ReadFile returns byte slice and error(if nothing went wrong then this will have nil value)

	if err != nil {
		// Log the error 
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// Split string into the substrings and returns a slice of substrings  
	s := strings.Split(string(bs), ",") //Convert byte to string first and split it into slice of string using Split
	return deck(s)
}

// Shuffling a Deck
func (d deck) shuffle()  {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i] //swap elements 
	}
}


// Testing with Go
// To make a test, create a new file ending in _test.go
// To test run command go test. check deck_test.go file



