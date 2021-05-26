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
	//entities.ApplyDiscount(p, 10)
	p.ApplyDiscount(10)
	fmt.Println(p)

	//grapes := entities.PerishableProduct{entities.Product{501, "Grapes", 60, 40, "Food", true}, "2 Days"}
	grapes := entities.NewPerishableProject(501, "Grapes", 60, 40, "Food", "2 Days")
	fmt.Println(grapes)
	fmt.Println(grapes.Cost)

	var productList entities.ProductList

	productList.AddProduct(entities.NewProduct(100, "Pen", 10, 100, "Stationary", true))
	productList.AddProduct(entities.NewProduct(101, "Book", 50, 115, "Stationary", true))
	productList.AddProduct(entities.NewProduct(102, "Pencil", 30, 54, "Stationary", true))
	productList.AddProduct(entities.NewProduct(103, "Ink", 20, 76, "Stationary", false))
	productList.AddProduct(entities.NewProduct(104, "Chart", 70, 87, "Stationary", false))

	fmt.Println(productList)
	idx := productList.GetIndex(*entities.NewProduct(103, "Ink", 20, 76, "Stationary", true))
	fmt.Println(idx)

	includes := productList.Includes(*entities.NewProduct(101, "Pen", 10, 100, "Stationary", true))
	fmt.Println(includes)

	allCheck := productList.All(costlyCheck)
	fmt.Println(allCheck)

	anyCheck := productList.Any(costlyCheck)
	fmt.Println(anyCheck)
}

func costlyCheck(p entities.Product) bool {
	return p.Cost > 50
}
