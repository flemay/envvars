package yml_test

import (
	"os"
	"strings"
	"testing"

	"github.com/flemay/envvars/internal/envvars"
	"github.com/flemay/envvars/internal/yml"
)

func TestDeclarationYML_implementsDeclarationReader(t *testing.T) {
	var _ envvars.DeclarationReader = new(yml.DeclarationYML)
}

func TestDeclarationYML_Read(t *testing.T) {
	thenDeclaration := envvars.Declaration{
		Tags: envvars.TagCollection{
			{
				Name: "tag1",
				Desc: "desc of tag1",
			},
		},
		Envvars: envvars.EnvvarCollection{
			{
				Name: "ENVVAR_1",
				Desc: "desc of ENVVAR_1",
			},
			{
				Name:     "ENVVAR_2",
				Desc:     "desc of ENVVAR_2",
				Optional: true,
				Example:  "example1",
			},
		},
	}
	testCases := map[string]struct {
		givenDeclarationFile string
		thenDeclaration      *envvars.Declaration
		thenErrorSubMessage  string
	}{
		"returns declaration":       {"./testdata/envvars.yml", &thenDeclaration, ""},
		"error if malformated file": {"./testdata/envvars_malformated.yml", nil, "error occurred when parsing the file"},
		"error if file not found":   {"nosuchfile.yml", nil, "no such file or directory"},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// given
			declarationYML := yml.NewDeclarationYML(tc.givenDeclarationFile)

			// when
			got, err := declarationYML.Read()

			// then
			if err != nil {
				if tc.thenErrorSubMessage == "" {
					t.Errorf("want no error, got %s", err.Error())
				} else if !strings.Contains(err.Error(), tc.thenErrorSubMessage) {
					t.Errorf("want %q to be in error %q", tc.thenErrorSubMessage, err.Error())
				}
				return
			}

			if got == nil {
				t.Error("want value, got nil")
				return
			}

			if !tc.thenDeclaration.Equal(*got) {
				t.Errorf("want %+v, got %+v", *tc.thenDeclaration, *got)
			}
		})
	}
}

func TestDeclarationYML_implementsDeclarationWriter(t *testing.T) {
	var _ envvars.DeclarationWriter = new(yml.DeclarationYML)
}

func TestDeclarationYML_Write(t *testing.T) {
	givenDeclaration := &envvars.Declaration{
		Envvars: []*envvars.Envvar{
			{
				Name: "ENVVAR_1",
				Desc: "desc of ENVVAR_1",
			},
		},
	}
	testCases := map[string]struct {
		givenDeclaration           *envvars.Declaration
		givenDeclarationFileExists bool
		whenOverwrite              bool
		thenDeclarationFile        string
		thenErrorSubMessage        string
	}{
		"write to declaration file":  {givenDeclaration, false, false, "./testdata/envvars.yml.golden", ""},
		"write fails if file exists": {givenDeclaration, true, false, "", ": file exists"},
		"overwrites existing file":   {givenDeclaration, true, true, "./testdata/envvars.yml.golden", ""},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// given
			filename := t.TempDir() + "/envvars.yml"
			if tc.givenDeclarationFileExists {
				helperCreateEmptyFile(t, filename)
			}
			writer := yml.NewDeclarationYML(filename)

			// when
			err := writer.Write(tc.givenDeclaration, tc.whenOverwrite)

			// then
			if err != nil {
				if tc.thenErrorSubMessage == "" {
					t.Errorf("want no error, got %q", err.Error())
				} else if !strings.Contains(err.Error(), tc.thenErrorSubMessage) {
					t.Errorf("want %q to be in error %q", tc.thenErrorSubMessage, err.Error())
				}
				return
			}

			want := helperReadFile(t, tc.thenDeclarationFile)
			got := helperReadFile(t, filename)
			if want != got {
				t.Errorf("want %s, got %s", want, got)
			}
		})
	}
}

func helperReadFile(t *testing.T, name string) string {
	f, err := os.ReadFile(name)
	if err != nil {
		t.Fatalf(err.Error())
	}
	return string(f)
}

func helperCreateEmptyFile(t *testing.T, name string) {
	f, err := os.Create(name)
	if err != nil {
		t.Fatal(err.Error())
	}
	if err := f.Close(); err != nil {
		t.Fatal(err.Error())
	}
}
