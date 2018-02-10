package envvars

import (
	"bufio"
	"fmt"
	"os"
)

// Dotenv generates a .env file (dotenvPath) that can be overwritten.
// It returns an error if the file already exists unless overwrite is true
func Dotenv(definition *Definition, dotenvPath string, overwrite bool) error {
	fileinfo, err := os.Stat(dotenvPath)
	if err != nil && os.IsNotExist(err) == false {
		return err
	}
	if err == nil && overwrite == false {
		return fmt.Errorf("error: %s already exist", dotenvPath)
	}
	if err == nil && fileinfo.IsDir() {
		return fmt.Errorf("error: %s is a folder, not a file", dotenvPath)
	}

	return writeDotenv(definition, dotenvPath)
}

func writeDotenv(definition *Definition, dotenvPath string) error {
	f, err := os.Create(dotenvPath)
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	for _, ev := range definition.Envvars {
		if _, err := w.WriteString(ev.Name + "\n"); err != nil {
			return err
		}
	}
	return w.Flush()
}
