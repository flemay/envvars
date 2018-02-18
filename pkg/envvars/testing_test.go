package envvars_test

import (
	"io/ioutil"
	"os"
	"testing"
)

func givenInvalidTagNameList() []string {
	return []string{"tagNotThere", "tagDuplicated", "tagDuplicated", ""}
}

func expectedInvalidTagNameListErrorMessage() string {
	return "tag 'tagNotThere' is not declared; tag 'tagDuplicated' is not declared; tag 'tagDuplicated' is not declared; tag '' is empty; tag 'tagDuplicated' is duplicated"
}

func removeFileOrDir(t *testing.T, name string) {
	if err := os.Remove(name); err != nil {
		t.Fatalf(err.Error())
	}
}

func readFile(t *testing.T, name string) string {
	f, err := ioutil.ReadFile(name)
	if err != nil {
		t.Fatalf(err.Error())
	}
	return string(f)
}

func createEmptyFile(t *testing.T, name string) {
	f, err := os.Create(name)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if err := f.Close(); err != nil {
		t.Fatalf(err.Error())
	}
}

func createDir(t *testing.T, name string) {
	os.Mkdir(name, os.ModePerm)
}
