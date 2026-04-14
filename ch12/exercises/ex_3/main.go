package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	for n := 0; n < 100_000; n += 1000 {
		s, err := Calculate(n)
		if err != nil {
			fmt.Println(n, err)
		} else {
			fmt.Println(n, s)
		}
	}
}

type SquareRootCalculator interface {
	Calculate(int) (float64, error)
}

func Calculate(number int) (float64, error) {
	squareRoots := cachedSquareRootCalculator()
	return squareRoots.Calculate(number)
}

var cachedSquareRootCalculator func() SquareRootCalculator = sync.OnceValue(initSquareRootCalculator)

func initSquareRootCalculator() SquareRootCalculator {
	fmt.Println("Initializing Square Root Calculator...")
	m := calculateSquareRoots()
	return SquareRoots{Map: m}
}

func calculateSquareRoots() map[int]float64 {
	var squareRoots = make(map[int]float64)
	for i := 0; i < 100_000; i++ {
		squareRoots[i] = math.Sqrt(float64(i))
	}
	return squareRoots
}

type SquareRoots struct {
	Map map[int]float64
}

func (s SquareRoots) Calculate(number int) (float64, error) {
	if number < 0 {
		return math.NaN(), fmt.Errorf("number must be positive")
	}
	if number >= 100_000 {
		return math.NaN(), fmt.Errorf("number must be less than 100000")
	}
	return s.Map[number], nil
}
