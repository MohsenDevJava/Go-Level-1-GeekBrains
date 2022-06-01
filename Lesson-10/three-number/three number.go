package three_number

import (
	"errors"
	"fmt"
)

var (
	number   int64
	ones     int64
	tens     int64
	hundreds int64
)

func ThreeNumber(number int64) (int, error) {
	if number <= 0 {
		fmt.Println(errors.New("only for positive"))
		return 0, nil
	}
	ones = number % 10
	number = number / 10
	tens = (number % 10) * 10
	hundreds = (number / 10) * 100
	fmt.Println("Number of ones :", ones)
	fmt.Println("Number of tens :", tens)
	fmt.Println("Number of hundreds :", hundreds)
	return 0, nil
}
