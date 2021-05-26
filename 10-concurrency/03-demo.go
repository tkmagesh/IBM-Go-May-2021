package main

import "fmt"

func add(x, y int) int {
	result := x + y
	return result
}

func main() {
	result := add(100, 200)
	fmt.Printf("Result = %v\n", result)
}
