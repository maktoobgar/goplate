package config

import _ "embed"

//go:embed config.yml
var ConfigData []byte
