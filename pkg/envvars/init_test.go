package envvars_test

import (
	"testing"

	"github.com/flemay/envvars/pkg/envvars"
	"github.com/flemay/envvars/pkg/yml"
)

func TestInit_toCreateDeclarationFile(t *testing.T) {
	// given
	filename := t.TempDir() + "/envvars.yml"
	w := yml.NewDeclarationYML(filename)

	// when
	err := envvars.Init(w)

	// then
	if err != nil {
		t.Errorf("want no error, got %q", err.Error())
		return
	}

	want := readFile(t, "./testdata/init_declaration_file.golden")
	got := readFile(t, filename)
	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestInit_toReturnErrorIfDeclarationExists(t *testing.T) {
	// given
	filename := t.TempDir() + "/envvars.yml"
	w := yml.NewDeclarationYML(filename)
	err := envvars.Init(w)
	if err != nil {
		t.Fatalf("envvars.Init: %v", err)
	}

	// when
	err = envvars.Init(w)

	// then
	if err == nil {
		t.Error("want error, got none")
		return
	}
	want := "open " + filename + ": file exists"
	if err.Error() != want {
		t.Errorf("want error %q, got %q", want, err.Error())
	}
}
