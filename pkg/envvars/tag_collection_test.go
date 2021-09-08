package envvars_test

import (
	"testing"

	"github.com/flemay/envvars/pkg/envvars"
)

func TestTagCollectionGet(t *testing.T) {
	tags := envvars.TagCollection{
		&envvars.Tag{Name: "tag1"},
		&envvars.Tag{Name: "tag2"},
	}
	testCases := map[string]struct {
		givenTags   envvars.TagCollection
		whenTagName string
		thenTag     *envvars.Tag
	}{
		"return tag if match":    {tags, "tag2", tags[1]},
		"return nil if no match": {tags, "NOT_DEFINED", nil},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// when
			got := tc.givenTags.Get(tc.whenTagName)
			// then
			if tc.thenTag != nil {
				if got == nil {
					t.Errorf("want tag %q, got nil", tc.thenTag.Name)
					return
				}
				if got.Name != tc.thenTag.Name {
					t.Errorf("want tag with name %q, got %q", tc.thenTag.Name, got.Name)
				}
				return
			}
			if got != nil {
				t.Errorf("want nil, got tag with name %q", got.Name)
			}
		})
	}
}

func TestTagCollectionGetAll(t *testing.T) {
	tags := envvars.TagCollection{
		&envvars.Tag{Name: "tag1"},
		&envvars.Tag{Name: "tag2"},
		&envvars.Tag{Name: "tag1"},
	}
	testCases := map[string]struct {
		givenTags   envvars.TagCollection
		whenTagName string
		thenTags    envvars.TagCollection
	}{
		"return matching envvars":             {tags, "tag1", envvars.TagCollection{tags[0], tags[2]}},
		"return empty collection if no match": {tags, "NOT_DEFINED", envvars.TagCollection{}},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// when
			got := tc.givenTags.GetAll(tc.whenTagName)
			// then
			if len(got) != len(tc.thenTags) {
				t.Errorf("want %d, got %d", len(tc.thenTags), len(got))
				return
			}
			for i, tag := range got {
				if tag.Name != tc.thenTags[i].Name {
					t.Errorf("want %q, got %q", tc.thenTags[i].Name, tag.Name)
				}
			}
		})
	}
}
