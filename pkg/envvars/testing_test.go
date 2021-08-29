package envvars_test

import (
	"os"
	"strings"
	"testing"
)

func givenInvalidTagNameList() []string {
	return []string{"tagNotThere", "tagDuplicated", "tagDuplicated", ""}
}

// readFile reads a file and returns it as string. It also removes trailing EOL.
func readFile(t *testing.T, name string) string {
	f, err := os.ReadFile(name)
	if err != nil {
		t.Fatalf(err.Error())
	}
	return strings.TrimSuffix(string(f), "\n")
}

type keyValue struct {
	Key   string
	Value string
}

// helperSetenv sets environment variables and returns an unset func to be called
// during defer.
func helperSetenv(t *testing.T, kvs ...keyValue) func() {
	for _, kv := range kvs {
		if err := os.Setenv(kv.Key, kv.Value); err != nil {
			t.Fatalf("os.Setenv: %v", err)
		}
	}
	return func() {
		for _, kv := range kvs {
			if err := os.Unsetenv(kv.Key); err != nil {
				t.Fatalf("os.Unsetenv: %v", err)
			}
		}
	}
}
