package envvars_test

import (
	"github.com/flemay/envvars/pkg/envvars"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnvfile_toReturnErrorIfInvalidDeclarationAndTagNameList(t *testing.T) {
	// given
	d, _ := envvars.NewDeclaration("testdata/declaration_file_invalid.yml")
	invalidList := givenInvalidTagNameList()

	// when
	err := envvars.Envfile(d, "", false, invalidList...)

	// then
	expectedErrorMsg := readFile(t, "testdata/declaration_file_with_tag_name_list_invalid_error_message.golden")
	assert.EqualError(t, err, expectedErrorMsg)
}

func TestEnvfile_toGenerateFileIfItDoesNotExist(t *testing.T) {
	// given
	d, _ := envvars.NewDeclaration("testdata/envfile_declaration_file.yml")
	name := "testdata/envfile_file.tmp"

	// when
	err := envvars.Envfile(d, name, false)

	// then
	assert.NoError(t, err)
	expected := readFile(t, "testdata/envfile_file.golden")
	actual := readFile(t, name)
	assert.Equal(t, expected, actual)
	removeFileOrDir(t, name)
}

func TestEnvfile_toGenerateFileWithOnlySpecifiedTags(t *testing.T) {
	// given
	d, _ := envvars.NewDeclaration("testdata/envfile_declaration_file.yml")
	name := "testdata/envfile_file_with_tag.tmp"

	// when
	err := envvars.Envfile(d, name, false, "tag1")

	// then
	assert.NoError(t, err)
	expected := readFile(t, "testdata/envfile_file_with_tag.golden")
	actual := readFile(t, name)
	assert.Equal(t, expected, actual)
	removeFileOrDir(t, name)
}

func TestEnvfile_toGenerateFileIfItExistsAndOverwrite(t *testing.T) {
	// given
	d, _ := envvars.NewDeclaration("testdata/envfile_declaration_file.yml")
	name := "testdata/envfile_file.tmp"
	createEmptyFile(t, name)

	// when
	err := envvars.Envfile(d, name, true)

	// then
	assert.NoError(t, err)
	expected := readFile(t, "testdata/envfile_file.golden")
	actual := readFile(t, name)
	assert.Equal(t, expected, actual)
	removeFileOrDir(t, name)
}

func TestEnvfile_toReturnErrorIfFileExistsAndNotOverwrite(t *testing.T) {
	// given
	d, _ := envvars.NewDeclaration("testdata/envfile_declaration_file.yml")
	name := "testdata/envfile_file.tmp"
	createEmptyFile(t, name)

	// when
	err := envvars.Envfile(d, name, false)

	// then
	assert.Error(t, err)
	assert.EqualError(t, err, "error: "+name+" already exist")
	removeFileOrDir(t, name)
}

func TestEnvfile_toReturnErrorIfPathIsFolderAndOverwrite(t *testing.T) {
	// given
	d, _ := envvars.NewDeclaration("testdata/envfile_declaration_file.yml")
	name := "testdata/tmp"
	createDir(t, name)

	// when
	err := envvars.Envfile(d, name, true)

	// then
	assert.Error(t, err)
	assert.EqualError(t, err, "error: "+name+" is a folder, not a file")
	removeFileOrDir(t, name)
}
