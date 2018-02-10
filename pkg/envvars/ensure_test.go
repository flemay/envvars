package envvars_test

import (
	"github.com/flemay/envvars/pkg/envvars"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestEnsure_toReturnNoErrorIfEnvvarsAreComplyWithDefinition(t *testing.T) {
	// given
	definition, _ := envvars.NewDefinition("testdata/ensure_envvars.toml")
	os.Setenv("NAME_1", "name1")
	os.Setenv("NAME_2", "name2")

	// when
	err := envvars.Ensure(definition)

	// then
	assert.NoError(t, err)
	os.Unsetenv("NAME_1")
	os.Unsetenv("NAME_2")
}

func TestEnsure_toReturnErrorIfEnvvarsDoNotComplyWithDefinition(t *testing.T) {
	// given
	definition, _ := envvars.NewDefinition("testdata/ensure_envvars.toml")

	// when
	err := envvars.Ensure(definition)

	// then
	expectedErrorMsg := readFile(t, "testdata/ensure_error_message.golden")
	assert.EqualError(t, err, expectedErrorMsg)
}
