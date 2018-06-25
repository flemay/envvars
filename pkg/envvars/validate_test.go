package envvars_test

import (
	"github.com/flemay/envvars/pkg/envvars"
	"github.com/flemay/envvars/pkg/mocks"
	"github.com/flemay/envvars/pkg/yml"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidate_toReturnNoErrorIfValidDeclaration(t *testing.T) {
	// given
	reader := yml.NewDeclarationYML("testdata/validate_declaration_file.yml")
	// when
	err := envvars.Validate(reader)
	// then
	assert.NoError(t, err)
}

func TestValidate_toReturnNoErrorIfValidDeclarationWithTags(t *testing.T) {
	// given
	reader := yml.NewDeclarationYML("testdata/validate_declaration_file_with_tags.yml")
	// when
	err := envvars.Validate(reader)
	// then
	assert.NoError(t, err)
}

func TestValidate_toReturnErrorIfInvalidDeclaration(t *testing.T) {
	// given
	reader := yml.NewDeclarationYML("testdata/declaration_file_invalid.yml")
	// when
	err := envvars.Validate(reader)
	// then
	expectedErrorMsg := readFile(t, "testdata/declaration_file_invalid_error_message.golden")
	assert.EqualError(t, err, expectedErrorMsg)
}

func TestValidate_toReturnErrorIfDeclarationIsEmpty(t *testing.T) {
	// given
	reader := yml.NewDeclarationYML("testdata/declaration_file_empty.yml")
	// when
	err := envvars.Validate(reader)
	// then
	assert.Error(t, err)
	assert.EqualError(t, err, "declaration must at least have 1 envvars")
}

func TestValidate_toReturnErrorIfDeclarationIsNil(t *testing.T) {
	// given
	mockReader := new(mocks.DeclarationReader)
	mockReader.On("Read").Return(nil, nil)

	// when
	err := envvars.Validate(mockReader)
	// then
	assert.EqualError(t, err, "declaration is nil")
}
