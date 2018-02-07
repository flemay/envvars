package envvars_test

import (
	"github.com/flemay/envvars"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMetadata_toReturnTheContentOfValidEnvvarsFile(t *testing.T) {
	// given
	envvarsFilePath := "testdata/envvars.toml"

	// when
	m, err := envvars.NewMetadata(envvarsFilePath)

	// then
	assert.NoError(t, err)
	assert.NotNil(t, m)
	assert.Len(t, m.Envvars, 2)
}
func TestNewMetadata_toReturnErrorIfMalformatedEnvvarsFile(t *testing.T) {
	// given
	envvarsFilePath := "testdata/malformated_envvars.toml"

	// when
	m, err := envvars.NewMetadata(envvarsFilePath)

	// then
	assert.Error(t, err)
	assert.Nil(t, m)
	assert.Contains(t, err.Error(), "error occurred when opening the file "+envvarsFilePath)
}
