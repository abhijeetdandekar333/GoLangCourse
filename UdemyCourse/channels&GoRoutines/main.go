package main

import (
	"fmt"
	"net/http"
	"time"

)

// GoRoutines : Allows us to perform actions in concurrent by making use of multiple threads.
// Concurrency is not parallelism.
// We need multiple cores/cpus for parallelism.
// Concurrency : We can have multiple threads executing code. If one thread blocks, another one is picked up
// and worked on
// Parallelism: Multiple threads executed at the exact same time. It requires multiple CPU's. 
// go keyword creates child goroutines.
// Channels: Used for communication between GoRoutines


func main()  {
	links := []string {
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	ch := make(chan string) //Creates a channel and assigns it to ch variable.


	
	for _,link := range links {

		go checkLink(link,ch)//Creates new GoRoutine or engine or thread for this function and run this function with it

	}
	// fmt.Println(<- ch)
	// fmt.Println(<- ch)
	// fmt.Println(<- ch)
	// fmt.Println(<- ch)
	// fmt.Println(<- ch)
	for l := range ch { //for works similar as slice with channels.
		// fmt.Println(<-ch)
		go func(link string) {  // Function literal same as anonymous function(func without name) in js
		// We always pass the argument in function which creates new copy of variable to make sure that main and
		// child GoRoutines use the different memory. 
			time.Sleep(5 * time.Second) //Pauses current GoRoutine for 5 seconds
			checkLink(link,ch)
		}(l) // (l) is Used to call the function and l is the argument that we are passing to the function
	}

}

func checkLink(link string, ch chan string){ //Passed channel parameter along with link
		_, err := http.Get(link)
		if err != nil {
			fmt.Println(link, "might be down!!")
			ch <- link
			return 
		}
		fmt.Println(link, "is up!!")
		ch <- link
}

// channel <- 5 : Send the value "5" into this channel
// myNumber <- channel : Wait for a value to be sent into the channel. When we get one, assign the 
// value to myNumber
// fmt.Println(<-channel) : Wait for the value to be sent into the channel. When we get one, log it out immediately.

