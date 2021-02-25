package main

import "fmt"

func main() {
	// ## constant dengan tipe data
	const firstName string = "Emanuel"

	fmt.Println(firstName)

	// ## constant tanpa tipe data
	const ages = 17

	fmt.Println(ages)
	fmt.Printf("%T", ages) // kira kira ini apa?
}
