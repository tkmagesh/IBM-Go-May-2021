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
	swap(&x, &y)
	fmt.Println("After swapping x and y are ", x, y)

	nos := []int{10, 20, 30}
	addValue(&nos, 40)
	fmt.Println(nos) //=> 10, 20, 30, 40
}

func addValue(nos *[]int, value int) {
	*nos = append(*nos, value)
}

func increment(noPtr *int) {
	*noPtr = *noPtr + 1
}

func swap(x, y *int) {
	*x, *y = *y, *x
}
