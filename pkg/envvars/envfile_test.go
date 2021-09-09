package envvars_test

import (
	"testing"

	"github.com/flemay/envvars/pkg/envfile"
	"github.com/flemay/envvars/pkg/envvars"
	"github.com/flemay/envvars/pkg/yml"
)

func TestEnvfile_toReturnErrorIfInvalidDeclarationAndTagNameList(t *testing.T) {
	// given
	r := yml.NewDeclarationYML("testdata/declaration_file_invalid.yml")
	tags := helperInvalidTagNames()
	filename := t.TempDir() + "/.env"
	w := envfile.NewEnvfile(filename, false, false)

	// when
	got := envvars.Envfile(r, w, tags...)

	// then
	want := helperReadFile(t, "testdata/declaration_file_with_tag_name_list_invalid_error_message.golden")
	if got.Error() != want {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestEnvfile_toWriteEnvfile(t *testing.T) {
	// given
	r := yml.NewDeclarationYML("testdata/envfile_declaration_file.yml")
	filename := t.TempDir() + "/.env"
	w := envfile.NewEnvfile(filename, false, false)

	// when
	err := envvars.Envfile(r, w)

	// then
	if err != nil {
		t.Errorf("want no error, got %q", err.Error())
		return
	}
	got := helperReadFile(t, filename)
	want := helperReadFile(t, "./testdata/envfile_file.golden")
	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}
func TestEnvfile_toWriteEnvfileWithOnlySpecifiedTags(t *testing.T) {
	// given
	r := yml.NewDeclarationYML("testdata/envfile_declaration_file.yml")
	filename := t.TempDir() + "/.env"
	w := envfile.NewEnvfile(filename, false, false)

	// when
	err := envvars.Envfile(r, w, "tag1")

	// then
	if err != nil {
		t.Errorf("want no error, got %q", err.Error())
		return
	}
	got := helperReadFile(t, filename)
	want := helperReadFile(t, "./testdata/envfile_file_with_tag.golden")
	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}
