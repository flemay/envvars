package envvars

import (
	"github.com/flemay/envvars/pkg/errorappender"
	"os"
)

// Ensure verifies that the environment variables comply to their Declaration. Tags can be passed along to only ensure environment variables with the tags
func Ensure(d *Declaration, tags ...string) error {
	c, err := List(d, tags...)
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
	errorAppender := errorappender.NewErrorAppender("; ")
	_, found := os.LookupEnv(ev.Name)
	if found == false {
		errorAppender.AppendString("must be set")
	}
	return errorAppender.Wrap("environment variable " + ev.Name + " ")
}
