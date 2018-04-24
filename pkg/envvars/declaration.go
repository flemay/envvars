package envvars

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
	Example  string
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
