package simple

func CalculateTax(amount float64) int8 {
	if amount <= 0.0 {
		return 0
	}
	if amount > 150.0 {
		return 12
	}
	if amount > 50.0 {
		return 10
	}
	return 5
}
