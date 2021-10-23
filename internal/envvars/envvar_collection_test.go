package envvars_test

import (
	"testing"

	"github.com/flemay/envvars/internal/envvars"
)

func TestEnvvarCollectionGet(t *testing.T) {
	collection := envvars.EnvvarCollection{
		&envvars.Envvar{Name: "ENVVAR_1"},
		&envvars.Envvar{Name: "ENVVAR_2"},
	}
	testCases := map[string]struct {
		givenCollection envvars.EnvvarCollection
		whenEnvvarName  string
		thenEnvvar      *envvars.Envvar
	}{
		"return envvar if match": {collection, "ENVVAR_2", collection[1]},
		"return nil if no match": {collection, "NOT_DEFINED", nil},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// when
			got := tc.givenCollection.Get(tc.whenEnvvarName)
			// then
			if tc.thenEnvvar != nil {
				if got == nil {
					t.Errorf("want an envvar with name %q, got nil", tc.thenEnvvar.Name)
					return
				}
				if got.Name != tc.thenEnvvar.Name {
					t.Errorf("want envvar with name %q, got %q", tc.thenEnvvar.Name, got.Name)
				}
				return
			}
			if got != nil {
				t.Errorf("want nil, got envvar with name %q", got.Name)
			}
		})
	}
}

func TestEnvvarCollectionGetAll(t *testing.T) {
	collection := envvars.EnvvarCollection{
		&envvars.Envvar{Name: "ENVVAR_1"},
		&envvars.Envvar{Name: "ENVVAR_2"},
		&envvars.Envvar{Name: "ENVVAR_1"},
	}
	testCases := map[string]struct {
		givenCollection envvars.EnvvarCollection
		whenEnvvarName  string
		thenEnvvars     envvars.EnvvarCollection
	}{
		"return matching envvars":             {collection, "ENVVAR_1", envvars.EnvvarCollection{collection[0], collection[2]}},
		"return empty collection if no match": {collection, "NOT_DEFINED", envvars.EnvvarCollection{}},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// when
			got := tc.givenCollection.GetAll(tc.whenEnvvarName)
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

func TestEnvvarCollectionWithTag(t *testing.T) {
	collection := envvars.EnvvarCollection{
		&envvars.Envvar{Name: "ENVVAR_1", Tags: []string{"T1", "T2"}},
		&envvars.Envvar{Name: "ENVVAR_2", Tags: []string{"T1"}},
		&envvars.Envvar{Name: "ENVVAR_3", Tags: []string{"T1", "T2"}},
	}
	testCases := map[string]struct {
		givenCollection envvars.EnvvarCollection
		whenTagName     string
		thenEnvvars     envvars.EnvvarCollection
	}{
		"return envvars with matching tag":    {collection, "T2", envvars.EnvvarCollection{collection[0], collection[2]}},
		"return empty collection if no match": {collection, "NOT_DEFINED", envvars.EnvvarCollection{}},
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
