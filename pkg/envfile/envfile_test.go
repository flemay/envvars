package envfile_test

import (
	"github.com/flemay/envvars/pkg/envfile"
	"github.com/flemay/envvars/pkg/yml"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestEnvfile_toGenerateFileIfItDoesNotExist(t *testing.T) {
	// given
	d, _ := yml.NewDeclaration("testdata/envfile_declaration_file.yml")
	name := "testdata/envfile_file.tmp"
	writer := envfile.NewEnvfile(name, false)

	// when
	err := writer.Write(d.Envvars)

	// then
	assert.NoError(t, err)
	expected := readFile(t, "testdata/envfile_file.golden")
	actual := readFile(t, name)
	assert.Equal(t, expected, actual)
	removeFileOrDir(t, name)
}

func TestEnvfile_toGenerateFileIfItExistsAndOverwrite(t *testing.T) {
	// given
	d, _ := yml.NewDeclaration("testdata/envfile_declaration_file.yml")
	name := "testdata/envfile_file.tmp"
	createEmptyFile(t, name)
	writer := envfile.NewEnvfile(name, true)

	// when
	err := writer.Write(d.Envvars)

	// then
	assert.NoError(t, err)
	expected := readFile(t, "testdata/envfile_file.golden")
	actual := readFile(t, name)
	assert.Equal(t, expected, actual)
	removeFileOrDir(t, name)
}

func TestEnvfile_toReturnErrorIfFileExistsAndNotOverwrite(t *testing.T) {
	// given
	d, _ := yml.NewDeclaration("testdata/envfile_declaration_file.yml")
	name := "testdata/envfile_file.tmp"
	createEmptyFile(t, name)
	writer := envfile.NewEnvfile(name, false)

	// when
	err := writer.Write(d.Envvars)

	// then
	assert.Error(t, err)
	assert.EqualError(t, err, "error: "+name+" already exist")
	removeFileOrDir(t, name)
}

func TestEnvfile_toReturnErrorIfPathIsFolderAndOverwrite(t *testing.T) {
	// given
	d, _ := yml.NewDeclaration("testdata/envfile_declaration_file.yml")
	name := "testdata/tmp"
	createDir(t, name)
	writer := envfile.NewEnvfile(name, true)

	// when
	err := writer.Write(d.Envvars)

	// then
	assert.Error(t, err)
	assert.EqualError(t, err, "error: "+name+" is a folder, not a file")
	removeFileOrDir(t, name)
}

func removeFileOrDir(t *testing.T, name string) {
	if err := os.Remove(name); err != nil {
		t.Fatalf(err.Error())
	}
}

func readFile(t *testing.T, name string) string {
	f, err := ioutil.ReadFile(name)
	if err != nil {
		t.Fatalf(err.Error())
	}
	return string(f)
}

func createEmptyFile(t *testing.T, name string) {
	f, err := os.Create(name)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if err := f.Close(); err != nil {
		t.Fatalf(err.Error())
	}
}

func createDir(t *testing.T, name string) {
	os.Mkdir(name, os.ModePerm)
}
