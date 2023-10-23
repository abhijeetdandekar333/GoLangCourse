package main

import (
	"fmt"
	"time"
	"sync"
)

// go keyword: Starts a new goroutine. A Goroutine is a lightweight thread managed by Go runtime.
// By default the main goroutine(thread) does not wait for other goRoutines so as i Have written a infinite for loop
// the code will execute but if I remove the for loop the main thread will exit without executing the printname function
// We need to tell main thread to wait. We will create a WaitGroup
// WaitGroup waits for the launched goRoutine to finish
// Package sync provides basic synchronization functionality
// Add: sets the number of goroutines to wait for
// wg.Add(count) Sets the counter of threads that the application should wait for
// wg.Wait() waits until the WaitGroup counter is 0
// wg.Done() Decrements the WaitGroup by 1 so this is called by goroutine to indicate that its finished.

var wg = sync.WaitGroup{}

func main()  {
	var userName string =  "Chris"
	var lastName string =  "Gayle"
	// var name = go printname(userName, lastName) 
	// for {
	wg.Add(1)
	go printname(userName, lastName)
	// fmt.Println("Output is \n", name)
	// }
	wg.Wait()
}

func printname(username string, lastname string)  {
	time.Sleep(10 * time.Second)   // The sleep functions stops or blocks the current "thread(goroutine) execution for defined duration"
	var name = fmt.Sprintf("Username is %v and LastName is %v\n",username,lastname)
	// return name
	fmt.Println("Output is \n", name)

	wg.Done()
	// Handling this blocking code with Goroutines 
	// We need concurrency to run application in multi-threads.

}