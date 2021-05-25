package main

import (
	"fmt"
	cal "modularity-demo/calculator"
)

func main() {
	fmt.Println(cal.Add(100, 200))
	fmt.Println(cal.Subtract(100, 200))
	fmt.Println(cal.GetCalcCount())
	fmt.Println(greet("Magesh"))
}
