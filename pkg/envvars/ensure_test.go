package envvars_test

import (
	"github.com/flemay/envvars/pkg/envvars"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestEnsure_toReturnNoErrorIfEnvvarsAreComplyToDefinition(t *testing.T) {
	// given
	definition, _ := envvars.NewDefinition("testdata/envvars.toml")
	os.Setenv("NAME_1", "name1")
	os.Setenv("NAME_2", "name2")

	// when
	err := envvars.Ensure(definition)

	// then
	assert.NoError(t, err)
	os.Unsetenv("NAME_1")
	os.Unsetenv("NAME_2")
}

func TestEnsure_toReturnErrorIfEnvvarsDoNotComplyToDefinition(t *testing.T) {
	// given
	definition, _ := envvars.NewDefinition("testdata/envvars.toml")

	// when
	err := envvars.Ensure(definition)

	// then
	expectedErrorMsg, _ := ioutil.ReadFile("testdata/ensure_error_message.golden")
	assert.EqualError(t, err, string(expectedErrorMsg))
}
