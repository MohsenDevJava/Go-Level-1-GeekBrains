package main

import (
	"fmt"
)

var prime []int

func main() {
	length := 0
	fmt.Println("Enter the number of inputs:")
	fmt.Scanln(&length)
	fmt.Println("Enter the inputs:")
	numbers := make([]int, length)
	for i := 0; i < length; i++ {
		fmt.Scanln(&numbers[i])
	}
	for _, s := range numbers {
		isPrime := primes_dns(s)
		fmt.Println(isPrime)
	}
	InsertSort(prime)
	fmt.Println(prime)

}
func primes_dns(num int) bool {
	for i := 2; i <= num/i; i++ {
		if num%i == 0 {
			return false
		}
	}
	prime = append(prime, num)
	return true
}

func InsertSort(prime []int) {
	for i := 0; i < len(prime); i++ {
		x := prime[i]
		j := i
		for ; j >= 1 && prime[j-1] > x; j-- {
			prime[j] = prime[j-1]
		}
		prime[j] = x
	}
}
