package envvars_test

import (
	"github.com/flemay/envvars/pkg/envvars"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestValidate_toReturnNoErrorIfValidDefinition(t *testing.T) {
	// given
	definition, _ := envvars.NewDefinition("testdata/envvars.toml")
	// when
	err := envvars.Validate(definition)
	// then
	assert.NoError(t, err)
}

func TestValidate_toReturnErrorIfInvalidDefinition(t *testing.T) {
	// given
	definition, _ := envvars.NewDefinition("testdata/invalid_envvars.toml")
	// when
	err := envvars.Validate(definition)
	// then
	expectedErrorMsg, _ := ioutil.ReadFile("testdata/validate_error_message.golden")
	assert.EqualError(t, err, string(expectedErrorMsg))
}
