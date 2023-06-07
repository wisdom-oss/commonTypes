package config

// EnvironmentConfguration matches the expected layout of the environment
// configuration file for the microservices.
type EnvironmentConfiguration struct {
	// Required contains an array of environment variable keys that contain
	// values required for the microservice to function.
	Required []string `json:"required"`

	// Optional contains a mapping of string to string values. The key of a
	// single mapping entry reflects the environment variable key and the
	// value denotes the default value for this environment variable if it was
	// not set.
	Optional map[string]string `json:"optional"`
}

// AuthorizationConfiguration contains the configuration for the authorization
// middleware
type AuthorizationConfiguration struct {
	// Enabled indicates if the authorization should be enabled and enforced
	Enabled bool `json:"enableAuth"`

	// RequireUserIdentification indicated if the user id needs to be present
	// to allow a request
	RequireUserIdentification bool `json:"requireUserID"`

	// RequiredUserGroup contains the identification string for the user group
	// that is required to access the microservice
	RequiredUserGroup string `json:"requiredUserGroup"`
}
