package envvars_test

import (
	"os"
	"testing"
)

func givenInvalidTagNameList() []string {
	return []string{"tagNotThere", "tagDuplicated", "tagDuplicated", ""}
}

func readFile(t *testing.T, name string) string {
	f, err := os.ReadFile(name)
	if err != nil {
		t.Fatalf(err.Error())
	}
	return string(f)
}
