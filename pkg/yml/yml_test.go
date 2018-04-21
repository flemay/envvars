package yml_test

import (
	"github.com/flemay/envvars/pkg/yml"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDeclaration_toReturnDeclarationBasedOnDeclarationFile(t *testing.T) {
	// given
	declarationFilePath := "testdata/declaration_file.yml"

	// when
	d, err := yml.NewDeclaration(declarationFilePath)

	// then
	assert.NoError(t, err)
	assert.NotNil(t, d)
	assert.Len(t, d.Envvars, 2)
}
func TestNewDeclaration_toReturnErrorIfMalformatedDeclarationFile(t *testing.T) {
	// given
	declarationFilePath := "testdata/declaration_file_malformated.yml"

	// when
	d, err := yml.NewDeclaration(declarationFilePath)

	// then
	assert.Error(t, err)
	assert.Nil(t, d)
	assert.Contains(t, err.Error(), "error occurred when parsing the file "+declarationFilePath)
}

func TestNewDeclaration_toReturnErrorIfFileNotFound(t *testing.T) {
	// given
	noSuchFilePath := "nosuchfile.yml"

	// when
	d, err := yml.NewDeclaration(noSuchFilePath)

	// then
	assert.Error(t, err)
	assert.Nil(t, d)
	assert.Contains(t, err.Error(), "open nosuchfile.yml: no such file or directory")
}
