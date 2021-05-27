package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

func main() {
	fileWg := &sync.WaitGroup{}
	processWg := &sync.WaitGroup{}
	dataCh := make(chan int)
	evenCh := make(chan int)
	oddCh := make(chan int)
	evenSum := make(chan int)
	oddSum := make(chan int)
	fileWg.Add(2)
	go source("data1.dat", dataCh, fileWg)
	go source("data2.dat", dataCh, fileWg)
	processWg.Add(4)
	go splitter(dataCh, evenCh, oddCh, processWg)
	go sum(evenCh, evenSum, processWg)
	go sum(oddCh, oddSum, processWg)
	go merger(evenCh, oddCh, "result.txt", processWg)
	fileWg.Wait()
	close(dataCh)
	processWg.Wait()
	fmt.Println("Done")
}

func source(filename string, out chan int, wg *sync.WaitGroup) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		str := scanner.Text()
		val, err := strconv.Atoi(str)
		if err != nil {
			log.Println(err)
		}
		out <- val
	}
	wg.Done()
}

func splitter(dataCh chan int, evenCh chan int, oddCh chan int, wg *sync.WaitGroup) {
	for no := range dataCh {
		if no%2 == 0 {
			evenCh <- no
		} else {
			oddCh <- no
		}
	}
	close(evenCh)
	close(oddCh)
	wg.Done()
}

func sum(in chan int, out chan int, wg *sync.WaitGroup) {
	result := 0
	for no := range in {
		result += no
	}
	out <- result
	wg.Done()
}

func merger(evenCh chan int, oddCh chan int, resultFile string, wg *sync.WaitGroup) {
	rf, err := os.Create(resultFile)
	if err != nil {
		panic(err)
	}
	defer rf.Close()
	for i := 0; i < 2; i++ {
		select {
		case oddResult := <-oddCh:
			rf.Write([]byte(fmt.Sprintf("Sum of odd numbers : %d\n", oddResult)))
		case evenResult := <-evenCh:
			rf.Write([]byte(fmt.Sprintf("Sum of even numbers : %d\n", evenResult)))
		}
	}
	wg.Done()
}
