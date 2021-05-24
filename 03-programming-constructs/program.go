package main

import "fmt"

var no1 int

func main() {
	/*
		programming constructs
			if, for, switch
	*/

	//if else
	/*
		no := 21
		if no%2 == 0 {
			fmt.Printf("No %d is an even number\n", no)
		} else {
			fmt.Printf("No %d is an odd number of type %T\n", no, no)
		}
	*/
	fmt.Println("Enter the number :")
	var no int
	fmt.Scanf("%d", &no)
	if no%2 == 0 {
		fmt.Printf("No %d is an even number\n", no)
	} else {
		fmt.Printf("No %d is an odd number of type %T\n", no, no)
	}

}
