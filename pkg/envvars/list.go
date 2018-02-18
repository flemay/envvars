package envvars

// List returns a list of Envvar matching the tags or all if no tags is provided.
// Returns error if any tag is not matching an Envvar.
func List(d *Declaration, tags ...string) (EnvvarCollection, error) {
	if err := validateDeclarationAndTagNameList(d, tags...); err != nil {
		return nil, err
	}

	if len(tags) == 0 {
		return d.Envvars, nil
	}

	taggedCollection := make(EnvvarCollection, 0, len(d.Envvars))
	for _, t := range tags {
		collection := d.Envvars.WithTag(t)
		taggedCollection = appendToEnvvarCollection(taggedCollection, collection...)
	}

	return taggedCollection, nil
}

func appendToEnvvarCollection(c EnvvarCollection, evs ...*Envvar) EnvvarCollection {
	toAppend := make(EnvvarCollection, 0, len(evs))
	for _, ev := range evs {
		if c.Get(ev.Name) == nil {
			toAppend = append(toAppend, ev)
		}
	}
	if len(toAppend) > 0 {
		return append(c, toAppend...)
	}
	return c
}
