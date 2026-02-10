package main

import (
	"fmt"
)

type Employee struct {
	firstName string
	lastName  string
	id        int
}

func main() {
	jb := Employee{"James", "Bond", 7}

	mcfly := Employee{firstName: "Marty", lastName: "McFly", id: 1985}

	var ripley = Employee{}
	ripley.firstName = "Ellen"
	ripley.lastName = "Ripley"
	ripley.id = 5156170

	fmt.Println(jb)
	fmt.Println(mcfly)
	fmt.Println(ripley)
}
