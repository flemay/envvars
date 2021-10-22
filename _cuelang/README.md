# Cuelang

This is an experiment on using cuelang to validate and extend envvars. Could it even replace envvars altogether?

## Usage

```bash
# Run cuelang with Docker
$ docker run -it --rm -v $(PWD):/opt/app -w /opt/app golang bash

# Install cuelang
$$ go install cuelang.org/go/cmd/cue@latest

# Vet
$$ cue vet schema.cue envvars.yml

# Eval
$$ cue eval schema.cue envvars.yml

# Format
$$ cue fmt schema.cue

$$ exit
```

## Todo

- Another cue file to extend validation functionalities. For example, a project may need to validate the environment variable ENV to have choice values such as "dev", "qa", "prod". Envvars would have a base schema with a set of validations and then a project could extend the validation.
- Can `pkg/tool/os` be used to access and validate environment variable from cue?
- Can `.env` file be generated from cue cmd? See `$ cue cmd --help`
- Integrate cue with envvars to validate `envvars.yml`. See [Processing CUE in Go](https://cuelang.org/docs/integrations/go/#processing-cue-in-go)

## References

- https://cuelang.org/
- https://cuetorials.com/
- https://github.com/cue-lang/cue/discussions/715
- How to use `pkg/tool/http`? https://github.com/cue-lang/cue/discussions/542

