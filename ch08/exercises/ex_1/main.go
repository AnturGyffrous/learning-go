package main

import "fmt"

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func Double[T Number](n T) T {
	return n * 2
}

func main() {
	var a = 100
	fmt.Println(Double(a))
	var b uint16 = 30_000
	fmt.Println(Double(b))
	var c float32 = 1.57
	fmt.Println(Double(c))
	fmt.Println(Double(1.5707963))
}
