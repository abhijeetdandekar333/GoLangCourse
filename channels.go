// USed when we want to transfer data between different goRoutines.  

package main

import (
	"fmt"
	"time"
	"math/rand"
	"sync"
)



func DoWork() int  {
	time.Sleep(time.Second)
	return rand.Intn(100)
}


func main()  {
	dataChan := make(chan int)
	// dataChan <- 789  // Gives deadlock error we need to execute it with n:= <-dataChan at same time so we will put it in
	// another GoRoutine
	// <- shows that we are putting value 789 into dataChan and n := <- dataChan shows we are putting values
	// from dataChan to variable n
	// Basically two threads in parallel one is for adding the data into channel and other is for receiving 
	// data from channel

	go func ()  {
		dataChan <- 789
	}()
	n := <- dataChan
	fmt.Printf("Value inside data channel is %d\n", n)

	// If we want to remove deadlock error we can initiate the channel with some memory
	dataChan1 := make(chan int, 1) //buffered channel
	dataChan1 <- 10
	m := <- dataChan1 
	fmt.Printf("Value inside data channel 1 is %d\n", m)

	// Channels are meant to be used in a multi threaded context.



	dataChan3 := make(chan int)
	
	go func() {                      //Runs on Background thread
		wg := sync.WaitGroup{}
		for i := 0; i < 10; i++ {
			wg.Add(1);
			go func() {             //Every iteration will have a different Thread. so iteration will be done in parallel in less time.
				defer wg.Done()
				result := DoWork()   
				dataChan3 <- result

			}()
		}
		wg.Wait()
		close(dataChan3)  //Avoids values of i from going out of channel by closing it.
	}()

	for p := range dataChan3 {       // Runs on Main Thread
		fmt.Printf("p = %d\n",p)
	}
}
