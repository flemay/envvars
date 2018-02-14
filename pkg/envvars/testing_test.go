package envvars_test

import (
	"io/ioutil"
	"os"
	"testing"
)

func givenInvalidTagNameList() []string {
	return []string{"TAG_NOT_THERE", "TAG_DUPLICATED", "TAG_DUPLICATED", ""}
}

func expectedInvalidTagNameListErrorMessage() string {
	return "tag 'TAG_NOT_THERE' is not defined; tag 'TAG_DUPLICATED' is not defined; tag 'TAG_DUPLICATED' is not defined; tag '' is empty; tag 'TAG_DUPLICATED' is duplicated"
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
