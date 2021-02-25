package main

import "fmt"

type hashName interface {
	getName() string
	getAge() int
}

type Person struct {
	name string
	age  int
}

func (person Person) getAge() int {
	return person.age
}

func (person Person) getName() string {
	return person.name
}

func main() {
	person := Person{"Villa", 20}
	sayHello(person)
}

func sayHello(person hashName) {
	fmt.Println("Hello ", person.getName(), "umur anda ", person.getAge())
}
