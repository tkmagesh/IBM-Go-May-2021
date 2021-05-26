package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	fileContents, err := ioutil.ReadFile("sample.txt")
	if err != nil {
		log.Fatalln(err)
		return
	}
	err2 := ioutil.WriteFile("sample-copy.txt", fileContents, 0x777)
	if err2 != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println("File copied")
}
