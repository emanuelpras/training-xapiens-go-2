package main

import "fmt"

func main() {
	//## for loop cara 1
	for i := 0; i < 5; i++ {
		fmt.Println("Angka ", i)
	}

	// ## for cara 2
	var x = 0

	for x < 5 {
		fmt.Println("Angka ", x)
		x++
	}

	fmt.Println("=========================")

	//## for & break
	var y = 0

	for {
		fmt.Println("Angka ", y)
		y++
		if y == 5 {
			break
		}
	}
}
