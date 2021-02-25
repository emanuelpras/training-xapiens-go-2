package main

import (
	"fmt"
)

// interface kosong itu ibarat dynamic data type
func emptyInterface(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "-"
	}
}

func assertion() interface{} {
	return 12
}

func main() {
	data := emptyInterface(2)

	fmt.Println(data)

	//bisa di skip
	var result interface{} = assertion()
	// resultString := result.(int)

	switch value := result.(type) {
	case string:
		fmt.Println("Data sudah benar karena tipe string dengan nilai ", value)
	default:
		fmt.Println("Tipe data salah karena nilai data", value)
	}

	// fmt.Println(resultString)
}
