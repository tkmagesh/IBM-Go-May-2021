package main

import (
	"fmt"
	"structs-demo/entities"
)

func main() {
	/* var p Product = Product{
		id:        100,
		name:      "Pen",
		cost:      10,
		units:     100,
		category:  "Stationary",
		isInStock: true,
	} */
	//var p entities.Product = entities.Product{100, "Pen", 10, 100, "Stationary", true}
	p := entities.NewProduct(100, "Pen", 10, 100, "Stationary", true)
	fmt.Println(p)
	fmt.Println("After applying 10% discount")
	entities.ApplyDiscount(p, 10)
	fmt.Println(p)

	//grapes := entities.PerishableProduct{entities.Product{501, "Grapes", 60, 40, "Food", true}, "2 Days"}
	grapes := entities.NewPerishableProject(501, "Grapes", 60, 40, "Food", "2 Days")
	fmt.Println(grapes)
	fmt.Println(grapes.Cost)

}
