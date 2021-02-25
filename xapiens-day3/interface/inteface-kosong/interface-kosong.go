package main

import "fmt"

func interfaceKosong(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "-"
	}
}

func main() {
	data := interfaceKosong(2)

	fmt.Println(data)
	fmt.Println("================================")

	//Deklarasi interface kosong di variable
	var contohIKosong interface{}
	contohIKosong = 2

	fmt.Println(contohIKosong)
	fmt.Printf("%T \n", contohIKosong)
	fmt.Println("================================")

	contohIKosong = "value kosong"
	fmt.Println(contohIKosong)
	fmt.Printf("%T", contohIKosong)

}
