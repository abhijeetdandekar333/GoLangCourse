package main

// import "fmt"

func main()  {
	// cards := deck{"Abhijeet Dandekar", newCard()}
	// cards = append(cards,"Sanket Dandekar")
	// cards.print()  //when print is executed "cards" variable gets passed into our print function as a variable called d.
	// var cards = newDeck()
	// fmt.Println(cards)

	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()

	// greeting := "Hi there"
	// fmt.Println([]byte(greeting)) //Converts string into bytes

	// cards := newDeck()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_cards");
	// cards.print();

	cards := newDeck()
	cards.shuffle()
	cards.print()
}

// func newCard() string {
// 	return "Sarvesh Dandekar"
// }

// Two ways of calling a function 
// 1. By receiver -> cards.print() function
// 2. By passing parameters 