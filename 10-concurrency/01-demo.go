package main

import "fmt"

func print(msg string) {
	fmt.Println(msg)
}

func main() {
	go print("Hello")
	go print("World")
	var input string
	fmt.Println("press ENTER to continue..")
	fmt.Scanln(&input)
}
