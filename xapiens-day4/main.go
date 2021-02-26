package main

import (
	"fmt"
	"xapiens-day4/helper"
)

func main() {
	var val string
	helper.SayHello("Pras")
	val = helper.Value

	fmt.Println("Ini variable dari helper ", val)
}
