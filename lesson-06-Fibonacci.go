package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter the number of terms : ")
	fmt.Scan(&n)
	fmt.Print("Fibonacci Series :")
	for i := 0; i < n; i++ {
		fmt.Print(" ", fibonacci(i))
	}
}
func fibonacci(num int) int {
	if num == 0 {
		return 0
	}
	if num == 1 {
		return 1
	}
	return fibonacci(num-1) + fibonacci(num-2)
}
