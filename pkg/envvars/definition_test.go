package envvars_test

import (
	"github.com/flemay/envvars/pkg/envvars"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDefinition_toReturnDefinitionBasedOnEnvvarsFile(t *testing.T) {
	// given
	envvarsFilePath := "testdata/envvars.toml"

	// when
	definition, err := envvars.NewDefinition(envvarsFilePath)

	// then
	assert.NoError(t, err)
	assert.NotNil(t, definition)
	assert.Len(t, definition.Envvars, 2)
}
func TestNewDefinition_toReturnErrorIfMalformatedEnvvarsFile(t *testing.T) {
	// given
	envvarsFilePath := "testdata/malformated_envvars.toml"

	// when
	definition, err := envvars.NewDefinition(envvarsFilePath)

	// then
	assert.Error(t, err)
	assert.Nil(t, definition)
	assert.Contains(t, err.Error(), "error occurred when opening the file "+envvarsFilePath)
}

func TestNewDefinition_toReturnErrorIfFileNotFound(t *testing.T) {
	// given
	noSuchFilePath := "nosuchfile.toml"

	// when
	definition, err := envvars.NewDefinition(noSuchFilePath)

	// then
	assert.Error(t, err)
	assert.Nil(t, definition)
	assert.Contains(t, err.Error(), "open nosuchfile.toml: no such file or directory")
}

func TestNewDefinitionAndValidate_toReturnDefinitionBasedOnValidEnvvarsFile(t *testing.T) {
	// given
	envvarsFilePath := "testdata/envvars.toml"

	// when
	definition, err := envvars.NewDefinitionAndValidate(envvarsFilePath)

	// then
	assert.NoError(t, err)
	assert.NotNil(t, definition)
	assert.Len(t, definition.Envvars, 2)
}

func TestNewDefinitionAndValidate_toReturnErrorIfInvalidEnvvarsFile(t *testing.T) {
	// given
	envvarsFilePath := "testdata/invalid_envvars.toml"

	// when
	definition, err := envvars.NewDefinitionAndValidate(envvarsFilePath)

	// then
	assert.Error(t, err)
	assert.NotNil(t, definition)
	expectedErrorMsg := readFile(t, "testdata/validate_error_message.golden")
	assert.EqualError(t, err, expectedErrorMsg)
}

func TestNewDefinitionAndValidate_toReturnErrorIfMalformatedEnvvarsFile(t *testing.T) {
	// given
	envvarsFilePath := "testdata/malformated_envvars.toml"

	// when
	definition, err := envvars.NewDefinitionAndValidate(envvarsFilePath)

	// then
	assert.Error(t, err)
	assert.Nil(t, definition)
	assert.Contains(t, err.Error(), "error occurred when opening the file "+envvarsFilePath)
}

func TestNewDefinitionAndValidate_toReturnErrorIfFileNotFound(t *testing.T) {
	// given
	noSuchFilePath := "nosuchfile.toml"

	// when
	definition, err := envvars.NewDefinitionAndValidate(noSuchFilePath)

	// then
	assert.Error(t, err)
	assert.Nil(t, definition)
	assert.Contains(t, err.Error(), "open nosuchfile.toml: no such file or directory")
}
