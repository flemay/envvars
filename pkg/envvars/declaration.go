package envvars

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
