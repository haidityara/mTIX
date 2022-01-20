package validation

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/haidityara/mtix/model/modelmovie"
)

func ValidateMovieStore(data modelmovie.Request) error {
	return validation.Errors{
		"title":        validation.Validate(data.Title, validation.Required),
		"description":  validation.Validate(data.Description, validation.Required),
		"duration":     validation.Validate(data.Duration, validation.Required),
		"language":     validation.Validate(data.Language, validation.Required),
		"genres":       validation.Validate(data.Genres, validation.Required),
		"realise_date": validation.Validate(data.RealiseDate, validation.Required),
		"country":      validation.Validate(data.Country, validation.Required),
	}.Filter()
}
