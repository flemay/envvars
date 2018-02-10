package envvars_test

import (
	"github.com/flemay/envvars/pkg/envvars"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDefinition_toReturnTheContentOfValidEnvvarsFile(t *testing.T) {
	// given
	envvarsFilePath := "testdata/envvars.toml"

	// when
	m, err := envvars.NewDefinition(envvarsFilePath)

	// then
	assert.NoError(t, err)
	assert.NotNil(t, m)
	assert.Len(t, m.Envvars, 2)
}
func TestNewDefinition_toReturnErrorIfMalformatedEnvvarsFile(t *testing.T) {
	// given
	envvarsFilePath := "testdata/malformated_envvars.toml"

	// when
	m, err := envvars.NewDefinition(envvarsFilePath)

	// then
	assert.Error(t, err)
	assert.Nil(t, m)
	assert.Contains(t, err.Error(), "error occurred when opening the file "+envvarsFilePath)
}
