package main

import (
	"fmt"
	"sync"
)

/*
	The following is an example for "Sharing memory for communication" - not preferred
	The preferred solution is "Communicate by sharing memory" (04-demo.go)
*/
var result int

var wg *sync.WaitGroup = &sync.WaitGroup{}

func add(x, y int) {
	result = x + y
	wg.Done()
}

func main() {
	wg.Add(1)
	go add(100, 200)
	wg.Wait()
	fmt.Printf("Result = %v\n", result)
}
