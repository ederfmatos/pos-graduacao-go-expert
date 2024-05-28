package simple

import "testing"

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
			if actual != tc.expect {
				t.Fatalf("expected %d, got %d", tc.expect, actual)
			}
		})
	}
}

func BenchmarkCalculateTaxBatch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500)
	}
}

func FuzzCalculateTaxBatch(f *testing.F) {
	seed := []float64{-1, -2, -3, -4.5, 5, 1, 1.1, 10, 20, 30, 25, 40, 69, 205, 4000}
	for _, amount := range seed {
		f.Add(amount)
	}
	f.Fuzz(func(t *testing.T, amount float64) {
		actual := CalculateTax(amount)
		if amount > 150 && actual != 12 {
			t.Fatalf("for value %f, expected 12, got %d", amount, actual)
		}
		if amount > 50 && amount <= 150 && actual != 10 {
			t.Fatalf("for value %f, expected 10, got %d", amount, actual)
		}
		if amount <= 0 && actual != 0 {
			t.Fatalf("for value %f, expected 0, got %d", amount, actual)
		}
		if amount > 0 && amount <= 50 && actual != 5 {
			t.Fatalf("for value %f, expected 5, got %d", amount, actual)
		}
	})
}
