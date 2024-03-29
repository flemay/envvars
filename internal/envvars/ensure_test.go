package envvars_test

import (
	"testing"

	"github.com/flemay/envvars/internal/envvars"
	"github.com/flemay/envvars/internal/yml"
)

func TestEnsure_toReturnErrorIfInvalidDeclarationAndTagNameList(t *testing.T) {
	// given
	reader := yml.NewDeclarationYML("testdata/declaration_file_invalid.yml")

	// when
	got := envvars.Ensure(reader, helperInvalidTagNames()...)

	// then
	want := helperReadFile(t, "testdata/declaration_file_with_tag_name_list_invalid_error_message.golden")
	if got.Error() != want {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestEnsure_toReturnNoErrorIfEnvvarsComply(t *testing.T) {
	// given
	reader := yml.NewDeclarationYML("testdata/ensure_declaration_file.yml")
	t.Setenv("ENVVAR_1", "name1")
	t.Setenv("ENVVAR_2", "name2")
	t.Setenv("ENVVAR_3", "name3")
	t.Setenv("OPTIONAL_ENVVAR", "optional")

	// when
	err := envvars.Ensure(reader)

	// then
	if err != nil {
		t.Errorf("want no error, got %q", err.Error())
	}
}

func TestEnsure_toReturnNoErrorIfOptionalEnvvarIsNotDefined(t *testing.T) {
	// given
	reader := yml.NewDeclarationYML("testdata/ensure_declaration_file.yml")
	t.Setenv("ENVVAR_1", "name1")
	t.Setenv("ENVVAR_2", "name2")
	t.Setenv("ENVVAR_3", "name3")

	// when
	got := envvars.Ensure(reader)

	// then
	if got != nil {
		t.Errorf("want no error, got %q", got.Error())
	}
}

func TestEnsure_toReturnNoErrorIfOptionalEnvvarHasEmptyValue(t *testing.T) {
	reader := yml.NewDeclarationYML("testdata/ensure_declaration_file.yml")
	t.Setenv("ENVVAR_1", "name1")
	t.Setenv("ENVVAR_2", "name2")
	t.Setenv("ENVVAR_3", "name3")
	t.Setenv("OPTIONAL_ENVVAR", "")

	// when
	err := envvars.Ensure(reader)

	// then
	if err != nil {
		t.Errorf("want no error, got %q", err.Error())
	}
}

func TestEnsure_toReturnErrorIfEnvvarsDoNotComply(t *testing.T) {
	// given
	reader := yml.NewDeclarationYML("testdata/ensure_declaration_file.yml")
	t.Setenv("ENVVAR_2", "")

	// when
	got := envvars.Ensure(reader)

	// then
	want := helperReadFile(t, "testdata/ensure_error_message.golden")
	if got.Error() != want {
		t.Errorf("want %q, got %q", want, got.Error())
	}
}

func TestEnsure_toReturnNoErrorIfTaggedEnvvarsComply(t *testing.T) {
	// given
	reader := yml.NewDeclarationYML("testdata/ensure_declaration_file.yml")
	t.Setenv("ENVVAR_2", "name2")

	// when
	got := envvars.Ensure(reader, "tag2")

	// then
	if got != nil {
		t.Errorf("want no error, got %q", got.Error())
	}
}

func TestEnsure_toReturnErrorIfTaggedEnvvarsDoNotComply(t *testing.T) {
	// given
	reader := yml.NewDeclarationYML("testdata/ensure_declaration_file.yml")
	t.Setenv("ENVVAR_2", "name2")

	// when
	got := envvars.Ensure(reader, "tag1")

	// then
	want := "environment variable ENVVAR_1 is not defined. Variable description: Desc of ENVVAR_1"
	if got.Error() != want {
		t.Errorf("want %q, got %q", want, got.Error())
	}
}
