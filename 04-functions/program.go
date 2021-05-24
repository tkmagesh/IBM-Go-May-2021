package main

import "fmt"

/*
var add func(int, int) int = func(x, y int) int {
	return x + y
}
*/

func main() {
	x, y := 100, 200
	/*
		add := func(x, y int) int {
			return x + y
		}
	*/
	add := getAdder()
	fmt.Printf("Adding %d and %d results in %d \n", x, y, add(x, y))

	quotient, remainder := divide(70, 8)
	fmt.Println(quotient, remainder)
}

/*
func add(x, y int) int {
	return x + y
}
*/

/*
func divide(x, y int) (int, int) {
	//fmt.Println(float32(x) / float32(y))
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
*/

func divide(x, y int) (quotient int, remainder int) {
	//fmt.Println(float32(x) / float32(y))
	quotient = x / y
	remainder = x % y
	return
}

func getAdder() func(int, int) int {
	return func(x, y int) int {
		return x + y
	}
}
