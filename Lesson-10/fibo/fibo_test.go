package fibo

import (
	"errors"
	"testing"
)

func TestFib(t *testing.T) {
	testCase := []struct {
		Input          int
		Expectedoutput int
		Name           string
		Expectederror  error
	}{
		{
			Name:           "case for 10",
			Input:          10,
			Expectedoutput: 55,
		},
		{
			Name:           "case for 9",
			Input:          9,
			Expectedoutput: 34,
		},
		{
			Name:          "negative number",
			Input:         -10,
			Expectederror: ErrorNonPositive,
		},
	}
	for _, cse := range testCase {
		cse := cse
		t.Run(cse.Name, func(t *testing.T) {
			result, err := Fib(cse.Input)
			if err != nil {
				if !errors.Is(err, cse.Expectederror) {
					t.Fatalf("unexpected error happened:%s", err.Error())
				}
				return
			}
			if result != cse.Expectedoutput {
				t.Fatalf("resualt incorecct ,got:%d ,expected:%d", result, cse.Expectedoutput)
			}
		})
	}

}

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Fib(15)
	}
}

func BenchmarkFibWithCache(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = FibWithCache(15, map[int]int{0: 0, 1: 1})
	}
}
