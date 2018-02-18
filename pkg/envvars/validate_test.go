package envvars_test

import (
	"github.com/flemay/envvars/pkg/envvars"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidate_toReturnNoErrorIfValidDeclaration(t *testing.T) {
	// given
	d, _ := envvars.NewDeclaration("testdata/validate_declaration_file.toml")
	// when
	err := envvars.Validate(d)
	// then
	assert.NoError(t, err)
}

func TestValidate_toReturnNoErrorIfValidDeclarationWithTags(t *testing.T) {
	// given
	d, _ := envvars.NewDeclaration("testdata/validate_declaration_file_with_tags.toml")
	// when
	err := envvars.Validate(d)
	// then
	assert.NoError(t, err)
}

func TestValidate_toReturnErrorIfInvalidDeclaration(t *testing.T) {
	// given
	d, _ := envvars.NewDeclaration("testdata/declaration_file_invalid.toml")
	// when
	err := envvars.Validate(d)
	// then
	expectedErrorMsg := readFile(t, "testdata/declaration_file_invalid_error_message.golden")
	assert.EqualError(t, err, expectedErrorMsg)
}

func TestValidate_toReturnErrorIfDeclarationIsEmpty(t *testing.T) {
	// given
	d, _ := envvars.NewDeclaration("testdata/declaration_file_empty.toml")
	// when
	err := envvars.Validate(d)
	// then
	assert.Error(t, err)
	assert.EqualError(t, err, "declaration must at least have 1 envvars")
}

func TestValidate_toReturnErrorIfDeclarationIsNil(t *testing.T) {
	// when
	err := envvars.Validate(nil)
	// then
	assert.EqualError(t, err, "declaration is nil")
}
