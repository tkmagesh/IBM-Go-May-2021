package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloServer(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
	fmt.Fprintf(res, "Hello %s", req.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", helloServer)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
