package envvars

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

// Definition describes the environment variables.
// envvars.toml follows this structure.
type Definition struct {
	Envvars []*Envvar
}

// Envvar contains the information of a single environment variable
type Envvar struct {
	Name string
	Desc string
}

// NewDefinition reads a definition file and creates the environment variables Definition out of it.
func NewDefinition(filepath string) (*Definition, error) {
	var definition Definition
	if _, err := toml.DecodeFile(filepath, &definition); err != nil {
		return nil, fmt.Errorf("error occurred when opening the file %s: %s", filepath, err.Error())
	}
	return &definition, nil
}
