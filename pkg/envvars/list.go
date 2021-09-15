package envvars

// List returns a list of Envvar matching the tags or all if no tags is provided.
// Returns error if any tag is not matching an Envvar.
func List(reader DeclarationReader, tags ...string) (EnvvarCollection, error) {
	d, err := reader.Read()
	if err != nil {
		return nil, err
	}
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
