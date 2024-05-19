package tax

import "testing"

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0
	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected %.2f, got %.2f", expected, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	tests := []struct {
		amount   float64
		expected float64
	}{
		{500.0, 5.0},
		{1000.0, 10.0},
	}

	for _, test := range tests {
		result := CalculateTax(test.amount)
		if result != test.expected {
			t.Errorf("Expected %.2f, got %.2f", test.expected, result)
		}
	}
}
