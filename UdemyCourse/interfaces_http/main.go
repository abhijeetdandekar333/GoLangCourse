// The HTTP Package : Http request to google.com and print response to terminal.
//  https://pkg.go.dev/std -> Go Packages link
 package main
 
import (
	"net/http"
	"fmt"
	"os"
)

 func main()  {
	resp, err := http.Get("http://google.com")
	if err!= nil {
		fmt.Println("Error Occured!!", err)
		os.Exit(1)
	}
	fmt.Println("Response is ", resp.Status)
	fmt.Println("Response is ", resp)
 }