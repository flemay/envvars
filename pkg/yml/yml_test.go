package yml_test

import (
	"github.com/flemay/envvars/pkg/envvars"
	"github.com/flemay/envvars/pkg/yml"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDeclaration_toReturnDeclarationBasedOnDeclarationFile(t *testing.T) {
	// given
	declarationFilePath := "testdata/envvars.yml"

	// when
	d, err := yml.NewDeclaration(declarationFilePath)

	// then
	assert.NoError(t, err)
	assert.NotNil(t, d)
	expectedTags := envvars.TagCollection{
		&envvars.Tag{
			Name: "tag1",
			Desc: "desc of tag1",
		},
	}
	assert.EqualValues(t, expectedTags, d.Tags)

	expectedEnvvars := envvars.EnvvarCollection{
		&envvars.Envvar{
			Name: "ENVVAR_1",
			Desc: "desc of ENVVAR_1",
		},
		&envvars.Envvar{
			Name:     "ENVVAR_2",
			Desc:     "desc of ENVVAR_2",
			Optional: true,
			Example:  "example1",
		},
	}

	assert.EqualValues(t, expectedEnvvars, d.Envvars)
}
func TestNewDeclaration_toReturnErrorIfMalformatedDeclarationFile(t *testing.T) {
	// given
	declarationFilePath := "testdata/envvars_malformated.yml"

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
