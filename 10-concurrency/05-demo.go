package main

import (
	"fmt"
	"time"
)

func print(msg string, in, out chan string) {
	for i := 0; i < 5; i++ {
		<-in
		fmt.Println(msg)
		time.Sleep(1 * time.Second)
		out <- "done"
	}
}

func print2(msg string, in, out chan string) {
	for i := 0; i < 5; i++ {
		<-in
		fmt.Println(msg)
		time.Sleep(2 * time.Second)
		out <- "done"
	}
}
func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go print("Hello", ch1, ch2)
	go print2("World", ch2, ch1)
	ch1 <- "start"
	var input string
	fmt.Scanln(&input)
}
