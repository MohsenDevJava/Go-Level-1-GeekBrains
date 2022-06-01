package testify

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	assert := assert.New(t)
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{-6, -4},
		{5555, 5557},
	}

	for _, test := range tests {
		assert.Equal(Calculate(test.input), test.expected)
	}
}
