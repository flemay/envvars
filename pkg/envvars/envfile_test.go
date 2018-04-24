package envvars_test

import (
	"github.com/flemay/envvars/pkg/envvars"
	"github.com/flemay/envvars/pkg/mocks"
	"github.com/flemay/envvars/pkg/yml"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestEnvfile_toReturnErrorIfInvalidDeclarationAndTagNameList(t *testing.T) {
	// given
	d, _ := yml.NewDeclaration("testdata/declaration_file_invalid.yml")
	invalidList := givenInvalidTagNameList()
	mockWriter := new(mocks.EnvfileWriter)
	// when
	err := envvars.Envfile(d, mockWriter, invalidList...)

	// then
	expectedErrorMsg := readFile(t, "testdata/declaration_file_with_tag_name_list_invalid_error_message.golden")
	assert.EqualError(t, err, expectedErrorMsg)
}

func TestEnvfile_toWriteEnvfile(t *testing.T) {
	// given
	d, _ := yml.NewDeclaration("testdata/envfile_declaration_file.yml")
	mockWriter := new(mocks.EnvfileWriter)
	mockWriter.On("Write", mock.Anything).Return(nil)

	// when
	err := envvars.Envfile(d, mockWriter)

	// then
	assert.NoError(t, err)
	mockWriter.AssertCalled(t, "Write", d.Envvars)
}
func TestEnvfile_toWriteEnvfileWithOnlySpecifiedTags(t *testing.T) {
	// given
	d, _ := yml.NewDeclaration("testdata/envfile_declaration_file.yml")
	mockWriter := new(mocks.EnvfileWriter)
	mockWriter.On("Write", mock.Anything).Return(nil)

	// when
	err := envvars.Envfile(d, mockWriter, "tag1")

	// then
	assert.NoError(t, err)
	mockWriter.AssertCalled(t, "Write", d.Envvars.WithTag("tag1"))
}
