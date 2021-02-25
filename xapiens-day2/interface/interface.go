package main

import "fmt"

// analogi buat bikin flow bisnis tuh ada fungsi apa aja
type hashName interface {
	getName() string
	getAge() int
}

type person struct {
	name string
	age  int
}

func sayHello(person hashName) {
	fmt.Println("Hello", person.getName(), "anda berumur ", person.getAge())
}

func (pers person) getName() string {
	return pers.name
}

func (pers person) getAge() int {
	return pers.age
}

func main() {
	prsn := person{"Villa", 19}
	sayHello(prsn)
}
