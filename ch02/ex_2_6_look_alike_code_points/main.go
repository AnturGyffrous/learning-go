package main

import "fmt"

func main() {
	ａ := "hello"   // Unicode U+FF41
	a := "goodbye" // standard lowercase a (Unicode U+0061)
	fmt.Println(ａ)
	fmt.Println(a)
}
