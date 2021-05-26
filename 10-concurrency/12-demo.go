package main

import (
	"fmt"
	"time"
)

func fibonacci(ch chan int) {
	x, y := 0, 1
	quit := time.After(10 * time.Second)
	for {
		select {
		case ch <- x:
			time.Sleep(1 * time.Second)
			x, y = y, x+y
		case <-quit:
			fmt.Println("Quitting")
			close(ch)
			return
		}
	}
}

func main() {
	ch := make(chan int, 10)
	go fibonacci(ch)
	go func() {
		for no := range ch {
			fmt.Println(no)
		}
	}()
	fmt.Println("press ENTER to stop...")
	var input string
	fmt.Scanln(&input)
	fmt.Println("Done")
}
