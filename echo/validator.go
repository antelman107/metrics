// Package echo provide implementations of custom functionality for the echo framework.
package echo

import (
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

// ValidatorWrapper implementation
type validatorWrapper struct {
	validator *validator.Validate
}

// NewValidatorWrapper constructor
func NewValidatorWrapper(v *validator.Validate) echo.Validator {
	return &validatorWrapper{
		validator: v,
	}
}

// Validate data
func (c *validatorWrapper) Validate(i interface{}) error {
	return c.validator.Struct(i)
}
