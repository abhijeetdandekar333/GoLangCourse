package main 


//  // Program to read content of file and print on terminal

// import (
// 	"io/ioutil"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// type deck []string


// func readFileContent(filename string) deck {
// 	bs, err := ioutil.ReadFile(filename)  //ReadFile returns byte slice and error(if nothing went wrong then this will have nil value)

// 	if err != nil {
// 		// Log the error 
// 		fmt.Println("Error:", err)
// 		os.Exit(1)
// 	}
// 	// Split string into the substrings and returns a slice of substrings  
// 	s := strings.Split(string(bs), "\n") //Convert byte to string first and split it into slice of string using Split
// 	return deck(s)
// }

// func (d deck) print()  {
// 	for _, element := range d{
// 		fmt.Println(element) 
// 	}
// }

// func main()  {
// 	r := readFileContent("abhijeet")
// 	r.print()
// }



//Solution:
import (
	"os"
	"fmt"
	"log"
	// "io/ioutil"
	// "strings"
	"io"
)


func main() {
	// fmt.Println(os.Args)
	// fmt.Println(os.Args[1]) //Prints second parameter that we provide in go run command

	fmt.Println(os.Args[1])
	file, err := os.Open(os.Args[1]) // For read access.
	if err != nil {
	log.Fatal(err)
	os.Exit(1)
	}
	io.Copy(os.Stdout, file)  //Prints content of file to terminal, It automatically implements the Reader interface
}

// To create or build a binary of this code which is executable we can run: 
// go build exercise1.go
// It builds exercise1 binary
// Run exercise1 binary using 
//  ./exericise1 exercise1.go -> Prints exercise1.go on terminal