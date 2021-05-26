package main

import (
	"fmt"
	"time"
)

func fibonacci(ch chan int, quit chan string) {
	x, y := 0, 1
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
	quit := make(chan string)
	go fibonacci(ch, quit)
	go func() {
		for no := range ch {
			fmt.Println(no)
		}
	}()
	fmt.Println("press ENTER to stop...")
	var input string
	fmt.Scanln(&input)
	quit <- "done"
	fmt.Println("Done")
}
