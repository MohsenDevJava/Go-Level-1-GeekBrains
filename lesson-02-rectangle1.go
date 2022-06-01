package main

import "fmt"

var (
	are, width, lenght int
)

func main() {
	fmt.Print("Please enter Width : ")
	fmt.Scan(&width)
	fmt.Print("Please enter Lenght : ")
	fmt.Scan(&lenght)
	fmt.Println("Area of rectangle : ", CalculateArea(width, lenght))
	CalcArea(width, lenght)
}

func CalculateArea(w, l int) (are int) {
	are = w * l
	return are
}

func CalcArea(we, le int) {
	aria := we * le
	fmt.Println("Area of rectangle : ", aria)
}

//////////////////Command:go run//////////////////////////////////

// PS C:\project> go run rectangle1.go
// Please enter Width : 43
// Please enter Lenght : 23
// Area of rectangle :  989
// Area of rectangle :  989

