package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func getResponse(ch chan *http.Response) {
	res, err := http.Get("https://jsonplaceholder.typicode.com/comments")
	if err != nil {
		log.Fatalln(err)
	}
	ch <- res
}

func main() {
	ch := make(chan *http.Response)
	timeout := time.After(200 * time.Millisecond)
	go getResponse(ch)
	for {
		select {
		case res := <-ch:
			io.Copy(os.Stdout, res.Body)
			return
		case <-timeout:
			log.Fatalln("Request timed out!!")
		}
	}

}
