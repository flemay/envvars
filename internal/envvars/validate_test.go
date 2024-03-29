package envvars_test

import (
	"testing"

	"github.com/flemay/envvars/internal/envvars"
	"github.com/flemay/envvars/internal/mock"
	"github.com/flemay/envvars/internal/yml"
)

func TestValidate_toReturnNoErrorIfValidDeclaration(t *testing.T) {
	// given
	r := yml.NewDeclarationYML("testdata/validate_declaration_file.yml")
	// when
	got := envvars.Validate(r)
	// then
	if got != nil {
		t.Errorf("want no error, got %q", got.Error())
	}
}

func TestValidate_toReturnNoErrorIfValidDeclarationWithTags(t *testing.T) {
	// given
	r := yml.NewDeclarationYML("testdata/validate_declaration_file_with_tags.yml")
	// when
	got := envvars.Validate(r)
	// then
	if got != nil {
		t.Errorf("want no error, got %q", got.Error())
	}
}

func TestValidate_toReturnErrorIfInvalidDeclaration(t *testing.T) {
	// given
	r := yml.NewDeclarationYML("testdata/declaration_file_invalid.yml")
	// when
	got := envvars.Validate(r)
	// then
	want := helperReadFile(t, "testdata/declaration_file_invalid_error_message.golden")
	if got.Error() != want {
		t.Errorf("want %q, got %q", want, got.Error())
	}
}

func TestValidate_toReturnErrorIfDeclarationIsEmpty(t *testing.T) {
	// given
	r := yml.NewDeclarationYML("testdata/declaration_file_empty.yml")
	// when
	got := envvars.Validate(r)
	// then
	want := "declaration must at least have 1 envvars"
	if got.Error() != want {
		t.Errorf("want %q, got %q", want, got.Error())
	}
}

func TestValidate_toReturnErrorIfDeclarationIsNil(t *testing.T) {
	// given
	r := mock.DeclarationReader{}
	// when
	got := envvars.Validate(r)
	// then
	if got == nil {
		t.Error("want error, got none")
		return
	}
	want := "declaration is nil"
	if got.Error() != want {
		t.Errorf("want %q, got %q", want, got)
	}
}
