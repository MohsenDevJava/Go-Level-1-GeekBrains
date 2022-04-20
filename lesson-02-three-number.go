package main

import (
	"fmt"
)

var (
	number   int64
	ones     int64
	tens     int64
	hundreds int64
)

func main() {
	fmt.Print("Please enter a three-digit number :")
	fmt.Scan(&number)
	ones = number % 10
	number = number / 10
	tens = (number % 10) * 10
	hundreds = (number / 10) * 100
	fmt.Println("Number of ones :", ones)
	fmt.Println("Number of tens :", tens)
	fmt.Println("Number of hundreds :", hundreds)
}

//////////////////Command:go run//////////////////////////////////
// PS C:\project> go run three-number.go
// Please enter a three-digit number :654
// Number of ones : 4      
// Number of tens : 50     
// Number of hundreds : 600
