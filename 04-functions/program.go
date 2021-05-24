package main

import "fmt"

/*
var add func(int, int) int = func(x, y int) int {
	return x + y
}
*/
func getCounter() func() int {
	var count int = 0

	return func() int {
		count++
		return count
	}
}

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

	counter := getCounter()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	//function composition
	loggedAdd := getLoggedOperation(add)
	loggedAdd(200, 300)

}

func getLoggedOperation(oper func(int, int) int) func(int, int) int {
	return func(x, y int) int {
		result := oper(x, y)
		fmt.Printf("Processing %d and %d & result = %d\n", x, y, result)
		return result
	}
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func multiple(x, y int) int {
	return x * y
}

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
