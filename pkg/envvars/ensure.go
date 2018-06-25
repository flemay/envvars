package envvars

import (
	"github.com/flemay/envvars/pkg/errorappender"
	"os"
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
	errorAppender := errorappender.NewErrorAppender("; ")
	value, found := os.LookupEnv(ev.Name)
	if found == false {
		errorAppender.AppendString("is not defined")
	} else if value == "" {
		errorAppender.AppendString("is empty")
	}

	return errorAppender.Wrap("environment variable " + ev.Name + " ")
}
