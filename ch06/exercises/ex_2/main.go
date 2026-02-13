package main

import "fmt"

func UpdateSlice(slice []string, value string) {
	slice[len(slice)-1] = value
	fmt.Println(slice)
}

func GrowSlice(slice []string, value string) {
	slice = append(slice, value)
	fmt.Println(slice)
}

func main() {
	slice1 := []string{"Matthew", "Mark", "Luke", "John"}
	fmt.Println(slice1)
	UpdateSlice(slice1, "Thomas")
	fmt.Println(slice1)

	slice2 := []string{"Matthew", "Mark", "Luke", "John"}
	fmt.Println(slice2)
	GrowSlice(slice2, "Thomas")
	fmt.Println(slice2)
}
