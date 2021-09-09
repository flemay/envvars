package envvars_test

import (
	"testing"

	"github.com/flemay/envvars/pkg/envvars"
	"github.com/flemay/envvars/pkg/yml"
)

func TestList_toReturnAllEnvvarsIfNoTagsSpecified(t *testing.T) {
	// given
	d := yml.NewDeclarationYML("testdata/list_declaration_file.yml")
	// when
	c, err := envvars.List(d)
	// then
	if err != nil {
		t.Errorf("want no error, got %q", err.Error())
		return
	}
	if len(c) != 3 {
		t.Errorf("want number of envvar to be 3, got %d", len(c))
	}
}

func TestList_toReturnTaggedEnvvarsIfTagsSpecified(t *testing.T) {
	// given
	reader := yml.NewDeclarationYML("testdata/list_declaration_file.yml")
	// when
	c, err := envvars.List(reader, "tag1")
	// then
	if err != nil {
		t.Errorf("want no error, got %q", err.Error())
		return
	}
	if len(c) != 2 {
		t.Errorf("want envvar collection 2, got %d", len(c))
	}
}

func TestList_toReturnErrorIfInvalidDeclarationAndTagNameList(t *testing.T) {
	// given
	reader := yml.NewDeclarationYML("testdata/declaration_file_invalid.yml")
	invalidList := helperInvalidTagNames()

	// when
	c, err := envvars.List(reader, invalidList...)

	// then
	want := helperReadFile(t, "testdata/declaration_file_with_tag_name_list_invalid_error_message.golden")
	if len(c) != 0 {
		t.Errorf("want empty envvar collection, got %d", len(c))
	}
	if err == nil {
		t.Errorf("want error %q, got none", want)
		return
	}
	if err.Error() != want {
		t.Errorf("want error %q, got %q", want, err.Error())
	}
}
