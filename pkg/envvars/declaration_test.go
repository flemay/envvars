package envvars_test

import (
	"github.com/flemay/envvars/pkg/envvars"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDeclaration_toReturnDeclarationBasedOnDeclarationFile(t *testing.T) {
	// given
	declarationFilePath := "testdata/declaration_file.toml"

	// when
	d, err := envvars.NewDeclaration(declarationFilePath)

	// then
	assert.NoError(t, err)
	assert.NotNil(t, d)
	assert.Len(t, d.Envvars, 2)
}
func TestNewDeclaration_toReturnErrorIfMalformatedDeclarationFile(t *testing.T) {
	// given
	declarationFilePath := "testdata/malformated_declaration_file.toml"

	// when
	d, err := envvars.NewDeclaration(declarationFilePath)

	// then
	assert.Error(t, err)
	assert.Nil(t, d)
	assert.Contains(t, err.Error(), "error occurred when opening the file "+declarationFilePath)
}

func TestNewDeclaration_toReturnErrorIfFileNotFound(t *testing.T) {
	// given
	noSuchFilePath := "nosuchfile.toml"

	// when
	d, err := envvars.NewDeclaration(noSuchFilePath)

	// then
	assert.Error(t, err)
	assert.Nil(t, d)
	assert.Contains(t, err.Error(), "open nosuchfile.toml: no such file or directory")
}

func TestEnvvar_HasTag_toReturnTrueIfTagIsPresent(t *testing.T) {
	// given
	ev := envvars.Envvar{Tags: []string{"T1", "T2"}}
	// when
	hasTag := ev.HasTag("T1")
	// then
	assert.True(t, hasTag)
}

func TestEnvvar_HasTag_toReturnFalseIfTagIsNotPresent(t *testing.T) {
	// given
	ev := envvars.Envvar{Tags: []string{"T1", "T2"}}
	// when
	hasTag := ev.HasTag("tagNotThere")
	// then
	assert.False(t, hasTag)
}
