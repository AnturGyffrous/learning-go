package main

import (
	"fmt"
	"math/rand"
)

func main() {
	numbers := make([]int, 100)
	for i := range numbers {
		numbers[i] = rand.Intn(100)
	}

loop:
	for _, v := range numbers {
		switch {
		case v%2 == 0 && v%3 == 0:
			fmt.Println("Six!")
			break loop
		case v%2 == 0:
			fmt.Println("Two!")
		case v%3 == 0:
			fmt.Println("Three!")
		default:
			fmt.Println("Never mind")
		}

	}
}
