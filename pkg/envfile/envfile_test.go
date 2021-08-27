package envfile_test

import (
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/flemay/envvars/pkg/envfile"
	"github.com/flemay/envvars/pkg/envvars"
	"github.com/flemay/envvars/pkg/yml"
)

func TestEnvfile_toImplementEnvfileWriterInterface(t *testing.T) {
	var _ envvars.EnvfileWriter = new(envfile.Envfile)
}

func TestEnvfileWrite(t *testing.T) {
	var tests = []struct {
		name                   string
		givenDeclarationFile   string
		givenEnvfileExists     bool
		givenEnvfileIsDir      bool
		whenExample            bool
		whenOverwrite          bool
		thenGoldenFile         string
		thenErrorMessageRegExp string
	}{
		{"generate file", "./testdata/envvars.yml", false, false, false, false, "./testdata/envfile.golden", ""},
		{"generate file with example", "./testdata/envvars.yml", false, false, true, false, "./testdata/envfile_example.golden", ""},
		{"overwrite existing file", "./testdata/envvars.yml", true, false, false, true, "./testdata/envfile.golden", ""},
		{"fail generate when existing file", "./testdata/envvars.yml", true, false, false, false, "", `.* already exist$`},
		{"fail overwrite a directory", "./testdata/envvars.yml", false, true, false, true, "", `.* is a folder, not a file$`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			reader := yml.NewDeclarationYML(tt.givenDeclarationFile)
			d, err := reader.Read()
			if err != nil {
				t.Fatalf("Reader.Read: %v", err)
			}

			if tt.givenEnvfileExists && tt.givenEnvfileIsDir {
				t.Fatalf("givenEnvfileExists and givenEnvfileIsDir cannot be both true")
			}

			envfileName := t.TempDir()
			if !tt.givenEnvfileIsDir {
				envfileName = envfileName + "/.env"
			}

			if tt.givenEnvfileExists {
				createEmptyFile(t, envfileName)
			}
			writer := envfile.NewEnvfile(envfileName, tt.whenExample, tt.whenOverwrite)

			// when
			if err := writer.Write(d.Envvars); err != nil {
				if tt.thenErrorMessageRegExp == "" {
					t.Fatalf("Writer.Write: %v", err)
				}
				// then
				re := regexp.MustCompile(tt.thenErrorMessageRegExp)
				if !re.MatchString(err.Error()) {
					t.Errorf("want regexp %q to match %q", tt.thenErrorMessageRegExp, err.Error())
				}
			}

			//then
			if tt.thenGoldenFile != "" {
				want := readFile(t, tt.thenGoldenFile)
				got := readFile(t, envfileName)
				if want != got {
					t.Errorf("want %q, got %q", want, got)
				}
			}
		})
	}
}

func TestEnvfile_toRemoveFileIfItExists(t *testing.T) {
	// given
	name := t.TempDir() + "/.env"
	createEmptyFile(t, name)

	// when
	err := envfile.Remove(name)

	// then
	if err != nil {
		t.Errorf("want no error, got %q", err.Error())
	}
}

func TestEnvfile_Remove_toReturnErrorIfEnvfileNotPresent(t *testing.T) {
	// given
	name := "testdata/.env"

	// when
	err := envfile.Remove(name)

	// then
	if err == nil {
		t.Error("want error, got no error")
	}
}

// readFile reads a file and returns it as string. It also removes trailing EOL.
func readFile(t *testing.T, name string) string {
	f, err := os.ReadFile(name)
	if err != nil {
		t.Fatal(err.Error())
	}
	return strings.TrimSuffix(string(f), "\n")
}

func createEmptyFile(t *testing.T, name string) {
	f, err := os.Create(name)
	if err != nil {
		t.Fatal(err.Error())
	}
	if err := f.Close(); err != nil {
		t.Fatal(err.Error())
	}
}
