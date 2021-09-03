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
