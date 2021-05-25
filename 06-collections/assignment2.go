package main

import "fmt"

func main() {
	operationsByChoice := map[int]func(int, int) int{
		1: add,
		2: subtract,
		3: multiply,
		4: divide,
	}
	for {
		userChoice := getUserChoice()
		if operation, exists := operationsByChoice[userChoice]; exists {
			x, y := getOperands()
			result := operation(x, y)
			fmt.Println("Result = ", result)
		} else {
			break
		}
	}
}

func getOperands() (int, int) {
	fmt.Println("Enter the operands")
	var x, y int
	fmt.Scanf("%d %d\n", &x, &y)
	return x, y
}

func add(x, y int) int {
	return x + y
}
func subtract(x, y int) int {
	return x - y
}
func multiply(x, y int) int {
	return x * y
}
func divide(x, y int) int {
	return x / y
}

func getUserChoice() int {
	fmt.Println("Enter your choice :")
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("5. Exit")
	var userChoice int
	fmt.Scanln(&userChoice)
	return userChoice
}
