package envvars

import (
	"errors"
	"fmt"
	"github.com/flemay/envvars/pkg/errorappender"
)

// Validate ensures the Definition is without any error.
// Consumer should always validate before doing any action with the Definition.
func Validate(definition *Definition) error {
	errorAppender := errorappender.NewErrorAppender("\n")
	for i, tag := range definition.Tags {
		tagErrorAppender := errorappender.NewErrorAppender("; ")
		tagErrorAppender.AppendError(validateTag(tag))
		tagErrorAppender.AppendError(validateTagNameUniqueness(tag.Name, definition.Tags))
		tagErrorAppender.AppendError(validateTagUsage(tag.Name, definition.Envvars))
		errorAppender.AppendError(tagErrorAppender.Wrap(fmt.Sprintf("Tag '%s' (#%d): ", tag.Name, i+1)))
	}
	for i, ev := range definition.Envvars {
		evErrorAppender := errorappender.NewErrorAppender("; ")
		evErrorAppender.AppendError(validateEnvvar(ev, definition.Tags))
		evErrorAppender.AppendError(validateEnvvarNameUniqueness(ev.Name, definition.Envvars))
		errorAppender.AppendError(evErrorAppender.Wrap(fmt.Sprintf("Envvar '%s' (#%d): ", ev.Name, i+1)))
	}

	return errorAppender.Error()
}

func validateEnvvar(ev *Envvar, tags []*Tag) error {
	errorAppender := errorappender.NewErrorAppender("; ")
	if ev.Name == "" {
		errorAppender.AppendString("name cannot be blank")
	}
	if ev.Desc == "" {
		errorAppender.AppendString("desc cannot be blank")
	}
	for _, tagName := range ev.Tags {
		tagNotDefined := true
		for _, tag := range tags {
			if tag.Name == tagName {
				tagNotDefined = false
				break
			}
		}
		if tagNotDefined {
			errorAppender.AppendString("tag '" + tagName + "' is not defined")
		}
	}
	return errorAppender.Error()
}

func validateEnvvarNameUniqueness(name string, evs []*Envvar) error {
	if name == "" {
		return nil
	}
	duplicateNb := 0
	for _, ev := range evs {
		if name == ev.Name {
			duplicateNb++
		}
	}
	if duplicateNb > 1 {
		return errors.New("name must be unique")
	}
	return nil
}

func validateTag(t *Tag) error {
	errorAppender := errorappender.NewErrorAppender("; ")
	if t.Name == "" {
		errorAppender.AppendString("name cannot be blank")
	}
	if t.Desc == "" {
		errorAppender.AppendString("desc cannot be blank")
	}
	return errorAppender.Error()
}

func validateTagNameUniqueness(name string, tags []*Tag) error {
	if name == "" {
		return nil
	}
	duplicateNb := 0
	for _, tag := range tags {
		if name == tag.Name {
			duplicateNb++
		}
	}
	if duplicateNb > 1 {
		return errors.New("name must be unique")
	}
	return nil
}

func validateTagUsage(name string, evs []*Envvar) error {
	if name == "" {
		return nil
	}
	notUsed := true
	for i := 0; i < len(evs) && notUsed; i++ {
		for _, tagName := range evs[i].Tags {
			if tagName == name {
				notUsed = false
				break
			}
		}
	}
	if notUsed {
		return errors.New("defined but not used")
	}
	return nil
}
