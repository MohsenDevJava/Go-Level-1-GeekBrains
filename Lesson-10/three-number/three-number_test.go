package three_number

import (
	"fmt"
	"testing"
)

func TestThreeNumber(t *testing.T) {
	testCaseNumber := []struct {
		testCase string
		number   int64
	}{
		{
			testCase: "No.1",
			number:   354,
		}, {
			testCase: "No.2",
			number:   234,
		}, {
			testCase: "No,3",
			number:   -543,
		},
	}

	for _, DigitNum := range testCaseNumber {
		DigitNum = DigitNum
		fmt.Printf("%v\n", DigitNum.testCase)
		t.Run(DigitNum.testCase, func(t *testing.T) {
			_, err := ThreeNumber(DigitNum.number)
			if err != nil {
				t.Fatalf("unexpected error happened:%s", err)
			}

		})
	}
}

func BenchmarkThreeNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ThreeNumber(213)
	}
}
