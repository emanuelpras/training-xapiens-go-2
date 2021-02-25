package main

import (
	"fmt"
)

func selectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}
		tmp := arr[i]
		arr[i] = arr[minIndex]
		arr[minIndex] = tmp
	}
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func main() {
	arr := []int{1, 2, 33, 44, 55, 33, 44, 22, 44, 33, 2, 77, 66, 1, 2, 3, 4, 5, 6, 7, 89, 3, 3, 8, 9, 75, 4, 3, 2, 2, 1, 3, 4, 5, 5, 6, 7, 8, 8, 9, 21, 15, 12, 98, 34, 11, 12, 13, 9, 99}
	selectionSort(arr)

	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Print("\n")
	limit := 25

	for i := 0; i < len(arr); i += limit {
		batch := arr[i:min(i+limit, len(arr))]
		fmt.Println(batch)
		fmt.Print("hasil :", sum(batch))
		fmt.Print("\n")
	}
	min := arr[0]

	for _, v := range arr {
		if v < min {
			min = v
		}
	}

	fmt.Print("minumum: ", min)
	fmt.Print("\n")
	max := arr[0]

	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	fmt.Print("maximum: ", max)
	fmt.Print("\n")
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func calSum(mid, m int, nums []int) bool {
	sum, count := 0, 0
	for _, v := range nums {
		sum += v
		if sum > mid {
			sum = v
			count++

			if count > m-1 {
				return false
			}
		}
	}
	return true
}

func findMinAndMax(a [5]int) (min int, max int) {
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
