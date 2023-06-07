package config

import (
	"os"
	"testing"
)

const environmentFilePath = "../res/environment.json"

func TestEnvironmentPopulateFromFile(t *testing.T) {
	file, err := os.Open(environmentFilePath)
	if err != nil {
		t.Fatal(err)
	}
	var c EnvironmentConfiguration
	err = c.PopulateFromFile(file)
	if err != nil {
		t.Fatal(err)
	}
	if c.Required[0] != "TEST_ENV_1" {
		t.Fatalf("expected required variable 1 to be 'TEST_ENV_1'. got: %s", c.Required[0])
	}
	if c.Required[1] != "TEST_ENV_2" {
		t.Fatalf("expected required variable 2 to be 'TEST_ENV_2'. got: %s", c.Required[1])
	}
	if c.Optional["TEST_ENV_3"] != "default" {
		t.Fatalf("expected optional variable 1 value to be 'default'. got: %s", c.Optional["TEST_ENV_3"])
	}
	if c.Optional["TEST_ENV_4"] != "" {
		t.Fatalf("expected optional variable 1 value to be ''. got: %s", c.Optional["TEST_ENV_3"])
	}
}

func TestEnvironmentPopulateFromFilePath(t *testing.T) {
	var c EnvironmentConfiguration
	err := c.PopulateFromFilePath(environmentFilePath)
	if err != nil {
		t.Fatal(err)
	}
	if c.Required[0] != "TEST_ENV_1" {
		t.Fatalf("expected required variable 1 to be 'TEST_ENV_1'. got: %s", c.Required[0])
	}
	if c.Required[1] != "TEST_ENV_2" {
		t.Fatalf("expected required variable 2 to be 'TEST_ENV_2'. got: %s", c.Required[1])
	}
	if c.Optional["TEST_ENV_3"] != "default" {
		t.Fatalf("expected optional variable 1 value to be 'default'. got: %s", c.Optional["TEST_ENV_3"])
	}
	if c.Optional["TEST_ENV_4"] != "" {
		t.Fatalf("expected optional variable 1 value to be ''. got: %s", c.Optional["TEST_ENV_3"])
	}
}

func TestParseEnvironment(t *testing.T) {
	var c EnvironmentConfiguration
	err := c.PopulateFromFilePath(environmentFilePath)
	os.Setenv("TEST_ENV_1", "1")
	os.Setenv("TEST_ENV_2", "2")
	if err != nil {
		t.Fatal(err)
	}

	_, err = c.ParseEnvironment()
	if err != nil {
		t.Fatal(err)
	}
}
