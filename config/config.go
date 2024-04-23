package config

import db "service/pkg/database"

type (
	Config struct {
		Logging               Logging      `yaml:"logging"`
		Language              Language     `yaml:language`
		Gateway               Microservice `yaml:"gateway"`
		Debug                 bool         `yaml:"debug"`
		Domain                string       `yaml:"domain"`
		PWD                   string       `yaml:"pwd"`
		AllowOrigins          string       `yaml:"allow_origins"`
		AllowHeaders          string       `yaml:"allow_headers"`
		MaxAge                int          `yaml:"max_age"`
		Timeout               int64        `yaml:"timeout"`
		MaxConcurrentRequests int          `yaml:"max_concurrent_requests"`
		SecretKey             string       `yaml:"secret_key"`
		Media                 string       `yaml:"media"`

		// SMS System
		OtpApiUrl     string `yaml:"otp_api_url"`
		OtpApiKey     string `yaml:"otp_api_key"`
		OtpTemplateId int    `yaml:"otp_template_id"`

		// Based on Days
		AccessTokenLifePeriod int64 `yaml:"access_token_life_period"`
		// Based on Months
		RefreshTokenLifePeriod int64 `yaml:"refresh_token_life_period"`
	}

	Logging struct {
		Path         string `yaml:"path"`
		Pattern      string `yaml:"pattern"`
		MaxAge       string `yaml:"max_age"`
		RotationTime string `yaml:"rotation_time"`
		RotationSize string `yaml:"rotation_size"`
	}

	Language struct {
		Path            string `yaml:"path"`
		DefaultLanguage string `yaml:"default_language"`
	}

	Microservice struct {
		Database db.Database `yaml:"database"`
		IP       string      `yaml:"ip"`
		Port     string      `yaml:"port"`
	}
)
