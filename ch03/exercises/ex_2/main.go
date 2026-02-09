package main

import "fmt"

func main() {
	message := "Hi 👩 and 👨"
	r := []rune(message)[3]
	fmt.Println(string(r))
}
