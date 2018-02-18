package envvars

import (
	"bufio"
	"fmt"
	"os"
)

// Envfile generates an env file that can be overwritten.
// It returns an error if the file already exists unless overwrite is true
func Envfile(d *Declaration, name string, overwrite bool, tags ...string) error {
	c, err := List(d, tags...)
	if err != nil {
		return err
	}
	fileinfo, err := os.Stat(name)
	if err != nil && os.IsNotExist(err) == false {
		return err
	}
	if err == nil && overwrite == false {
		return fmt.Errorf("error: %s already exist", name)
	}
	if err == nil && fileinfo.IsDir() {
		return fmt.Errorf("error: %s is a folder, not a file", name)
	}

	return writeEnvfile(c, name)
}

func writeEnvfile(c EnvvarCollection, name string) error {
	f, err := os.Create(name)
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
