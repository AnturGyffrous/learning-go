package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName string, lastName string, age int) Person {
	person := Person{FirstName: firstName, LastName: lastName, Age: age}
	return person
}

func MakePersonPointer(firstName string, lastName string, age int) *Person {
	person := Person{FirstName: firstName, LastName: lastName, Age: age}
	return &person
}

func main() {
	person1 := MakePerson("Marty", "McFly", 17)
	person2 := MakePersonPointer("Lorraine", "Baines", 47)
	fmt.Println(person1)
	fmt.Println(person2)
}
