package envvars

// EnvfileWriter defines an interface that is used to write into an envfile.
type EnvfileWriter interface {
	Write(c EnvvarCollection) error
}

// Envfile generates an env file that can be overwritten.
// It returns an error if the file already exists unless overwrite is true
func Envfile(d *Declaration, writer EnvfileWriter, tags ...string) error {
	c, err := List(d, tags...)
	if err != nil {
		return err
	}
	return writer.Write(c)
}
