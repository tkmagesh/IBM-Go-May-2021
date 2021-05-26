package entities

import (
	"fmt"
	"sort"
)

type ProductList []Product

func (productList *ProductList) AddProduct(p *Product) {
	*productList = append(*productList, *p)
}

func (productList ProductList) GetIndex(product Product) int {
	for idx, p := range productList {
		if p == product {
			return idx
		}
	}
	return -1
}

func (productList ProductList) Includes(product Product) bool {
	return productList.GetIndex(product) >= 0
}

func (productList ProductList) All(criteria func(Product) bool) bool {
	for _, x := range productList {
		if !criteria(x) {
			return false
		}
	}
	return true
}

func (productList *ProductList) Any(criteria func(Product) bool) bool {
	for _, x := range *productList {
		if criteria(x) {
			return true
		}
	}
	return false
}

//implement the sort (by using the sort.Sort() method )
/*
	initial implementation should sort the products by id
	provide implementations for sorting the products by cost, units, category etc
*/

func (productList ProductList) Sort() {
	sort.Sort(productList)
}

/* Implementing for the Interface interface used by sort.Sort() */
func (productList ProductList) Len() int {
	return len(productList)
}

func (productList ProductList) Less(i, j int) bool {
	return productList[i].Id < productList[j].Id
}

func (productList ProductList) Swap(i, j int) {
	productList[i], productList[j] = productList[j], productList[i]
}

func (productList ProductList) String() string {
	result := ""
	for _, product := range productList {
		result += fmt.Sprintf("%s\n", product)
	}
	return result
}

type ByCost struct{ ProductList }

func (byCost ByCost) Less(i, j int) bool {
	return byCost.ProductList[i].Cost < byCost.ProductList[j].Cost
}

func (productList ProductList) SortByCost() {
	sort.Sort(ByCost{productList})
}
