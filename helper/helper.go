// Multiple packages like main we will create more packages
// if we want to use the helper package then we need to import it in main file on given path in go.mod file
// Export this function to use in other packages
// capitalize the functions and variables to export 
package helper

func HelloWorld(message string) string {
	return message;
}

