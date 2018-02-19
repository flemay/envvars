package envvars_test

import (
	"github.com/flemay/envvars/pkg/envvars"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestEnsure_toReturnErrorIfInvalidDeclarationAndTagNameList(t *testing.T) {
	// given
	d, _ := envvars.NewDeclaration("testdata/declaration_file_invalid.toml")
	invalidList := givenInvalidTagNameList()

	// when
	err := envvars.Ensure(d, invalidList...)

	// then
	expectedErrorMsg := readFile(t, "testdata/declaration_file_with_tag_name_list_invalid_error_message.golden")
	assert.EqualError(t, err, expectedErrorMsg)
}

func TestEnsure_toReturnNoErrorIfEnvvarsComply(t *testing.T) {
	// given
	d, _ := envvars.NewDeclaration("testdata/ensure_declaration_file.toml")
	os.Setenv("ENVVAR_1", "name1")
	os.Setenv("ENVVAR_2", "name2")
	os.Setenv("ENVVAR_3", "name3")

	// when
	err := envvars.Ensure(d)

	// then
	assert.NoError(t, err)
	os.Unsetenv("ENVVAR_1")
	os.Unsetenv("ENVVAR_2")
	os.Unsetenv("ENVVAR_3")
}

func TestEnsure_toReturnErrorIfEnvvarsDoNotComply(t *testing.T) {
	// given
	d, _ := envvars.NewDeclaration("testdata/ensure_declaration_file.toml")
	os.Setenv("ENVVAR_2", "")

	// when
	err := envvars.Ensure(d)

	// then
	os.Unsetenv("ENVVAR_2")
	expectedErrorMsg := readFile(t, "testdata/ensure_error_message.golden")
	assert.EqualError(t, err, expectedErrorMsg)
}

func TestEnsure_toReturnNoErrorIfTaggedEnvvarsComply(t *testing.T) {
	// given
	d, _ := envvars.NewDeclaration("testdata/ensure_declaration_file.toml")
	os.Setenv("ENVVAR_2", "name2")
	// when
	err := envvars.Ensure(d, "tag2")
	// then
	assert.NoError(t, err)
	os.Unsetenv("ENVVAR_2")
}

func TestEnsure_toReturnErrorIfTaggedEnvvarsDoNotComply(t *testing.T) {
	// given
	d, _ := envvars.NewDeclaration("testdata/ensure_declaration_file.toml")
	os.Setenv("ENVVAR_2", "name2")
	// when
	err := envvars.Ensure(d, "tag1")
	// then
	assert.EqualError(t, err, "environment variable ENVVAR_1 is not defined")
}
