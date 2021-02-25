package main

import (
	"fmt"
	"sort"
)

func main() {
	data := [...]int{1, 2, 33, 44, 55, 33, 44, 22, 44, 33, 2, 77, 66, 1, 2, 3, 4, 5, 6, 7, 89, 3, 3, 8, 9, 75, 4, 3, 2, 2, 1, 3, 5, 6, 6, 34, 2, 6, 72, 2, 42, 42, 43, 31, 36, 9, 10, 7, 19, 20}
	fmt.Println("========== Data ==========")
	fmt.Println(data)
	fmt.Printf("\n")
	length := len(data) / 2
	newData1 := data[:length] //kelompok 1
	newData2 := data[length:] //kelompok 2

	fmt.Println("========== Bagi Data ke Dalam 2 kelompok ==========")
	fmt.Println("Kelompok Data 1 : ", newData1) //output kelompok1
	fmt.Println("Kelompok Data 2 : ", newData2) //output kelompok2
	fmt.Printf("\n")

	fmt.Println("========== Kelompok Data 1 ==========")
	fmt.Println(newData1)
	fmt.Printf("\n")
	fmt.Println("Urutkan data dari kecil ke besar :")
	sort.Sort(sort.IntSlice(newData1))
	fmt.Println(newData1)
	fmt.Printf("\n")
	fmt.Println("Tampilkan total data :")
	fmt.Println(total(newData1))
	fmt.Printf("\n")
	fmt.Println("Rata2 data :")
	fmt.Println(float64(total(newData1)) / float64(len(newData1)))
	fmt.Printf("\n")

	fmt.Println("Min & Max :")
	min, max := findMinAndMax(newData1)
	fmt.Println("Data 1: ", newData1)
	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)
	fmt.Printf("\n")

	fmt.Println("========== Kelompok Data 2 ==========")
	fmt.Println(newData2)
	fmt.Printf("\n")
	fmt.Println("Urutkan data dari kecil ke besar :")
	sort.Sort(sort.IntSlice(newData2))
	fmt.Println(newData2)
	fmt.Printf("\n")
	fmt.Println("Tampilkan total data :")
	fmt.Println(total(newData2))
	fmt.Printf("\n")
	fmt.Println("Rata2 data :")
	fmt.Println(float64(total(newData2)) / float64(len(newData2)))
	fmt.Printf("\n")

	fmt.Println("Min & Max :")
	min, max = findMinAndMax(newData2)
	fmt.Println("Data 1: ", newData2)
	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)
	fmt.Printf("\n")
}

func total(val []int) int {
	totals := 0
	for i := 0; i < len(val); i++ {
		totals += val[i]
	}
	return totals
}

func findMinAndMax(a []int) (min int, max int) {
	min = a[0]
	max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}
