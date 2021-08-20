package yml

import (
	"fmt"
	"os"

	"github.com/flemay/envvars/pkg/envvars"
	"gopkg.in/yaml.v2"
)

// DeclarationYML handles read/write of a Declaration YML file
type DeclarationYML struct {
	filename string
}

// NewDeclarationYML returns a DeclarationYML object
func NewDeclarationYML(filename string) *DeclarationYML {
	return &DeclarationYML{filename}
}

// Read returns a Declaration from a yml file
func (declarationYML *DeclarationYML) Read() (*envvars.Declaration, error) {
	data, err := os.ReadFile(declarationYML.filename)
	if err != nil {
		return nil, fmt.Errorf("error occurred when reading the file %s: %s", declarationYML.filename, err.Error())
	}

	var d envvars.Declaration
	if err := yaml.Unmarshal(data, &d); err != nil {
		return nil, fmt.Errorf("error occurred when parsing the file %s: %s", declarationYML.filename, err.Error())
	}
	return &d, nil
}

// Write writes a Declaration object to a yml file.
// Returns error if the file already exists unless overwrite is true
func (declarationYML *DeclarationYML) Write(d *envvars.Declaration, overwrite bool) error {
	data, err := yaml.Marshal(d)
	if err != nil {
		return err
	}
	flag := os.O_CREATE | os.O_WRONLY
	if !overwrite {
		flag |= os.O_EXCL
	}
	file, err := os.OpenFile(declarationYML.filename, flag, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(data)
	return err
}
