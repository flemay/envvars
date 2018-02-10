package envvars_test

import (
	"github.com/flemay/envvars/pkg/envvars"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestValidate_toReturnNoErrorIfValidMetadata(t *testing.T) {
	// given
	metadata, _ := envvars.NewMetadata("testdata/envvars.toml")
	// when
	err := envvars.Validate(metadata)
	// then
	assert.NoError(t, err)
}

func TestValidate_toReturnErrorIfInvalidMetadata(t *testing.T) {
	// given
	metadata, _ := envvars.NewMetadata("testdata/invalid_envvars.toml")
	// when
	err := envvars.Validate(metadata)
	// then
	expectedErrorMsg, _ := ioutil.ReadFile("testdata/validate_error_message.golden")
	assert.EqualError(t, err, string(expectedErrorMsg))
}
