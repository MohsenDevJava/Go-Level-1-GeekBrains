package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter the number of terms : ")
	fmt.Scan(&n)
	fmt.Print("Fibonacci Series :")
	fmt.Print("\n", fib(n))
}
func fib(n int) int {
	baseCases := map[int]int{
		1: 0,
		2: 1,
	}
	fmt.Print(baseCases[1], baseCases[2])
	return computeCache(n, baseCases)
}

func computeCache(n int, cache map[int]int) int {
	if val, found := cache[n]; found {
		return val
	}
	cache[n] = computeCache(n-1, cache) + computeCache(n-2, cache)
	fmt.Print(" ", cache[n])
	return cache[n]
}
