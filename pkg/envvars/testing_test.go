package envvars_test

import (
	"io/ioutil"
	"testing"
)

func givenInvalidTagNameList() []string {
	return []string{"tagNotThere", "tagDuplicated", "tagDuplicated", ""}
}

func expectedInvalidTagNameListErrorMessage() string {
	return "tag 'tagNotThere' is not declared; tag 'tagDuplicated' is not declared; tag 'tagDuplicated' is not declared; tag '' is empty; tag 'tagDuplicated' is duplicated"
}

func readFile(t *testing.T, name string) string {
	f, err := ioutil.ReadFile(name)
	if err != nil {
		t.Fatalf(err.Error())
	}
	return string(f)
}
