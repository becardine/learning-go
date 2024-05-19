package testify

import (
	"errors"
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

func TestCalculateTaxAndSave(t *testing.T) {
	amount := 500.0
	repository := &TaxRepositoryMock{}
	repository.On("SaveTax", 5.0).Return(nil).Once()
	repository.On("SaveTax", 0.0).Return(errors.New("error saving tax"))

	err := CalculateTaxAndSave(amount, repository)
	assert.Nil(t, err)

	err = CalculateTaxAndSave(0, repository)
	assert.Error(t, err, "error saving tax")
	repository.AssertExpectations(t)
	repository.AssertNumberOfCalls(t, "SaveTax", 2)
}
