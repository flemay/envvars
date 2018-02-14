package envvars_test

import (
	"github.com/flemay/envvars/pkg/envvars"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestEnsure_toReturnErrorIfInvalidDefinitionAndTagNameList(t *testing.T) {
	// given
	d, _ := envvars.NewDefinition("testdata/invalid_envvars.toml")
	invalidList := givenInvalidTagNameList()

	// when
	err := envvars.Ensure(d, invalidList...)

	// then
	expectedErrorMsg := readFile(t, "testdata/invalid_envvars_with_tag_name_list_error_message.golden")
	assert.EqualError(t, err, expectedErrorMsg)
}

func TestEnsure_toReturnNoErrorIfEnvvarsComply(t *testing.T) {
	// given
	definition, _ := envvars.NewDefinition("testdata/ensure_envvars.toml")
	os.Setenv("ENVVAR_1", "name1")
	os.Setenv("ENVVAR_2", "name2")

	// when
	err := envvars.Ensure(definition)

	// then
	assert.NoError(t, err)
	os.Unsetenv("ENVVAR_1")
	os.Unsetenv("ENVVAR_2")
}

func TestEnsure_toReturnErrorIfEnvvarsDoNotComply(t *testing.T) {
	// given
	definition, _ := envvars.NewDefinition("testdata/ensure_envvars.toml")

	// when
	err := envvars.Ensure(definition)

	// then
	expectedErrorMsg := readFile(t, "testdata/ensure_error_message.golden")
	assert.EqualError(t, err, expectedErrorMsg)
}

func TestEnsure_toReturnNoErrorIfTaggedEnvvarsComply(t *testing.T) {
	// given
	d, _ := envvars.NewDefinition("testdata/ensure_envvars.toml")
	os.Setenv("ENVVAR_2", "name2")
	// when
	err := envvars.Ensure(d, "TAG_2")
	// then
	assert.NoError(t, err)
	os.Unsetenv("ENVVAR_2")
}

func TestEnsure_toReturnErrorIfTaggedEnvvarsDoNotComply(t *testing.T) {
	// given
	d, _ := envvars.NewDefinition("testdata/ensure_envvars.toml")
	os.Setenv("ENVVAR_2", "name2")
	// when
	err := envvars.Ensure(d, "TAG_1")
	// then
	assert.EqualError(t, err, "environment variable ENVVAR_1 must be set")
}
