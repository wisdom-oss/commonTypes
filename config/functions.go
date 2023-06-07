package wisdomType

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

// PopulateFromFile takes an already opened file handle and reads the files
// contents into the environment configuration and returns an error if one
// occurs.
//
// This function is a shorthand call for json.NewDecoder(reader).Decode(target)
func (c *EnvironmentConfiguration) PopulateFromFile(f *os.File) error {
	return json.NewDecoder(f).Decode(c)
}

// PopulateFromFilePath takes a file path to the configuration file and
// populates the values using the PopulateFromFile function and returns
// an error if one occurrs.
func (c *EnvironmentConfiguration) PopulateFromFilePath(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	return c.PopulateFromFile(file)
}

// ParseEnvironment takes the already processed environment configuration and
// checks the acutal environment for the needed variables and populates them
// accordingly. If a required variable is missing, a error will be returned.
//
// The function supports reading docker secrets. The environment variable key
// will be extended with `_FILE` and checked for a path.
func (c EnvironmentConfiguration) ParseEnvironment() (map[string]string, error) {
	// create an empty mapping to allow appending found values
	parsedVariables := make(map[string]string)

	// start by parsing the required environment variables
	for _, envKey := range c.Required {
		// check if the environment variable is set
		envValue, envSet := os.LookupEnv(envKey)
		if !envSet {
			// now check if the file variant for the environment variable
			fileKey := fmt.Sprintf("%s_FILE", envKey)
			filePath, fileSet := os.LookupEnv(fileKey)
			if !fileSet {
				return nil, ErrRequiredVariableMissing
			}
			// now trim any excess spaces from the file path
			filePath = strings.TrimSpace(filePath)
			// now check if the file path is a empty string
			if filePath == "" {
				return nil, ErrFilePathEmpty
			}
			// now open the file
			file, err := os.Open(filePath)
			if err != nil {
				return nil, err
			}
			// and read its contents
			fileContentBytes, err := io.ReadAll(file)
			if err != nil {
				return nil, err
			}
			// now trim any excess spaces at the start and end of the file
			fileContents := strings.TrimSpace(string(fileContentBytes))
			// now set the value to the original environment key
			parsedVariables[envKey] = fileContents
			// handle the next entry
			continue
		}

		// trim excess spaces from the environment variable value
		envValue = strings.TrimSpace(envValue)
		// and now check if the value is empty
		if envValue == "" {
			return nil, ErrRequiredVariableEmpty
		}
		// now set the value
		parsedVariables[envKey] = envValue
	}

	// start by parsing the required environment variables
	for envKey, defaultValue := range c.Optional {
		// check if the environment variable is set
		envValue, envSet := os.LookupEnv(envKey)
		if !envSet {
			// now check if the file variant for the environment variable
			fileKey := fmt.Sprintf("%s_FILE", envKey)
			filePath, fileSet := os.LookupEnv(fileKey)
			if !fileSet {
				parsedVariables[envKey] = defaultValue
				continue
			}
			// now trim any excess spaces from the file path
			filePath = strings.TrimSpace(filePath)
			// now check if the file path is a empty string
			if filePath == "" {
				parsedVariables[envKey] = defaultValue
				continue
			}
			// now open the file
			file, err := os.Open(filePath)
			if err != nil {
				return nil, err
			}
			// and read its contents
			fileContentBytes, err := io.ReadAll(file)
			if err != nil {
				return nil, err
			}
			// now trim any excess spaces at the start and end of the file
			fileContents := strings.TrimSpace(string(fileContentBytes))
			// now set the value to the original environment key
			parsedVariables[envKey] = fileContents
			// handle the next entry
			continue
		}

		// trim excess spaces from the environment variable value
		envValue = strings.TrimSpace(envValue)
		// and now check if the value is empty
		if envValue == "" {
			parsedVariables[envKey] = defaultValue
			continue
		}
		// now set the value
		parsedVariables[envKey] = envValue
	}
	return parsedVariables, nil
}

// PopulateFromFile takes an already opened file handle and reads the files
// contents into the authorization configuration and returns an error if one
// occurs.
//
// This function is a shorthand call for json.NewDecoder(reader).Decode(target)
func (c *AuthorizationConfiguration) PopulateFromFile(f *os.File) error {
	return json.NewDecoder(f).Decode(c)
}

// PopulateFromFilePath takes a file path to the configuration file and
// populates the values using the PopulateFromFile function and returns
// an error if one occurrs.
func (c *AuthorizationConfiguration) PopulateFromFilePath(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	return c.PopulateFromFile(file)
}
