package main

import (
	"errors"
	"fmt"
)

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

	quotient, remainder, _ := divide(70, 8)
	fmt.Println(quotient, remainder)

	counter := getCounter()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	//function composition
	loggedAdd := getLoggedOperation("Add", add)
	loggedAdd(200, 300) // => "Operation:Add, Processing 200 and 300 & result = 500"

	//modify the following line accordingly
	loggedSubtract := getLoggedOperation("Subtract", subtract)
	loggedSubtract(100, 50) // => "Operation:Subtract, Processing 100 and 50 & result = 50"

	increment, _ := getSpinner()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	/*
		fmt.Println(decrement())
		fmt.Println(decrement())
		fmt.Println(decrement())
		fmt.Println(decrement())
	*/

	q, r, err := divide(100, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(q, r)
	}

}

func getLoggedOperation(operName string, oper func(int, int) int) func(int, int) int {
	return func(x, y int) int {
		result := oper(x, y)
		fmt.Printf("Operation : %s, Processing %d and %d & result = %d\n", operName, x, y, result)
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

func divide(x, y int) (quotient int, remainder int, err error) {
	//fmt.Println(float32(x) / float32(y))
	if y == 0 {
		err = errors.New("invalid argument error")
		return
	}
	quotient = x / y
	remainder = x % y
	return
}

func getAdder() func(int, int) int {
	return func(x, y int) int {
		return x + y
	}
}

func getSpinner() (func() int, func() int) {
	count := 0
	increment := func() int {
		count++
		return count
	}
	decrement := func() int {
		count--
		return count
	}
	return increment, decrement
}
