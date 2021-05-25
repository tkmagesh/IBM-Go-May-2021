package main

import (
	"fmt"
	cal "modularity-demo/calculator"

	"github.com/fatih/color"
)

func main() {
	fmt.Println(cal.Add(100, 200))
	fmt.Println(cal.Subtract(100, 200))
	fmt.Println(cal.GetCalcCount())
	color.Blue(greet("Magesh"))
}
