package main

import (
	"fmt"
	"strings"
)

func main() {
	sayHello("Emanuel", containsChar)
}

func sayHello(name string, filter func(string) string) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func containsChar(name string) string {
	if strings.Contains(name, "d") {
		return "..."
	} else {
		return name
	}
}
