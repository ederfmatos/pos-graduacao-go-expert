package simple

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateTaxBatch(t *testing.T) {
	type test struct {
		amount float64
		expect int8
	}
	tests := []test{
		{amount: 0, expect: 0},
		{amount: -2, expect: 0},
		{amount: 20, expect: 5},
		{amount: 50, expect: 5},
		{amount: 51, expect: 10},
		{amount: 100, expect: 10},
		{amount: 150.1, expect: 12},
		{amount: 205, expect: 12},
	}
	for _, tc := range tests {
		t.Run("TestCalculateTaxBatch", func(t *testing.T) {
			actual := CalculateTax(tc.amount)
			assert.Equal(t, tc.expect, actual)
		})
	}
}
