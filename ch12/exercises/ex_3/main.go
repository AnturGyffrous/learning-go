package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	for n := 0; n < 100_000; n += 1000 {
		fmt.Println(n, cachedSquareRootMap()[n])
	}
}

var cachedSquareRootMap = sync.OnceValue(calculateSquareRoots)

func calculateSquareRoots() map[int]float64 {
	fmt.Println("Initializing Square Root Map...")
	var squareRoots = make(map[int]float64)
	for i := 0; i < 100_000; i++ {
		squareRoots[i] = math.Sqrt(float64(i))
	}
	return squareRoots
}
