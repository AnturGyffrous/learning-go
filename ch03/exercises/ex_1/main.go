package main

import "fmt"

func main() {
	var s []string = []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	s1 := s[:2]
	s2 := s[1:4]
	s3 := s[3:]
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}
