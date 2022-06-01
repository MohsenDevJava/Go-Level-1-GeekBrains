package main

import (
	"fmt"
	"math"
)

func main() {
	var n1, n2 int
	fmt.Print("Enter First number :")
	fmt.Scanln(&n1)
	fmt.Print("Enter second number :")
	fmt.Scanln(&n2)
	PrimeNumber(n1, n2)
}

func PrimeNumber(num1, num2 int) {
	if num1 < 2 || num2 < 2 {
		fmt.Println("Number must be grater than 2")
		return
	}
	for num1 < num2 {
		isPrime := true
		for i := 2; i <= int(math.Sqrt(float64(num1))); i++ {
			if num1%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Printf("%d %s", num1, " ")
		}
		num1++
	}
	fmt.Println()
}
