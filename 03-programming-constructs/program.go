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
	//fmt.Println("Enter the number :")
	var no int = 20
	//fmt.Scanf("%d", &no)
	if no%2 == 0 {
		fmt.Printf("No %d is an even number\n", no)
	} else {
		fmt.Printf("No %d is an odd number of type %T\n", no, no)
	}

	//for construct
	/*
		sum := 0
		for i := 1; i <= 10; i++ {
			sum += i
		}
		fmt.Println(sum)
	*/

	//'for' used like a 'while' construct
	/*
		sum := 1
		for sum < 100 {
			sum += sum
		}
		fmt.Println(sum)
	*/

	//infinite loop
	/*
		sum := 1
		for {
			if sum > 100 {
				break
			}
			sum += sum
		}
		fmt.Println(sum)
	*/

}
