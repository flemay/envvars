package envvars_test

import (
	"os"
	"testing"

	"github.com/flemay/envvars/pkg/envvars"
	"github.com/flemay/envvars/pkg/mocks"
	"github.com/flemay/envvars/pkg/yml"
	"github.com/stretchr/testify/assert"
)

func TestEnsure_toReturnErrorIfInvalidDeclarationAndTagNameList(t *testing.T) {
	// given
	reader := yml.NewDeclarationYML("testdata/declaration_file_invalid.yml")
	invalidList := givenInvalidTagNameList()

	// when
	err := envvars.Ensure(reader, invalidList...)

	// then
	expectedErrorMsg := readFile(t, "testdata/declaration_file_with_tag_name_list_invalid_error_message.golden")
	assert.EqualError(t, err, expectedErrorMsg)
}

func TestEnsure_toReturnNoErrorIfEnvvarsComply(t *testing.T) {
	// given
	reader := yml.NewDeclarationYML("testdata/ensure_declaration_file.yml")
	defer helperSetenv(t,
		keyValue{"ENVVAR_1", "name1"},
		keyValue{"ENVVAR_2", "name2"},
		keyValue{"ENVVAR_3", "name3"})()

	// when
	err := envvars.Ensure(reader)

	// then
	assert.NoError(t, err)
}

func TestEnsure_toReturnNoErrorIfOptionalEnvvarIsNotDefined(t *testing.T) {
	// given
	d := &envvars.Declaration{
		Envvars: []*envvars.Envvar{
			&envvars.Envvar{
				Name:     "NAME",
				Desc:     "Desc",
				Optional: true,
			},
		},
	}
	mockReader := new(mocks.DeclarationReader)
	mockReader.On("Read").Return(d, nil)

	// when
	err := envvars.Ensure(mockReader)

	// then
	assert.NoError(t, err)
}

func TestEnsure_toReturnNoErrorIfOptionalEnvvarHasEmptyValue(t *testing.T) {
	// given
	d := &envvars.Declaration{
		Envvars: []*envvars.Envvar{
			&envvars.Envvar{
				Name:     "NAME",
				Desc:     "Desc",
				Optional: true,
			},
		},
	}
	mockReader := new(mocks.DeclarationReader)
	mockReader.On("Read").Return(d, nil)
	os.Setenv("NAME", "")

	// when
	err := envvars.Ensure(mockReader)

	// then
	os.Unsetenv("NAME")
	assert.NoError(t, err)
}

func TestEnsure_toReturnErrorIfEnvvarsDoNotComply(t *testing.T) {
	// given
	reader := yml.NewDeclarationYML("testdata/ensure_declaration_file.yml")
	os.Setenv("ENVVAR_2", "")
	defer os.Unsetenv("ENVVAR_2")

	// when
	got := envvars.Ensure(reader)

	// then
	want := readFile(t, "testdata/ensure_error_message.golden")
	assert.EqualError(t, got, want)
}

func TestEnsure_toReturnNoErrorIfTaggedEnvvarsComply(t *testing.T) {
	// given
	reader := yml.NewDeclarationYML("testdata/ensure_declaration_file.yml")
	os.Setenv("ENVVAR_2", "name2")
	// when
	err := envvars.Ensure(reader, "tag2")
	// then
	assert.NoError(t, err)
	os.Unsetenv("ENVVAR_2")
}

func TestEnsure_toReturnErrorIfTaggedEnvvarsDoNotComply(t *testing.T) {
	// given
	reader := yml.NewDeclarationYML("testdata/ensure_declaration_file.yml")
	os.Setenv("ENVVAR_2", "name2")
	// when
	err := envvars.Ensure(reader, "tag1")
	// then
	assert.EqualError(t, err, "environment variable ENVVAR_1 is not defined. Variable description: Desc of ENVVAR_1")
}
