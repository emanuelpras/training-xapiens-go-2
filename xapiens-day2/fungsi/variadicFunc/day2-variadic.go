package main

import (
	"fmt"
)

func main() {
	var avr = funcCalculate(1, 2, 3, 4, 5)
	fmt.Println(avr)
}

func funcCalculate(angka ...int) float64 {
	var total int = 0

	for _, number := range angka {
		total += number
	}

	var avg = float64(total) / float64(len(angka))
	return avg
}
