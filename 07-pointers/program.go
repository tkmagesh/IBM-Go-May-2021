package main

import "fmt"

func main() {
	no := 10
	var noPtr *int = &no
	fmt.Println("Address of no = ", &no)
	fmt.Println("Address of no = ", noPtr)

	//deferencing
	fmt.Printf("Value at %v is %d\n", noPtr, *noPtr)
	fmt.Println("Value before increment = ", no)
	increment(&no)
	fmt.Println("Value after increment = ", no)

	x, y := 100, 200
	fmt.Println("Before swapping x and y are ", x, y)
	swap( /*  */ )
	fmt.Println("After swapping x and y are ", x, y)
}

func increment(noPtr *int) {
	*noPtr = *noPtr + 1
}

func swap( /*  */ ) {

}
