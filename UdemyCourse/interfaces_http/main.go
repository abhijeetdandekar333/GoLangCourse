// The HTTP Package : Http request to google.com and print response to terminal.
//  https://pkg.go.dev/std -> Go Packages link
 package main
 
import (
	"net/http"
	"fmt"
	"os"
	"io"
)

type logWriter struct{}

 func main()  {
	resp, err := http.Get("http://google.com")
	if err!= nil {
		fmt.Println("Error Occured!!", err)
		os.Exit(1)
	}
	// fmt.Println("Response is ", resp.Status)
	// fmt.Println("Response is ", resp)

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// Writer and Reader :
	// io.Copy(os.Stdout, resp.Body)  
	// first argument(destination) implements writer interface and second argument implements reader interface (source)

	// Own Custom Writer
	lw := logWriter{}
	io.Copy(lw, resp.Body)

 }

 func (logWriter) Write(bs []byte) (int, error)  {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
 }