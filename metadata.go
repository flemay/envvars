package envvars

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

// Metadata describes the environment variables.
// envvars.toml follows this structure.
type Metadata struct {
	Envvars []Envvar
}

// Envvar contains the information of a single environment variable
type Envvar struct {
	Name        string
	Description string
}

// NewMetadata reads a file and creates the environment variables Metadata out of it.
func NewMetadata(filepath string) (*Metadata, error) {
	var metadata Metadata
	if _, err := toml.DecodeFile(filepath, &metadata); err != nil {
		return nil, fmt.Errorf("error occurred when opening the file %s: %s", filepath, err.Error())
	}
	return &metadata, nil
}
