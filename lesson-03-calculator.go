package main

import "fmt"

func main() {
	var num1, num2 float64
	var operators string
	fmt.Print("Enter the First Number: ")
	fmt.Scanln(&num1)
	fmt.Print("Enter the Second Number:")
	fmt.Scanln(&num2)
	fmt.Print("Chose the operator (+ - * / factorial ) :")
	fmt.Scanln(&operators)

	switch operators {
	case "+":
		fmt.Printf("%.2f %s %.2f = %.2f", num1, operators, num2, num1+num2)
	case "-":
		fmt.Printf("%.2f %s %.2f = %.2f", num1, operators, num2, num1-num2)
	case "*":
		fmt.Printf("%.2f %s %.2f = %.2f", num1, operators, num2, num1*num2)
	case "/":
		if num2 == 0 {
			fmt.Println("Divide by zero Situation")
		} else {
			fmt.Printf("%.2f %s %.2f = %.2f", num1, operators, num2, num1/num2)
		}
	case "factorial":
		fmt.Printf("%.f \n", factorial(num1))
		fmt.Printf("%.f \n", factorial(num2))
	default:
		fmt.Println("Invalid operator")
	}
}
func factorial(num float64) float64 {
	if num == 1 || num == 0 {
		return num
	}
	return num * factorial(num-1)
}
