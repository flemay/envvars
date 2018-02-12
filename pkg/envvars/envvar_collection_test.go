package envvars_test

import (
	"github.com/flemay/envvars/pkg/envvars"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnvvarCollection_Get_toReturnEnvvarMatchingName(t *testing.T) {
	// given
	c := envvars.EnvvarCollection{
		&envvars.Envvar{Name: "ENVVAR_1"},
		&envvars.Envvar{Name: "ENVVAR_2"},
	}
	// when
	ev := c.Get("ENVVAR_2")
	// then
	assert.NotNil(t, ev)
	assert.Equal(t, "ENVVAR_2", ev.Name)
}

func TestEnvvarCollection_Get_toReturnNilIfNoEnvvarMatchingName(t *testing.T) {
	// given
	c := envvars.EnvvarCollection{
		&envvars.Envvar{Name: "ENVVAR_1"},
		&envvars.Envvar{Name: "ENVVAR_2"},
	}
	// when
	ev := c.Get("NOT_DEFINED")
	// then
	assert.Nil(t, ev)
}

func TestEnvvarCollection_GetAll_toReturnEnvvarCollectionMatchingName(t *testing.T) {
	// given
	c := envvars.EnvvarCollection{
		&envvars.Envvar{Name: "ENVVAR_1"},
		&envvars.Envvar{Name: "ENVVAR_2"},
		&envvars.Envvar{Name: "ENVVAR_1"},
	}
	// when
	evs := c.GetAll("ENVVAR_1")
	// then
	assert.Len(t, evs, 2)
	assert.Equal(t, "ENVVAR_1", evs[0].Name)
	assert.Equal(t, "ENVVAR_1", evs[1].Name)
}

func TestEnvvarCollection_GetAll_toReturnEmtpyEnvvarCollectionIfNoneMatchingName(t *testing.T) {
	// given
	c := envvars.EnvvarCollection{
		&envvars.Envvar{Name: "ENVVAR_1"},
		&envvars.Envvar{Name: "ENVVAR_2"},
		&envvars.Envvar{Name: "ENVVAR_1"},
	}
	// when
	evs := c.GetAll("NOT_DEFINED")
	// then
	assert.Len(t, evs, 0)
}

func TestEnvvarCollection_WithTag_toReturnEnvvarsWithTagName(t *testing.T) {
	// given
	c := envvars.EnvvarCollection{
		&envvars.Envvar{Name: "ENVVAR_1", Tags: []string{"T1", "T2"}},
		&envvars.Envvar{Name: "ENVVAR_2", Tags: []string{"T1"}},
		&envvars.Envvar{Name: "ENVVAR_3", Tags: []string{"T1", "T2"}},
	}
	// when
	evs := c.WithTag("T2")
	// then
	assert.Len(t, evs, 2)
	assert.Equal(t, "ENVVAR_1", evs[0].Name)
	assert.Equal(t, "ENVVAR_3", evs[1].Name)
}

func TestEnvvarCollection_WithTag_toReturnEmptyIfNoEnvvarWithTagName(t *testing.T) {
	// given
	c := envvars.EnvvarCollection{
		&envvars.Envvar{Name: "ENVVAR_1", Tags: []string{"T1", "T2"}},
		&envvars.Envvar{Name: "ENVVAR_2", Tags: []string{"T1"}},
		&envvars.Envvar{Name: "ENVVAR_3", Tags: []string{"T1", "T2"}},
	}
	// when
	evs := c.WithTag("NOT_THERE")
	// then
	assert.Len(t, evs, 0)
}
