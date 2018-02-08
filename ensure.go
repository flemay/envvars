package envvars

import (
	"os"
)

// Ensure verifies that the environment variables are comform to the Metadata.
func Ensure(metadata *Metadata) error {
	errorAppender := NewErrorAppender("\n")
	for _, ev := range metadata.Envvars {
		errorAppender.AppendError(ensureEnvvar(ev))
	}
	return errorAppender.Error()
}

func ensureEnvvar(ev *Envvar) error {
	errorAppender := NewErrorAppender("; ")
	_, found := os.LookupEnv(ev.Name)
	if found == false {
		errorAppender.AppendString("must be set")
	}
	return errorAppender.Wrap("environment variable " + ev.Name + " ")
}
