package main

import "fmt"

func main() {
	var angka = []int{1, 2, 3, 4, 5, 4, 3, 2, 1}
	var newNumbers = func(min int) []int {
		var r []int
		for _, e := range angka {
			if e < min {
				continue
			}
			r = append(r, e)
		}
		return r
	}(3)

	fmt.Println(angka)
	fmt.Println(newNumbers)
}
