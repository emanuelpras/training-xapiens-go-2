package helper

import "fmt"

var value = "private"
var Value = "public"

func sayHello(name string) { //func private huruf depannya kecil
	fmt.Println("Hello ", name)
}

func SayHello(name string) { //func public huruf depannya besar
	fmt.Println("Hello ", name)
}
