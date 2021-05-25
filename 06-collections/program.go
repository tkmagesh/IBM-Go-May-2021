package main

import (
	"fmt"
)

func main() {

	//Array
	//var nos [5]int

	//var nos [5]int = [5]int{3, 1, 5, 2, 4}
	//var nos [5]int = [...]int{3, 1, 4, 2, 5}
	//var nos = [...]int{3, 1, 2, 4, 5}
	nos := [...]int{3, 1, 4, 2, 5}
	fmt.Println(nos)
	fmt.Println(len(nos))

	/*
		for i := 0; i < len(nos); i++ {
			fmt.Printf("index=%d, value=%d\n", i, nos[i])
		}
	*/

	for idx, value := range nos {
		fmt.Printf("index=%d, value=%d\n", idx, value)
	}

	//slice
	var names []string = []string{"Magesh", "Suresh", "Ramesh"}
	fmt.Println(names, len(names))
	newNames := []string{"Ganesh", "Rajesh"}
	names = append(names, "Philip", "John")
	names = append(names, newNames...)
	fmt.Println(names, len(names))

	replaceName(names, 0, "Sajid")
	fmt.Println("After replacing Magesh with Sajid")
	fmt.Println(names)
	//slicing
	/*
		list[lo : hi] => from lo to hi-1
		list[:hi] => from 0 to hi-1
		list[lo :] => from lo to len(list)
		list[lo : lo] => empty
		list[lo : lo+1] => list[lo]
	*/
	fmt.Println("names[2:4] => ", names[2:4])
	fmt.Println("names[:3] => ", names[:3])
	fmt.Println("names[3:] =>", names[3:])

	//map
	cityRanks := map[string]int{
		"Mysuru":    2,
		"Bengaluru": 5,
		"Udupi":     1,
	}
	fmt.Println("Rank of Bengalure => ", cityRanks["Bengaluru"])
	//cityRanks["Mangaluru"] = 3
	fmt.Println("After adding Mangaluru")
	addCityRank(cityRanks, "Mangaluru", 3)
	fmt.Println(cityRanks)

	for city, rank := range cityRanks {
		fmt.Printf("City = %s, Rank = %d\n", city, rank)
	}

	//cityRanks["Chennai"] = 6
	fmt.Println("Check if Chennai exists")
	if rank, exists := cityRanks["Chennai"]; exists {
		fmt.Println("Chennai is ranked at ", rank)
	} else {
		fmt.Println("Chennai is not ranked")
	}

	fmt.Println("After deleting Mysuru")
	delete(cityRanks, "Mysuru")
	fmt.Println(cityRanks)

}

func addCityRank(cityRanks map[string]int, city string, rank int) {
	cityRanks[city] = rank
}

func replaceName(names []string, idx int, name string) {
	names[idx] = name
}
