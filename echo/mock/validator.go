// Package mock provide mock objects.
package mock

import (
	"github.com/stretchr/testify/mock"
)

// Validator mock.
type Validator struct {
	mock.Mock
}

// Validate mock implementation.
func (v *Validator) Validate(i interface{}) error {
	return v.Called(i).Error(0)
}
