package testify

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0
	result, err := CalculateTax(amount)
	assert.Nil(t, err)
	assert.Equal(t, expected, result, "they should be equal")

	result, err = CalculateTax(0)
	assert.Error(t, err, "amount should be greater than 0")
	assert.Equal(t, 0.0, result)
}
