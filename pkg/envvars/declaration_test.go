package envvars_test

import (
	"github.com/flemay/envvars/pkg/envvars"
	"github.com/stretchr/testify/assert"
	"testing"
)

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
