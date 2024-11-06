package main

import "fmt"

// main runs a simple calculator application in a loop, allowing the user to select
// operations such as addition, subtraction, multiplication, and division. The user
// can perform calculations with two floating-point numbers and choose to continue
// or exit the application. The screen is cleared before displaying the menu and
// results. Division by zero is handled with an error message.
func main() {
	for {
		fmt.Println("\033[H\033[2J") // clear screen
		fmt.Println("Simple Calculator")
		fmt.Println("----------------")
		name := "mersad"

		fmt.Println(name)
		for {
			fmt.Println("1. Addition")
			fmt.Println("2. Subtraction")
			fmt.Println("3. Multiplication")
			fmt.Println("4. Division")
			fmt.Println("5. Exit")

			var choice int
			fmt.Print("Please select an option: ")
			fmt.Scanln(&choice)

			if choice == 5 {
				return
			}

			var num1, num2 float64
			fmt.Print("Please enter the first number: ")
			fmt.Scanln(&num1)
			fmt.Print("Please enter the second number: ")
			fmt.Scanln(&num2)

			var result float64
			switch choice {
			case 1:
				result = num1 + num2
			case 2:
				result = num1 - num2
			case 3:
				result = num1 * num2
			case 4:
				if num2 != 0 {
					result = num1 / num2
				} else {
					fmt.Println("Error: Division by zero!")
					continue
				}
			default:
				fmt.Println("Invalid option!")
				continue
			}

			fmt.Println("\033[H\033[2J")               // clear screen
			fmt.Printf("\033[1m%.2f\033[0m\n", result) // print result in bold

			var cont string
			fmt.Print("Do you want to continue? (yes/no): ")
			fmt.Scanln(&cont)

			if cont != "yes" {
				return
			}
		}
	}
}
