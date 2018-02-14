package envvars_test

import (
	"github.com/flemay/envvars/pkg/envvars"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDefinition_toReturnDefinitionBasedOnEnvvarsFile(t *testing.T) {
	// given
	envvarsFilePath := "testdata/definition_envvars.toml"

	// when
	definition, err := envvars.NewDefinition(envvarsFilePath)

	// then
	assert.NoError(t, err)
	assert.NotNil(t, definition)
	assert.Len(t, definition.Envvars, 2)
}
func TestNewDefinition_toReturnErrorIfMalformatedEnvvarsFile(t *testing.T) {
	// given
	envvarsFilePath := "testdata/definition_malformated_envvars.toml"

	// when
	definition, err := envvars.NewDefinition(envvarsFilePath)

	// then
	assert.Error(t, err)
	assert.Nil(t, definition)
	assert.Contains(t, err.Error(), "error occurred when opening the file "+envvarsFilePath)
}

func TestNewDefinition_toReturnErrorIfFileNotFound(t *testing.T) {
	// given
	noSuchFilePath := "nosuchfile.toml"

	// when
	definition, err := envvars.NewDefinition(noSuchFilePath)

	// then
	assert.Error(t, err)
	assert.Nil(t, definition)
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
	hasTag := ev.HasTag("TAG_NOT_THERE")
	// then
	assert.False(t, hasTag)
}
