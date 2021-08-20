package envfile

import (
	"bufio"
	"fmt"
	"os"

	"github.com/flemay/envvars/pkg/envvars"
)

// Envfile define a struct which is responsible to generates an env file.
type Envfile struct {
	filename  string
	example   bool
	overwrite bool
}

// Write generates an env file that can be overwritten.
// It returns an error if the file already exists unless overwrite is true
func (e *Envfile) Write(c envvars.EnvvarCollection) error {
	fileinfo, err := os.Stat(e.filename)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	if err == nil && !e.overwrite {
		return fmt.Errorf("error: %s already exist", e.filename)
	}
	if err == nil && fileinfo.IsDir() {
		return fmt.Errorf("error: %s is a folder, not a file", e.filename)
	}

	f, err := os.Create(e.filename)
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	for _, ev := range c {
		line := ev.Name
		if e.example && ev.Example != "" {
			line += "=" + ev.Example
		}
		line += "\n"

		if _, err := w.WriteString(line); err != nil {
			return err
		}
	}
	return w.Flush()
}

// NewEnvfile returns an Envfile struct
func NewEnvfile(filename string, example bool, overwrite bool) *Envfile {
	return &Envfile{filename, example, overwrite}
}

// Remove removes the env file. Returns an error if the file does not exist or any permission issue.
func Remove(filename string) error {
	return os.Remove(filename)
}
