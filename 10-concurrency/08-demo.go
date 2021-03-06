package main

import (
	"fmt"
)

func writeData(ch chan int) {
	fmt.Println("Before writing 10 into the channel")
	ch <- 10
	fmt.Println("After writing 10 into the channel")
	fmt.Println("Before writing 20 into the channel")
	ch <- 20
	fmt.Println("After writing 20 into the channel")
	fmt.Println("Before writing 30 into the channel")
	ch <- 30
	fmt.Println("After writing 30 into the channel")
	fmt.Println("Before writing 40 into the channel")
	ch <- 40
	fmt.Println("After writing 40 into the channel")
	close(ch)

}
func main() {
	fmt.Println("Starting the main function")
	ch := make(chan int, 2)
	go writeData(ch)
	for no := range ch {
		fmt.Println("Before reading data from the channel")
		fmt.Println(no)
		fmt.Println("After reading data from the channel")
	}
	fmt.Println("Exiting from main function")
}
