package envvars_test

import (
	"testing"

	"github.com/flemay/envvars/pkg/envvars"
)

func TestEnvvarHasTag(t *testing.T) {
	testCases := map[string]struct {
		givenTags  []string
		whenTag    string
		thenHasTag bool
	}{
		"true if tag is present":      {[]string{"T1", "T2"}, "T1", true},
		"false if tag is not present": {[]string{"T1", "T2"}, "T3", false},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// given
			ev := envvars.Envvar{Tags: tc.givenTags}
			// when
			got := ev.HasTag(tc.whenTag)
			// then
			if got != tc.thenHasTag {
				t.Errorf("want %t, got %t", tc.thenHasTag, got)
			}
		})
	}
}
