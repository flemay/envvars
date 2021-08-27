package envvars

import "strings"

// DeclarationReader is a reader interface to get a Declaration object
type DeclarationReader interface {
	Read() (*Declaration, error)
}

// DeclarationWriter is a writer interface to write a Declaration object
type DeclarationWriter interface {
	Write(d *Declaration, overwrite bool) error
}

// Declaration describes the environment variables.
type Declaration struct {
	Tags    TagCollection `yaml:",omitempty"`
	Envvars EnvvarCollection
}

// Equal returns true if the current Declaration has the same values as the one
// passed in argument
func (d Declaration) Equal(declaration Declaration) bool {
	if len(d.Tags) != len(declaration.Tags) {
		return false
	}
	for i, t := range d.Tags {
		if t.Name != declaration.Tags[i].Name {
			return false
		}
		if t.Desc != declaration.Tags[i].Desc {
			return false
		}
	}

	if len(d.Envvars) != len(declaration.Envvars) {
		return false
	}
	for i, ev := range d.Envvars {
		if ev.Name != declaration.Envvars[i].Name {
			return false
		}
		if ev.Desc != declaration.Envvars[i].Desc {
			return false
		}
		if ev.Optional != declaration.Envvars[i].Optional {
			return false
		}
		if ev.Example != declaration.Envvars[i].Example {
			return false
		}
		if strings.Join(ev.Tags, "") != strings.Join(declaration.Envvars[i].Tags, "") {
			return false
		}
	}

	return true
}

// Envvar contains the information of a single environment variable.
type Envvar struct {
	Name     string
	Desc     string
	Tags     []string `yaml:",omitempty"`
	Optional bool     `yaml:",omitempty"`
	Example  string   `yaml:",omitempty"`
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
