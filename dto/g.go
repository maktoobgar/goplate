package dto

import (
	"github.com/golodash/galidator"
)

var g = galidator.G().CustomValidators(galidator.Validators{}).CustomMessages(galidator.Messages{
	// Overrides the default galidator messages
	"min":      "MinLength",
	"max":      "MaxLength",
	"len":      "Len",
	"phone":    "Phone",
	"choices":  "Choices",
	"required": "Required",
	"email":    "Email",
})
