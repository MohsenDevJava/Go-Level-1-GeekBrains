package prime

import (
	"errors"
	"fmt"
	"math"
)

var ErrorNon = errors.New("Number must be grater than 2")

func PrimeNumber(num1, num2 int) (int, error) {
	if num1 < 2 || num2 < 2 {
		fmt.Println("")
		return 0, ErrorNon
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
	return 0, nil
}
