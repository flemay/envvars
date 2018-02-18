package envvars_test

import (
	"github.com/flemay/envvars/pkg/envvars"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList_toReturnAllEnvvarsIfNoTagsSpecified(t *testing.T) {
	// given
	d, _ := envvars.NewDeclaration("testdata/list_declaration_file.toml")
	// when
	c, err := envvars.List(d)
	// then
	assert.NoError(t, err)
	assert.Len(t, c, 3)
}

func TestList_toReturnTaggedEnvvarsIfTagsSpecified(t *testing.T) {
	// given
	d, _ := envvars.NewDeclaration("testdata/list_declaration_file.toml")
	// when
	c, err := envvars.List(d, "tag1")
	// then
	assert.NoError(t, err)
	assert.Len(t, c, 2)
}

func TestList_toReturnErrorIfInvalidDeclarationAndTagNameList(t *testing.T) {
	// given
	d, _ := envvars.NewDeclaration("testdata/declaration_file_invalid.toml")
	invalidList := givenInvalidTagNameList()

	// when
	c, err := envvars.List(d, invalidList...)

	// then
	expectedErrorMsg := readFile(t, "testdata/declaration_file_with_tag_name_list_invalid_error_message.golden")
	assert.EqualError(t, err, expectedErrorMsg)
	assert.Len(t, c, 0)
}
