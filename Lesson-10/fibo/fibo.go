package fibo

import (
	"errors"
)

var ErrorNonPositive = errors.New("only for positive")

func Fib(n int) (int, error) {
	if n < 0 {
		return 0, ErrorNonPositive
	}
	if n <= 1 {
		return n, nil
	}
	a, err := Fib(n - 1)
	if err != nil {
		return 0, err
	}
	b, err := Fib(n - 2)
	if err != nil {
		return 0, err
	}
	return a + b, nil
}

func FibWithCache(n int, cache map[int]int) int {
	if val, ok := cache[n]; ok {
		return val
	}
	cache[n] = FibWithCache(n-1, cache) + FibWithCache(n-2, cache)

	return cache[n]
}

//var n int
//t1 := 0
//t2 := 1
//nextTerm := 0

//fmt.Print("Enter the number of terms : ")
//fmt.Scan(&n)
//fmt.Print("Fibonacci Series :")

//for i := 1; i <= n; i++ {
//if i == 1 {
//fmt.Print(" ", t1)
//continue
//}
//if i == 2 {
//fmt.Print(" ", t2)
//continue
//}
//nextTerm = t1 + t2
//t1 = t2
//t2 = nextTerm
//fmt.Print(" ", nextTerm)
//}
//type Fibona struct {
//	T1       int
//	T2       int
//	NextTerm int
//}
//
//func (f *Fibona) fibo(num int) {
//	for i := 1; i <= num; i++ {
//		if i == 1 {
//			fmt.Print(" ", f.T1)
//			continue
//		}
//		if i == 2 {
//			fmt.Print(" ", f.T2)
//			continue
//		}
//		f.NextTerm = f.T1 + f.T2
//		f.T1 = f.T2
//		f.T2 = f.NextTerm
//		fmt.Print(" ", f.NextTerm)
//	}
//
//}
//
//func main() {
//	var n int
//	fmt.Print("Enter the number of terms : ")
//	fmt.Scan(&n)
//	fmt.Print("Fibonacci Series :")
//	resualt := Fibona{}
//	resualt.fibo(n)
//}
