package main

import (
	"fmt"
)

func main() {
	// template header
	fmt.Println("============================## Mini Calculator Apps ##============================")
	header()
	var keyword int
	var history []float64

	var number1, number2 float64
	var operator string
	var status bool

	var i int

	fmt.Print("How can i help you? : ")
	fmt.Scanln(&keyword)
menu:
	for keyword > 0 {
		switch keyword {
		case 1:
			status = true
			i = 0
		operate:
			for status != false {
				if i == 0 {
					fmt.Printf("\n")
					fmt.Println("============================ ## Calculator ## ============================")
					fmt.Print("Please enter First number: ")
					fmt.Scanln(&number1)
					fmt.Print("Please enter Operator (+,-,/,*): ")
					fmt.Scanln(&operator)
					if operator == "#" {
						break operate
					}
					fmt.Print("Please enter Second number: ")
					fmt.Scanln(&number2)

				} else {
					fmt.Print("Please enter Operator (+,-,/,%,*): ")
					fmt.Scanln(&operator)
					if operator == "#" {
						break operate
					}
					fmt.Print("Please enter Second number: ")
					fmt.Scanln(&number2)
				}

				var output float64
				switch operator {
				case "+":
					output = number1 + number2
				case "-":
					output = number1 - number2
				case "*":
					output = number1 * number2
				case "/":
					output = number1 / number2
				default:
					fmt.Println("Invalid Operation")

					break operate
				}
				i++
				history = append(history, output)
				fmt.Println("Result :", output)
				number1 = output
			}
		case 2:
			fmt.Printf("\n")
			fmt.Println("============================ ## History ## ============================")
			for a := 0; a < len(history); a++ {
				fmt.Printf("%.1f, ", history[a])
			}
			fmt.Printf("\n")
		case 3:
			fmt.Printf("\n")
			fmt.Println("Thanks for using me! ^_^")
			fmt.Println("============================## End ##============================")
			break menu
		default:
			fmt.Println("Wrong keyword! Please check the Menu above! Your keyword : ", keyword)
		}
		fmt.Printf("\n")
		header()
		fmt.Print("What are you looking for? : ")
		fmt.Scanln(&keyword)
	}

}

func header() {
	fmt.Println("Menu :")
	fmt.Println("=== 1. Calculator")
	fmt.Println("=== 2. History")
	fmt.Println("=== 3. Exit")
	fmt.Println("### Make sure that your input is a number")
}
