package main

import "fmt"

func fibonacci(count int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < count; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func main() {
	ch := make(chan int, 10)
	go fibonacci(20, ch)
	for no := range ch {
		fmt.Println(no)
	}
	fmt.Println("Done")
}
