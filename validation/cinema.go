package validation

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/haidityara/mtix/model/modelcinema"
)

func ValidateStoreCinema(data modelcinema.Request) error {
	return validation.Errors{
		"name":    validation.Validate(data.Name, validation.Required),
		"city_id": validation.Validate(data.CityID, validation.Required),
	}.Filter()
}
