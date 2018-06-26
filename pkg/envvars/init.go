package envvars

// Init creates a Declaration File to serve as an example.
// Return an error if the Declaration File already exists.
func Init(writer DeclarationWriter) error {
	d := &Declaration{
		Envvars: EnvvarCollection{
			&Envvar{
				Name:    "ECHO",
				Desc:    "Description of ECHO",
				Example: "Hello World",
			},
		},
	}
	return writer.Write(d, false)
}
