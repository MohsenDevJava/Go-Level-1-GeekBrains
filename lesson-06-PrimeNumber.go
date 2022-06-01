package main

import (
	"fmt"
	"math"
)

var prime []int

func main() {
	var length int
	fmt.Println("Enter the number of inputs:")
	fmt.Scanln(&length)
	fmt.Println("Enter the inputs:")

	numbers := make([]int, length)
	for i := 0; i < length; i++ {
		fmt.Scanln(&numbers[i])
	}
	for _, s := range numbers {
		checkPrimeNumber(&s)
	}
	fmt.Println(prime)

}
func checkPrimeNumber(num *int) {
	if *num < 2 {
		fmt.Println("Number must be greater than 2.")
		return
	}
	sq_root := int(math.Sqrt(float64(*num)))
	for i := 2; i <= sq_root; i++ {
		if *num%i == 0 {
			return
		}
	}
	prime = append(prime, *num)
	return
}
