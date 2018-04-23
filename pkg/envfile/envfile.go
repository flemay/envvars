package envfile

import (
	"bufio"
	"fmt"
	"github.com/flemay/envvars/pkg/envvars"
	"os"
)

// Envfile define a struct which is reponsible to generates an env file.
type Envfile struct {
	filename  string
	overwrite bool
}

// Write generates an env file that can be overwritten.
// It returns an error if the file already exists unless overwrite is true
func (e *Envfile) Write(c envvars.EnvvarCollection) error {
	fileinfo, err := os.Stat(e.filename)
	if err != nil && os.IsNotExist(err) == false {
		return err
	}
	if err == nil && e.overwrite == false {
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
		if _, err := w.WriteString(ev.Name + "\n"); err != nil {
			return err
		}
	}
	return w.Flush()
}

// NewEnvfile returns an Envfile struct
func NewEnvfile(name string, overwrite bool) *Envfile {
	return &Envfile{name, overwrite}
}
