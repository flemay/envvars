package envvars

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// Declaration describes the environment variables.
type Declaration struct {
	Tags    TagCollection
	Envvars EnvvarCollection
}

// Envvar contains the information of a single environment variable.
type Envvar struct {
	Name     string
	Desc     string
	Tags     []string
	Optional bool
}

func (ev *Envvar) HasTag(name string) bool {
	for _, t := range ev.Tags {
		if t == name {
			return true
		}
	}
	return false
}

// Tag allows targetting environnement variables for a specific purpose.
type Tag struct {
	Name string
	Desc string
}

// NewDeclaration reads a declaration file and returns a Declaration.
func NewDeclaration(filepath string) (*Declaration, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("error occurred when reading the file %s: %s", filepath, err.Error())
	}

	var d Declaration
	if err := yaml.Unmarshal(data, &d); err != nil {
		return nil, fmt.Errorf("error occurred when parsing the file %s: %s", filepath, err.Error())
	}
	return &d, nil
}
