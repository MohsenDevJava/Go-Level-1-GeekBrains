package prime

import (
	"fmt"
	"testing"
)

func TestPrimeNumber(t *testing.T) {
	testNumber := []struct {
		Case    string
		Number1 int
		Number2 int
	}{
		{
			Case:    "Number-1",
			Number1: 10,
			Number2: 30,
		}, {
			Case:    "Number-2",
			Number1: 20,
			Number2: 50,
		}, {
			Case:    "Number-3",
			Number1: 30,
			Number2: 60,
		},
	}
	for _, Pnum := range testNumber {
		Pnum = Pnum
		fmt.Printf("Name Case:%v\n", Pnum.Case)
		t.Run(Pnum.Case, func(t *testing.T) {
			_, err := PrimeNumber(Pnum.Number1, Pnum.Number2)
			if err != nil {
				t.Fatalf("unexpected error happened:%s", err.Error())
			}
		})
	}
}

func BenchmarkPrimeNumber(b *testing.B) {

	for i := 0; i < b.N; i++ {
		_, _ = PrimeNumber(20, 50)
	}
}
