package envvars_test

import (
	"github.com/flemay/envvars/pkg/envvars"
	"io/ioutil"
	"os"
	"testing"
)

func givenDefinition(t *testing.T, name string) *envvars.Definition {
	definition, err := envvars.NewDefinitionAndValidate("testdata/" + name)
	if err != nil {
		t.Fatalf(err.Error())
	}
	return definition
}

func removeFile(t *testing.T, name string) {
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
