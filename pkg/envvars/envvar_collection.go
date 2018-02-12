package envvars

// EnvvarCollection is an Envvar collection with helper functions.
type EnvvarCollection []*Envvar

// WithTag returns a collection of Envvar that are tagged with tag name.
func (c EnvvarCollection) WithTag(name string) EnvvarCollection {
	taggedEvs := make(EnvvarCollection, 0, len(c))
	for _, ev := range c {
		if ev.HasTag(name) {
			taggedEvs = append(taggedEvs, ev)
		}
	}
	return taggedEvs
}

// Get returns the Envvar that has the name.
func (c EnvvarCollection) Get(name string) *Envvar {
	for _, ev := range c {
		if ev.Name == name {
			return ev
		}
	}
	return nil
}

// GetAll returns all Envvar that has name.
func (c EnvvarCollection) GetAll(name string) EnvvarCollection {
	ec := make(EnvvarCollection, 0)
	for _, ev := range c {
		if ev.Name == name {
			ec = append(ec, ev)
		}
	}
	return ec
}
