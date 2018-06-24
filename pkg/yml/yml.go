package yml

import (
	"fmt"
	"github.com/flemay/envvars/pkg/envvars"
	"gopkg.in/yaml.v2"
	"io/ioutil"
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
	data, err := ioutil.ReadFile(declarationYML.filename)
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
	return nil
}

// NewDeclaration returns a Declaration from a yml file
func NewDeclaration(filename string) (*envvars.Declaration, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error occurred when reading the file %s: %s", filename, err.Error())
	}

	var d envvars.Declaration
	if err := yaml.Unmarshal(data, &d); err != nil {
		return nil, fmt.Errorf("error occurred when parsing the file %s: %s", filename, err.Error())
	}
	return &d, nil
}
