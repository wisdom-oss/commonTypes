package config

import "errors"

var (
	ErrRequiredVariableMissing = errors.New("required environment variable missing")
	ErrRequiredVariableEmpty   = errors.New("required environment variable empty")
	ErrFilePathEmpty           = errors.New("empty file path for docker secret")
)
