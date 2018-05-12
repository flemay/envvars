package envfile_test

import (
	"github.com/flemay/envvars/pkg/envfile"
	"github.com/flemay/envvars/pkg/envvars"
	"github.com/flemay/envvars/pkg/yml"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestEnvfile_toImplementEnvfileWriterInterface(t *testing.T) {
	assert.Implements(t, (*envvars.EnvfileWriter)(nil), new(envfile.Envfile))
}

func TestEnvfile_toGenerateFileIfItDoesNotExist(t *testing.T) {
	// given
	d, _ := yml.NewDeclaration("testdata/envvars.yml")
	name := "testdata/envfile.tmp"
	writer := envfile.NewEnvfile(name, false, false)

	// when
	err := writer.Write(d.Envvars)

	// then
	assert.NoError(t, err)
	expected := readFile(t, "testdata/envfile.golden")
	actual := readFile(t, name)
	assert.Equal(t, expected, actual)
	removeFileOrDir(t, name)
}

func TestEnvfile_toGenerateFileWithExampleIfItDoesNotExist(t *testing.T) {
	// given
	d, _ := yml.NewDeclaration("testdata/envvars.yml")
	name := "testdata/envfile.tmp"
	writer := envfile.NewEnvfile(name, true, false)

	// when
	err := writer.Write(d.Envvars)

	// then
	assert.NoError(t, err)
	expected := readFile(t, "testdata/envfile_example.golden")
	actual := readFile(t, name)
	assert.Equal(t, expected, actual)
	removeFileOrDir(t, name)
}

func TestEnvfile_toGenerateFileIfItExistsAndOverwrite(t *testing.T) {
	// given
	d, _ := yml.NewDeclaration("testdata/envvars.yml")
	name := "testdata/envfile.tmp"
	createEmptyFile(t, name)
	writer := envfile.NewEnvfile(name, false, true)

	// when
	err := writer.Write(d.Envvars)

	// then
	assert.NoError(t, err)
	expected := readFile(t, "testdata/envfile.golden")
	actual := readFile(t, name)
	assert.Equal(t, expected, actual)
	removeFileOrDir(t, name)
}

func TestEnvfile_toReturnErrorIfFileExistsAndNotOverwrite(t *testing.T) {
	// given
	d, _ := yml.NewDeclaration("testdata/envvars.yml")
	name := "testdata/envfile.tmp"
	createEmptyFile(t, name)
	writer := envfile.NewEnvfile(name, false, false)

	// when
	err := writer.Write(d.Envvars)

	// then
	assert.Error(t, err)
	assert.EqualError(t, err, "error: "+name+" already exist")
	removeFileOrDir(t, name)
}

func TestEnvfile_toReturnErrorIfPathIsFolderAndOverwrite(t *testing.T) {
	// given
	d, _ := yml.NewDeclaration("testdata/envvars.yml")
	name := "testdata/tmp"
	createDir(t, name)
	writer := envfile.NewEnvfile(name, false, true)

	// when
	err := writer.Write(d.Envvars)

	// then
	assert.Error(t, err)
	assert.EqualError(t, err, "error: "+name+" is a folder, not a file")
	removeFileOrDir(t, name)
}

func TestEnvfile_toRemoveFileIfItExists(t *testing.T) {
	// given
	name := "testdata/envfile.tmp"
	createEmptyFile(t, name)

	// when
	err := envfile.Remove(name)

	// then
	assert.NoError(t, err)
}

func TestEnvfile_Remove_toReturnErrorIfEnvfileNotPresent(t *testing.T) {
	// given
	name := "testdata/envfile.tmp"

	// when
	err := envfile.Remove(name)

	// then
	assert.Error(t, err)
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
