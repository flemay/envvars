package yml

import (
	"fmt"
	"github.com/flemay/envvars/pkg/envvars"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

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
