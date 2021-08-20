package yml_test

import (
	"os"
	"testing"

	"github.com/flemay/envvars/pkg/envvars"
	"github.com/flemay/envvars/pkg/yml"
	"github.com/stretchr/testify/assert"
)

func TestDeclarationYML_Read_toReturnDeclarationBasedOnDeclarationFile(t *testing.T) {
	// given
	declarationYML := yml.NewDeclarationYML("testdata/envvars.yml")

	// when
	d, err := declarationYML.Read()

	// then
	assert.NoError(t, err)
	assert.NotNil(t, d)
	expectedTags := envvars.TagCollection{
		&envvars.Tag{
			Name: "tag1",
			Desc: "desc of tag1",
		},
	}
	assert.EqualValues(t, expectedTags, d.Tags)

	expectedEnvvars := envvars.EnvvarCollection{
		&envvars.Envvar{
			Name: "ENVVAR_1",
			Desc: "desc of ENVVAR_1",
		},
		&envvars.Envvar{
			Name:     "ENVVAR_2",
			Desc:     "desc of ENVVAR_2",
			Optional: true,
			Example:  "example1",
		},
	}

	assert.EqualValues(t, expectedEnvvars, d.Envvars)
}
func TestDeclarationYML_Read_toReturnErrorIfMalformatedDeclarationFile(t *testing.T) {
	// given
	declarationFilePath := "testdata/envvars_malformated.yml"
	declarationYML := yml.NewDeclarationYML(declarationFilePath)

	// when
	d, err := declarationYML.Read()

	// then
	assert.Error(t, err)
	assert.Nil(t, d)
	assert.Contains(t, err.Error(), "error occurred when parsing the file "+declarationFilePath)
}

func TestDeclarationYML_Read_toReturnErrorIfFileNotFound(t *testing.T) {
	// given
	noSuchFilePath := "nosuchfile.yml"
	declarationYML := yml.NewDeclarationYML(noSuchFilePath)

	// when
	d, err := declarationYML.Read()

	// then
	assert.Error(t, err)
	assert.Nil(t, d)
	assert.Contains(t, err.Error(), "open nosuchfile.yml: no such file or directory")
}

func TestDeclarationYML_Write_toWriteDeclarationInYMLFile(t *testing.T) {
	// given
	filename := t.TempDir() + "/envvars.yml"
	writer := yml.NewDeclarationYML(filename)
	d := &envvars.Declaration{
		Envvars: []*envvars.Envvar{
			&envvars.Envvar{
				Name: "ENVVAR_1",
				Desc: "desc of ENVVAR_1",
			},
		},
	}

	// when
	err := writer.Write(d, false)

	// then
	assert.NoError(t, err)
	expectedFile := readFile(t, "testdata/envvars.yml.golden")
	actualFile := readFile(t, filename)
	assert.Equal(t, expectedFile, actualFile)
}

func TestDeclarationYML_Write_toReturnErrorIfFileExists(t *testing.T) {
	// given
	filename := t.TempDir() + "/envvars.yml"
	writer := yml.NewDeclarationYML(filename)
	d := &envvars.Declaration{
		Envvars: []*envvars.Envvar{
			&envvars.Envvar{
				Name: "ENVVAR_1",
				Desc: "desc of ENVVAR_1",
			},
		},
	}
	err := writer.Write(d, false)
	assert.NoError(t, err)

	// when
	err = writer.Write(d, false)

	// then
	assert.EqualError(t, err, "open "+filename+": file exists")
}

func TestDeclarationYML_Write_toOverwriteExistingFile(t *testing.T) {
	// given
	filename := t.TempDir() + "/envvars.yml"
	writer := yml.NewDeclarationYML(filename)
	d := &envvars.Declaration{
		Envvars: []*envvars.Envvar{
			&envvars.Envvar{
				Name: "ENVVAR_1",
				Desc: "desc of ENVVAR_1",
			},
		},
	}
	err := writer.Write(d, false)
	assert.NoError(t, err)

	// when
	err = writer.Write(d, true)

	// then
	assert.NoError(t, err)
	expectedFile := readFile(t, "testdata/envvars.yml.golden")
	actualFile := readFile(t, filename)
	assert.Equal(t, expectedFile, actualFile)
}

func readFile(t *testing.T, name string) string {
	f, err := os.ReadFile(name)
	if err != nil {
		t.Fatalf(err.Error())
	}
	return string(f)
}
