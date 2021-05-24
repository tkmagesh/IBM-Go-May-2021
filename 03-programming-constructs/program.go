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

	//switch case
	var newNo int = 21
	//fmt.Scanf("%d", &no)
	switch newNo % 2 {
	case 1:
		fmt.Printf("newNo %d is an odd number of type %T\n", newNo, newNo)
	case 0:
		fmt.Printf("newNo %d is an even number\n", newNo)
	}

	/*
		score
			0 - 3 => "Terrible"
			4 - 7 => "Not Bad"
			8 - 9 => "Good"
			10 => "Excellent"
			else => "Unknown score"
	*/

	score := 7
	switch score {
	case 0, 1, 2, 3:
		fmt.Println("Terrible")
	case 4, 5, 6, 7:
		fmt.Println("Not bad!")
	case 8, 9:
		fmt.Println("Good")
	case 10:
		fmt.Println("Excellect")
	default:
		fmt.Println("Unknown score")
	}
}
