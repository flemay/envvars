package envvars_test

import (
	"github.com/flemay/envvars"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestEnsure_toReturnNoErrorIfEnvvarsAreComformToMetadata(t *testing.T) {
	// given
	metadata, _ := envvars.NewMetadata("testdata/envvars.toml")
	os.Setenv("NAME_1", "name1")
	os.Setenv("NAME_2", "name2")

	// when
	err := envvars.Ensure(metadata)

	// then
	assert.NoError(t, err)
	os.Unsetenv("NAME_1")
	os.Unsetenv("NAME_2")
}

func TestEnsure_toReturnErrorIfEnvvarsAreNotComformToMetadata(t *testing.T) {
	// given
	metadata, _ := envvars.NewMetadata("testdata/envvars.toml")

	// when
	err := envvars.Ensure(metadata)

	// then
	expectedErrorMsg, _ := ioutil.ReadFile("testdata/ensure_error_message.golden")
	assert.EqualError(t, err, string(expectedErrorMsg))
}
