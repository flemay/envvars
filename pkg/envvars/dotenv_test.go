package envvars_test

import (
	"github.com/flemay/envvars/pkg/envvars"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDotenv_toGenerateFileIfItDoesNotExist(t *testing.T) {
	// given
	d := givenDefinition(t, "dotenv_envvars.toml")
	dotenvPath := "testdata/dotenv_file.tmp"

	// when
	err := envvars.Dotenv(d, dotenvPath, false)

	// then
	assert.NoError(t, err)
	expected := readFile(t, "testdata/dotenv_file.golden")
	actual := readFile(t, dotenvPath)
	assert.Equal(t, expected, actual)
	removeFileOrDir(t, dotenvPath)
}

func TestDotenv_toGenerateFileIfItExistsAndOverwrite(t *testing.T) {
	// given
	d := givenDefinition(t, "dotenv_envvars.toml")
	dotenvPath := "testdata/dotenv_file.tmp"
	createEmptyFile(t, dotenvPath)

	// when
	err := envvars.Dotenv(d, dotenvPath, true)

	// then
	assert.NoError(t, err)
	expected := readFile(t, "testdata/dotenv_file.golden")
	actual := readFile(t, dotenvPath)
	assert.Equal(t, expected, actual)
	removeFileOrDir(t, dotenvPath)
}

func TestDotenv_toReturnErrorIfFileExistsAndNotOverwrite(t *testing.T) {
	// given
	d := givenDefinition(t, "dotenv_envvars.toml")
	dotenvPath := "testdata/dotenv_file.tmp"
	createEmptyFile(t, dotenvPath)

	// when
	err := envvars.Dotenv(d, dotenvPath, false)

	// then
	assert.Error(t, err)
	assert.EqualError(t, err, "error: "+dotenvPath+" already exist")
	removeFileOrDir(t, dotenvPath)
}

func TestDotenv_toReturnErrorIfPathIsFolderAndOverwrite(t *testing.T) {
	// given
	d := givenDefinition(t, "dotenv_envvars.toml")
	dotenvPath := "testdata/tmp"
	createDir(t, dotenvPath)

	// when
	err := envvars.Dotenv(d, dotenvPath, true)

	// then
	assert.Error(t, err)
	assert.EqualError(t, err, "error: "+dotenvPath+" is a folder, not a file")
	removeFileOrDir(t, dotenvPath)
}
