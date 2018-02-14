package envvars_test

import (
	"github.com/flemay/envvars/pkg/envvars"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidate_toReturnNoErrorIfValidDefinition(t *testing.T) {
	// given
	definition, _ := envvars.NewDefinition("testdata/validate_envvars.toml")
	// when
	err := envvars.Validate(definition)
	// then
	assert.NoError(t, err)
}

func TestValidate_toReturnNoErrorIfValidDefinitionWithTags(t *testing.T) {
	// given
	definition, _ := envvars.NewDefinition("testdata/validate_envvars_with_tags.toml")
	// when
	err := envvars.Validate(definition)
	// then
	assert.NoError(t, err)
}

func TestValidate_toReturnErrorIfInvalidDefinition(t *testing.T) {
	// given
	definition, _ := envvars.NewDefinition("testdata/invalid_envvars.toml")
	// when
	err := envvars.Validate(definition)
	// then
	expectedErrorMsg := readFile(t, "testdata/invalid_envvars_error_message.golden")
	assert.EqualError(t, err, expectedErrorMsg)
}

func TestValidate_toReturnErrorIfDefinitionIsEmpty(t *testing.T) {
	// given
	definition, _ := envvars.NewDefinition("testdata/validate_empty_envvars_file.toml")
	// when
	err := envvars.Validate(definition)
	// then
	assert.Error(t, err)
}
