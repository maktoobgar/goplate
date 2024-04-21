package build

import "embed"

//go:embed translations
var Translations embed.FS

//go:embed config/config.yml
var Config []byte
