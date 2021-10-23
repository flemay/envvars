package envvars

// TagCollection is a Tag collection with helper functions.
type TagCollection []*Tag

// Get returns the Tag that has name.
func (c TagCollection) Get(name string) *Tag {
	for _, t := range c {
		if t.Name == name {
			return t
		}
	}
	return nil
}

// GetAll returns all Tag that has name.
func (c TagCollection) GetAll(name string) TagCollection {
	tags := make(TagCollection, 0)
	for _, t := range c {
		if t.Name == name {
			tags = append(tags, t)
		}
	}
	return tags
}
