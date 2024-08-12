package validators

import (
	"github.com/golodash/galidator/v2"
)

var Generator = galidator.G().CustomValidators(galidator.Validators{
	"email_is_unique":        EmailIsUnique,
	"phone_number_is_unique": PhoneNumberIsUnique,
	"image_type":             ImageType("image/png", "image/jpg", "image/jpeg"),
}).CustomMessages(galidator.Messages{
	// Overrides the default galidator messages
	"min":      "MinLength",
	"max":      "MaxLength",
	"phone":    "Phone",
	"choices":  "Choices",
	"required": "Required",
	"email":    "Email",

	// Custom validators
	"email_is_unique":        "EmailIsUnique",
	"phone_number_is_unique": "PhoneNumberIsUnique",
	"image_type":             "ImageType",
})
