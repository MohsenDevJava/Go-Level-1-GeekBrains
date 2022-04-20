package main

import "fmt"

var area int

func main() {
	var lenght, width int
	fmt.Print("Enter Length of Rectangle : ")
	fmt.Scan(&lenght)
	fmt.Print("Enter Width of Rectangle : ")
	fmt.Scan(&width)
	area = lenght * width
	fmt.Println("Area of Rectangle : ", area)
}


//////////////////Command:go run//////////////////////////////////

// PS C:\project> go run rectangle.go
// Enter Length of Rectangle : 23
// Enter Width of Rectangle : 13
// Area of Rectangle :  299
