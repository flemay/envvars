package envvars

import (
	"github.com/flemay/envvars/pkg/errorappender"
	"os"
)

// Ensure verifies that the environment variables comply to their Definition.
func Ensure(definition *Definition) error {
	errorAppender := errorappender.NewErrorAppender("\n")
	for _, ev := range definition.Envvars {
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
