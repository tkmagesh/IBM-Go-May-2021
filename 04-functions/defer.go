package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	defer func() {
		fmt.Println("deferred from main")
		if r := recover(); r != nil {
			fmt.Println(r)
			fmt.Println("Have to shutdown! Contact the administrator")
			debug.PrintStack()
		}
	}()

	fmt.Println(fn())

}

func fn() (result int, err error) {
	fmt.Println("fn triggered")
	/*
		deferredFn := func() {
			fmt.Println("fn completed")
		}
		defer deferredFn()
	*/
	defer func() {
		/* if r := recover(); r != nil {
			err = errors.New("someone panicked")
		} */
		fmt.Println("fn completed - 1")
		panic("panicking from a deferred fn")
	}()
	defer func() {
		fmt.Println("fn completed - 2")
	}()
	defer func() {
		fmt.Println("fn completed - 3")
	}()
	//panic("something went wrong!")
	fmt.Println(divide(100, 10))
	fmt.Println("end of fn execution")
	return
}

func getData() {
	/*
		open file
		defer close file
		for{
			read each line until EOF
			if err {
				return
			}
			parse the data
			if err {
				return
			}
			manipulate the data
			if err {
				return
			}
		}
		return result
	*/
}

func divide(x, y int) (quotient int, remainder int, err error) {
	quotient = x / y
	remainder = x % y
	return
}
