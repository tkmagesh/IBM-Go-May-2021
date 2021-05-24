//package declaration
package main

//import dependent packages
import "fmt"

//package level variables

//package init()

//main function
func main() {

	/*
		var greetMsg string
		greetMsg = "Hello World!"
	*/

	/*
		var greetMsg string = "Hello World!"
	*/

	/*
		var greetMsg = "Hello World!"
	*/

	/*
		// usage of := is limited to function variables
		greetMsg := "Hello World!"
	*/
	//fmt.Println(greetMsg)

	/*
		var name string
		var greetMsg string
		name = "Magesh"
		greetMsg = "Have a nice day!"
	*/

	/*
		var (
			name     string
			greetMsg string
		)
		name = "Magesh"
		greetMsg = "Have a nice day!"
	*/

	/*
		var (
			name     string = "Magesh"
			greetMsg string = "Have a nice day!"
		)
	*/

	/*
		var (
			name, greetMsg string = "Magesh", "Have a nice day!"
		)
	*/

	/*
		var name, greetMsg string = "Magesh", "Have a nice day!"
	*/

	/*
		var name, greetMsg = "Magesh", "Have a nice day!"
	*/
	name, greetMsg := "Magesh", "Have a nice day!"

	fmt.Println(name, greetMsg)
}

//other functions
