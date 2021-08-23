package envvars

import (
	"fmt"
	"os"

	"github.com/flemay/envvars/pkg/errorappender"
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
	value, found := os.LookupEnv(ev.Name)

	var errMessage string
	if !found {
		errMessage = "is not defined"
	} else if value == "" {
		errMessage = "is empty"
	}

	if errMessage != "" {
		if ev.Desc != "" {
			errMessage = fmt.Sprintf("%s. Variable description: %s", errMessage, ev.Desc)
		}
		return fmt.Errorf("environment variable %s %s", ev.Name, errMessage)
	}
	return nil
}
