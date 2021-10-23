package envvars

import (
	"errors"
	"fmt"
	"os"

	"github.com/flemay/envvars/internal/errorappender"
)

// Ensure verifies that the environment variables comply to their Declaration. Tags can be passed along to only ensure environment variables with the tags
func Ensure(reader DeclarationReader, tags ...string) error {
	c, err := List(reader, tags...)
	if err != nil {
		return err
	}
	errorAppender := errorappender.NewErrorAppender("\n")
	for _, ev := range c {
		errorAppender.AppendError(ensureEnvvar(ev))
	}
	return errorAppender.Error()
}

func ensureEnvvar(ev *Envvar) error {
	if ev.Optional {
		return nil
	}

	newError := func(envIsMessage string) error {
		message := fmt.Sprintf("environment variable %s %s", ev.Name, envIsMessage)
		if ev.Desc != "" {
			message = fmt.Sprintf("%s. Variable description: %s", message, ev.Desc)
		}
		return errors.New(message)
	}

	value, found := os.LookupEnv(ev.Name)
	if !found {
		return newError("is not defined")
	}

	if value == "" {
		return newError("is empty")
	}

	return nil
}
