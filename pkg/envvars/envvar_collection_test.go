package envvars_test

import (
	"testing"

	"github.com/flemay/envvars/pkg/envvars"
)

func TestEnvvarCollectionGet(t *testing.T) {
	givenEnvvars := envvars.EnvvarCollection{
		&envvars.Envvar{Name: "ENVVAR_1"},
		&envvars.Envvar{Name: "ENVVAR_2"},
	}
	testCases := map[string]struct {
		givenCollection envvars.EnvvarCollection
		whenEnvvarName  string
		thenEnvvar      *envvars.Envvar
	}{
		"return envvar if match": {givenEnvvars, "ENVVAR_2", givenEnvvars[1]},
		"return nil if no match": {givenEnvvars, "NOT_DEFINED", nil},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// when
			got := tc.givenCollection.Get(tc.whenEnvvarName)
			// then
			if tc.thenEnvvar != nil {
				if got == nil {
					t.Error("want an envvar, got nil")
					return
				}
				if got.Name != tc.thenEnvvar.Name {
					t.Errorf("want envvar with name %q, got %q", tc.thenEnvvar.Name, got.Name)
				}
				return
			}
			if got != nil {
				t.Error("want nil, got not nil")
			}
		})
	}
}

func TestEnvvarCollectionGetAll(t *testing.T) {
	givenEnvvars := envvars.EnvvarCollection{
		&envvars.Envvar{Name: "ENVVAR_1"},
		&envvars.Envvar{Name: "ENVVAR_2"},
		&envvars.Envvar{Name: "ENVVAR_1"},
	}
	testCases := map[string]struct {
		givenCollection envvars.EnvvarCollection
		whenEnvvarName  string
		thenEnvvars     envvars.EnvvarCollection
	}{
		"return matching envvars":             {givenEnvvars, "ENVVAR_1", envvars.EnvvarCollection{givenEnvvars[0], givenEnvvars[2]}},
		"return empty collection if no match": {givenEnvvars, "NOT_DEFINED", envvars.EnvvarCollection{}},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// when
			got := tc.givenCollection.GetAll(tc.whenEnvvarName)
			// then
			if len(got) != len(tc.thenEnvvars) {
				t.Errorf("want %q, got %q", len(tc.thenEnvvars), len(got))
				return
			}
			for i, envvar := range got {
				if envvar.Name != tc.thenEnvvars[i].Name {
					t.Errorf("want %q, got %q", tc.thenEnvvars[i].Name, envvar.Name)
				}
			}
		})
	}
}

func TestEnvvarCollectionWithTag(t *testing.T) {
	givenEnvvars := envvars.EnvvarCollection{
		&envvars.Envvar{Name: "ENVVAR_1", Tags: []string{"T1", "T2"}},
		&envvars.Envvar{Name: "ENVVAR_2", Tags: []string{"T1"}},
		&envvars.Envvar{Name: "ENVVAR_3", Tags: []string{"T1", "T2"}},
	}
	testCases := map[string]struct {
		givenCollection envvars.EnvvarCollection
		whenTagName     string
		thenEnvvars     envvars.EnvvarCollection
	}{
		"return envvars with matching tag":    {givenEnvvars, "T2", envvars.EnvvarCollection{givenEnvvars[0], givenEnvvars[2]}},
		"return empty collection if no match": {givenEnvvars, "NOT_DEFINED", envvars.EnvvarCollection{}},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// when
			got := tc.givenCollection.WithTag(tc.whenTagName)
			// then
			if len(got) != len(tc.thenEnvvars) {
				t.Errorf("want %d, got %d", len(tc.thenEnvvars), len(got))
				return
			}
			for i, envvar := range got {
				if envvar.Name != tc.thenEnvvars[i].Name {
					t.Errorf("want %q, got %q", tc.thenEnvvars[i].Name, envvar.Name)
				}
			}
		})
	}
}
