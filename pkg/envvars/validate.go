package envvars

import (
	"errors"
	"fmt"
	"github.com/flemay/envvars/pkg/errorappender"
)

// Validate ensures the Definition is without any error.
// Consumer should always validate before doing any action with the Definition
func Validate(definition *Definition) error {
	errorAppender := errorappender.NewErrorAppender("\n")
	for i, ev := range definition.Envvars {
		evErrorAppender := errorappender.NewErrorAppender("; ")
		evErrorAppender.AppendError(validateEnvvar(ev))
		evErrorAppender.AppendError(validateNameUniqueness(ev.Name, definition.Envvars))
		errorAppender.AppendError(evErrorAppender.Wrap(fmt.Sprintf("Envvar #%d: ", i+1)))
	}

	return errorAppender.Error()
}

func validateEnvvar(ev *Envvar) error {
	errorAppender := errorappender.NewErrorAppender("; ")
	if ev.Name == "" {
		errorAppender.AppendString("name cannot be blank")
	}
	if ev.Desc == "" {
		errorAppender.AppendString("desc cannot be blank")
	}
	return errorAppender.Error()
}

func validateNameUniqueness(name string, evs []*Envvar) error {
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
