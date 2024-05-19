package testify

import "github.com/stretchr/testify/mock"

// Repository is a mock of Repository interface
type TaxRepositoryMock struct {
	mock.Mock
}

// SaveTax mocks the SaveTax method
func (m *TaxRepositoryMock) SaveTax(tax float64) error {
	args := m.Called(tax)
	return args.Error(0)
}
