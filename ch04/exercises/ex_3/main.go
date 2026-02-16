package main

import "fmt"

func main() {
	var total int
	for i := 0; i < 10; i++ {
		total = total + 1
		fmt.Println(total)
	}

	fmt.Println(total)
}
