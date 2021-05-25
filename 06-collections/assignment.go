package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randomNos := generateRandomNos()
	primeNos := filter(randomNos, isPrime)
	evenNos := filter(randomNos, isEven)
	fmt.Println("Prime Nos :=> ", primeNos)
	fmt.Println("Even Nos :=> ", evenNos)
}

func generateRandomNos() []int {
	randomNos := [20]int{}
	for i := 0; i < 20; i++ {
		randomNos[i] = rand.Intn(100)
	}
	return randomNos[:]
}

/*
func filterPrime(nos []int) []int {
	result := []int{}
	for _, randomNo := range nos {
		if isPrime(randomNo) {
			result = append(result, randomNo)
		}
	}
	return result
}

func filterEven(nos []int) []int {
	result := []int{}
	for _, randomNo := range nos {
		if isEven(randomNo) {
			result = append(result, randomNo)
		}
	}
	return result
}
*/

func filter(nos []int, predicate func(int) bool) []int {
	result := []int{}
	for _, randomNo := range nos {
		if predicate(randomNo) {
			result = append(result, randomNo)
		}
	}
	return result
}

func isPrime(n int) bool {
	for i := 2; i < (n / 2); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func isEven(n int) bool {
	return n%2 == 0
}
