package main

import (
	"fmt"
	"sync"
)

var wg *sync.WaitGroup = &sync.WaitGroup{}

func print(msg string) {
	fmt.Println(msg)
	wg.Done()
}

func main() {
	wg.Add(2)
	go print("Hello")
	go print("World")
	wg.Wait()
	fmt.Println("Exiting from the application")
}
