package main

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	people := make([]Person, 0, 10000000)
	for i := 0; i < 10000000; i++ {
		people = append(people, Person{"Marty", "McFly", 17})
	}
}
