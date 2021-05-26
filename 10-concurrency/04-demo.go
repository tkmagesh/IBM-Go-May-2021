package main

import "fmt"

func add(x, y int, ch chan int) {
	result := x + y
	//writing data into the channel
	fmt.Println("Before writing data into the channel")
	ch <- result
	fmt.Println("After writing data into the channel")
}

func main() {
	resultCh := make(chan int)
	go add(100, 200, resultCh)
	fmt.Println("Before reading data from the channel")
	result := <-resultCh
	fmt.Println("Result = ", result)
	fmt.Println("After reading data from the channel")
	/*fmt.Println("Result = ", result) */
	var input string
	fmt.Scanln(&input)
}
