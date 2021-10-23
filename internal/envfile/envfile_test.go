package envfile_test

import (
	"os"
	"strings"
	"testing"

	"github.com/flemay/envvars/internal/envfile"
	"github.com/flemay/envvars/internal/envvars"
	"github.com/flemay/envvars/internal/yml"
)

func TestEnvfile_implementsEnvfileWriter(t *testing.T) {
	var _ envvars.EnvfileWriter = new(envfile.Envfile)
}

func TestEnvfileWrite(t *testing.T) {
	testCases := map[string]struct {
		givenDeclarationFile string
		givenEnvfileExists   bool
		givenEnvfileIsDir    bool
		whenExample          bool
		whenOverwrite        bool
		thenGoldenFile       string
		thenErrorSubMessage  string
	}{
		"generate file":                    {"./testdata/envvars.yml", false, false, false, false, "./testdata/envfile.golden", ""},
		"generate file with example":       {"./testdata/envvars.yml", false, false, true, false, "./testdata/envfile_example.golden", ""},
		"overwrite existing file":          {"./testdata/envvars.yml", true, false, false, true, "./testdata/envfile.golden", ""},
		"fail generate when existing file": {"./testdata/envvars.yml", true, false, false, false, "", "already exist"},
		"fail overwrite a directory":       {"./testdata/envvars.yml", false, true, false, true, "", "is a folder, not a file"},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// given
			reader := yml.NewDeclarationYML(tc.givenDeclarationFile)
			d, err := reader.Read()
			if err != nil {
				t.Fatalf("Reader.Read: %v", err)
			}

			if tc.givenEnvfileExists && tc.givenEnvfileIsDir {
				t.Fatalf("givenEnvfileExists and givenEnvfileIsDir cannot be both true")
			}

			envfileName := t.TempDir()
			if !tc.givenEnvfileIsDir {
				envfileName = envfileName + "/.env"
			}

			if tc.givenEnvfileExists {
				helperCreateEmptyFile(t, envfileName)
			}
			writer := envfile.NewEnvfile(envfileName, tc.whenExample, tc.whenOverwrite)

			// when
			if err := writer.Write(d.Envvars); err != nil {
				if tc.thenErrorSubMessage == "" {
					t.Fatalf("Writer.Write: %v", err)
				}
				// then
				if !strings.Contains(err.Error(), tc.thenErrorSubMessage) {
					t.Errorf("want %q to be in error %q", tc.thenErrorSubMessage, err.Error())
				}
				return
			}

			//then
			want := helperReadFile(t, tc.thenGoldenFile)
			got := helperReadFile(t, envfileName)
			if want != got {
				t.Errorf("want %q, got %q", want, got)
			}
		})
	}
}

func TestEnvfile_toRemoveFileIfItExists(t *testing.T) {
	// given
	name := t.TempDir() + "/.env"
	helperCreateEmptyFile(t, name)

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

// helperReadFile reads a file and returns it as string. It also removes trailing EOL.
func helperReadFile(t *testing.T, name string) string {
	f, err := os.ReadFile(name)
	if err != nil {
		t.Fatal(err.Error())
	}
	return strings.TrimSuffix(string(f), "\n")
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
