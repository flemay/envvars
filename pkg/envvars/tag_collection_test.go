package envvars_test

import (
	"github.com/flemay/envvars/pkg/envvars"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTagCollection_Get_toReturnMatchingNameTag(t *testing.T) {
	// given
	c := envvars.TagCollection{
		&envvars.Tag{Name: "TAG_1"},
		&envvars.Tag{Name: "TAG_2"},
	}
	// when
	tag := c.Get("TAG_2")
	// then
	assert.NotNil(t, tag)
	assert.Equal(t, "TAG_2", tag.Name)
}

func TestTagCollection_Get_toReturnNilIfNoneMatchingName(t *testing.T) {
	// given
	c := envvars.TagCollection{
		&envvars.Tag{Name: "TAG_1"},
		&envvars.Tag{Name: "TAG_2"},
	}
	// when
	tag := c.Get("NOT_DEFINED")
	// then
	assert.Nil(t, tag)
}

func TestTagCollection_GetAll_toReturnTagCollectionMatchingName(t *testing.T) {
	// given
	c := envvars.TagCollection{
		&envvars.Tag{Name: "TAG_1"},
		&envvars.Tag{Name: "TAG_2"},
		&envvars.Tag{Name: "TAG_1"},
	}
	// when
	tags := c.GetAll("TAG_1")
	// then
	assert.Len(t, tags, 2)
	assert.Equal(t, "TAG_1", tags[0].Name)
	assert.Equal(t, "TAG_1", tags[1].Name)
}

func TestTagCollection_GetAll_toReturnEmtpyTagCollectionIfNoneMatchingName(t *testing.T) {
	// given
	c := envvars.TagCollection{
		&envvars.Tag{Name: "TAG_1"},
		&envvars.Tag{Name: "TAG_2"},
		&envvars.Tag{Name: "TAG_1"},
	}
	// when
	tags := c.GetAll("NOT_DEFINED")
	// then
	assert.Len(t, tags, 0)
}
