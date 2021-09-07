package envvars_test

import (
	"testing"

	"github.com/flemay/envvars/pkg/envfile"
	"github.com/flemay/envvars/pkg/envvars"
	"github.com/flemay/envvars/pkg/mocks"
	"github.com/flemay/envvars/pkg/yml"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestEnvfile_toReturnErrorIfInvalidDeclarationAndTagNameList(t *testing.T) {
	// given
	r := yml.NewDeclarationYML("testdata/declaration_file_invalid.yml")
	tags := givenInvalidTagNameList()
	filename := t.TempDir() + "/.env"
	w := envfile.NewEnvfile(filename, false, false)

	// when
	got := envvars.Envfile(r, w, tags...)

	// then
	want := readFile(t, "testdata/declaration_file_with_tag_name_list_invalid_error_message.golden")
	if got.Error() != want {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestEnvfile_toWriteEnvfile(t *testing.T) {
	// given
	reader := yml.NewDeclarationYML("testdata/envfile_declaration_file.yml")
	mockWriter := new(mocks.EnvfileWriter)
	mockWriter.On("Write", mock.Anything).Return(nil)

	// when
	err := envvars.Envfile(reader, mockWriter)

	// then
	assert.NoError(t, err)
	d, _ := reader.Read()
	mockWriter.AssertCalled(t, "Write", d.Envvars)
}
func TestEnvfile_toWriteEnvfileWithOnlySpecifiedTags(t *testing.T) {
	// given
	reader := yml.NewDeclarationYML("testdata/envfile_declaration_file.yml")
	mockWriter := new(mocks.EnvfileWriter)
	mockWriter.On("Write", mock.Anything).Return(nil)

	// when
	err := envvars.Envfile(reader, mockWriter, "tag1")

	// then
	d, _ := reader.Read()
	assert.NoError(t, err)
	mockWriter.AssertCalled(t, "Write", d.Envvars.WithTag("tag1"))
}
