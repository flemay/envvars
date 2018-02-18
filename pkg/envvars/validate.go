package envvars

import (
	"errors"
	"fmt"
	"github.com/flemay/envvars/pkg/errorappender"
)

// Validate ensures the Declaration is without any error.
func Validate(d *Declaration) error {
	return validateDeclarationAndTagNameList(d)
}

func validateDeclaration(d *Declaration) error {
	errorAppender := errorappender.NewErrorAppender("\n")
	for i, tag := range d.Tags {
		tagErrorAppender := errorappender.NewErrorAppender("; ")
		tagErrorAppender.AppendError(validateTag(tag))
		tagErrorAppender.AppendError(validateTagNameUniqueness(tag.Name, d.Tags))
		tagErrorAppender.AppendError(validateTagUsage(tag.Name, d.Envvars))
		errorAppender.AppendError(tagErrorAppender.Wrap(fmt.Sprintf("tag '%s' (#%d): ", tag.Name, i+1)))
	}

	if len(d.Envvars) == 0 {
		errorAppender.AppendString("declaration must at least have 1 envvars")
	}

	for i, ev := range d.Envvars {
		evErrorAppender := errorappender.NewErrorAppender("; ")
		evErrorAppender.AppendError(validateEnvvar(ev, d.Tags))
		evErrorAppender.AppendError(validateEnvvarNameUniqueness(ev.Name, d.Envvars))
		errorAppender.AppendError(evErrorAppender.Wrap(fmt.Sprintf("envvar '%s' (#%d): ", ev.Name, i+1)))
	}

	return errorAppender.Error()
}

func validateDeclarationAndTagNameList(d *Declaration, tagNames ...string) error {
	if d == nil {
		return errors.New("declaration is nil")
	}
	errorAppender := errorappender.NewErrorAppender("\n")
	errorAppender.AppendError(validateDeclaration(d))
	errorAppender.AppendError(validateTagNameList(tagNames, d.Tags))
	return errorAppender.Error()
}

func validateEnvvar(ev *Envvar, tags TagCollection) error {
	errorAppender := errorappender.NewErrorAppender("; ")
	if ev.Name == "" {
		errorAppender.AppendString("name cannot be blank")
	}
	if ev.Desc == "" {
		errorAppender.AppendString("desc cannot be blank")
	}
	errorAppender.AppendError(validateTagNameList(ev.Tags, tags))

	return errorAppender.Error()
}

func validateEnvvarNameUniqueness(name string, c EnvvarCollection) error {
	if len(c.GetAll(name)) > 1 {
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

func validateTagNameUniqueness(name string, tags TagCollection) error {
	if len(tags.GetAll(name)) > 1 {
		return errors.New("name must be unique")
	}
	return nil
}

func validateTagUsage(name string, c EnvvarCollection) error {
	if len(c.WithTag(name)) == 0 {
		return errors.New("is not used")
	}
	return nil
}

func validateTagNameList(names []string, tags TagCollection) error {
	errorAppender := errorappender.NewErrorAppender("; ")
	counts := make(map[string]int)
	for _, name := range names {
		if name == "" {
			errorAppender.AppendString("tag '' is empty")
			continue
		}
		if tags.Get(name) == nil {
			errorAppender.AppendError(fmt.Errorf("tag '%v' is not declared", name))
		}
		counts[name] = counts[name] + 1
	}
	for name, counts := range counts {
		if counts > 1 {
			errorAppender.AppendError(fmt.Errorf("tag '%v' is duplicated", name))
		}
	}
	return errorAppender.Error()
}
