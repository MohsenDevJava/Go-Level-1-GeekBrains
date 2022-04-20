package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.14
	var (
		area   float64
		radius float64
		circum float64
	)
	fmt.Print("Please enter area of Circle : ")
	fmt.Scan(&area)
	radius = math.Sqrt(area / PI)
	circum = 2 * PI * radius
	fmt.Printf("Radius of Circle is :%.2f\n", radius)
	fmt.Printf("Circumference of Circle is :%.2f\n", circum)
}

//////////////////Command:go run//////////////////////////////////
// PS C:\project> go run circle.go
// Please enter area of Circle : 34
// Radius of Circle is :3.29        
// Circumference of Circle is :20.66


