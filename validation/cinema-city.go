package validation

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/haidityara/mtix/model/modelcinemacity"
)

func ValidateStoreCinemaCity(data modelcinemacity.Request) error {
	return validation.Errors{
		"name":  validation.Validate(data.Name, validation.Required),
		"state": validation.Validate(data.State, validation.Required),
		"zip":   validation.Validate(data.Zip, validation.Required),
	}.Filter()
}
