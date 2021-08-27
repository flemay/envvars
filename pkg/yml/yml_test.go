package yml_test

import (
	"os"
	"strings"
	"testing"

	"github.com/flemay/envvars/pkg/envvars"
	"github.com/flemay/envvars/pkg/yml"
	"github.com/stretchr/testify/assert"
)

func TestDelcarationYML_Read(t *testing.T) {
	thenDeclaration := envvars.Declaration{
		Tags: envvars.TagCollection{
			&envvars.Tag{
				Name: "tag1",
				Desc: "desc of tag1",
			},
		},
		Envvars: envvars.EnvvarCollection{
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
		},
	}
	var tests = []struct {
		name                 string
		givenDeclarationFile string
		thenDeclaration      *envvars.Declaration
		thenErrorSubMessage  string
	}{
		{"returns declaration", "./testdata/envvars.yml", &thenDeclaration, ""},
		{"error if malformated file", "./testdata/envvars_malformated.yml", nil, "error occurred when parsing the file"},
		{"error if file not found", "nosuchfile.yml", nil, "no such file or directory"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			declarationYML := yml.NewDeclarationYML(tt.givenDeclarationFile)

			// when
			got, err := declarationYML.Read()

			// then
			if err != nil {
				if tt.thenErrorSubMessage == "" {
					t.Fatalf("Reader.Read: %v", err)
				}
				if !strings.Contains(err.Error(), tt.thenErrorSubMessage) {
					t.Errorf("want %q to be in error %q", tt.thenErrorSubMessage, err.Error())
				}
				return
			}

			if got == nil {
				t.Error("want value, got nil")
				return
			}

			if !tt.thenDeclaration.Equal(*got) {
				t.Error("want true, got false")
			}
		})
	}
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
