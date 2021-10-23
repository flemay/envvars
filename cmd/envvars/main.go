package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"text/template"

	"github.com/flemay/envvars/internal/envfile"
	"github.com/flemay/envvars/internal/envvars"
	"github.com/flemay/envvars/internal/yml"
)

//go:embed "version.json"
var versionJSONFileEmbed []byte

func main() {
	const appName = "envvars"

	log.SetFlags(0)
	log.SetPrefix(appName + ": ")
	if err := run(appName, versionJSONFileEmbed); err != nil {
		log.Fatalf("%s\n", err)
	}
}

func run(appName string, versionJSONFile []byte) error {
	cmds := commands{
		ensureCmd(),
		envfileCmd(),
		initCmd(),
		listCmd(),
		validateCmd(),
		versionCmd(versionJSONFile),
	}

	usage, err := defaultUsage(appName, cmds)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s\n", usage)
	}

	var helpFlag bool
	flag.BoolVar(&helpFlag, "help", false, "Help")
	flag.BoolVar(&helpFlag, "h", false, "Help (shorthand)")

	// My understanding so far is if there is an error on parsing,
	// the error will be displayed followed by the usage.
	flag.Parse()

	if helpFlag {
		flag.Usage()
		// same behavior as COMMAND --help, so no error
		return nil
	}

	if len(os.Args) < 2 {
		// Maybe the usage should be printed so it is consistant with flag.Usage()
		// and flagSet.ExitOnError
		return fmt.Errorf("a command is required. See \"%s --help\".", appName)
	}

	cmd, exists := cmds.Get(os.Args[1])
	if !exists {
		return fmt.Errorf("%q is not an %s command. See \"%s --help\"", os.Args[1], appName, appName)
	}

	if cmd.Run == nil {
		return fmt.Errorf("the command %q does not implement the methor Run.", cmd.Name)
	}
	return cmd.Run(os.Args[2:])
}

func defaultUsage(appName string, cmds commands) (string, error) {
	data := struct {
		AppName string
		Cmds    commands
	}{
		appName,
		cmds,
	}
	usageTpl := `Usage:
    {{.AppName}} [flags]
    {{.AppName}} [command]

Flags:
    {{printf "%-15s" "-h"}} Output usage information (shorthand)
    {{printf "%-15s" "-help"}} Output usage information

Commands:
{{- range .Cmds}}
    {{printf "%-15s" .Name}} {{.Desc}}
{{- end}}

Use "{{.AppName}} [command] --help" for more information about a command.

Examples:
    Create a file to get started
        $ {{.AppName}} init
    Validate the file while developing
        $ {{.AppName}} validate
    Make sure the current environment variables comply to the Declaration file
        $ {{.AppName}} ensure`

	tmpl, err := template.New("usage").Parse(usageTpl)
	if err != nil {
		return "", err
	}

	var tmplResult bytes.Buffer
	if err := tmpl.Execute(&tmplResult, data); err != nil {
		return "", err
	}
	return tmplResult.String(), nil
}

type command struct {
	Name string
	Desc string
	Run  func(args []string) error
}

type commands []command

func (cmds commands) Get(name string) (command, bool) {
	for _, cmd := range cmds {
		if cmd.Name == name {
			return cmd, true
		}
	}
	return command{}, false
}

func defineFileFlag(fs *flag.FlagSet) *string {
	var file string
	fs.StringVar(&file, "file", "envvars.yml", "declaration file")
	fs.StringVar(&file, "f", "envvars.yml", "declaration file (shorthand)")
	return &file
}

func defineTagsFlag(fs *flag.FlagSet) *string {
	var tags string
	fs.StringVar(&tags, "tags", "", "comma-separeted list of tags")
	fs.StringVar(&tags, "t", "", "comma-separeted list of tags (shorthand)")
	return &tags
}

func commaSeparateTags(tags string) []string {
	var tagsSeparated []string
	if tags != "" {
		tagsSeparated = strings.Split(tags, ",")
	}
	return tagsSeparated
}

func initCmd() command {
	cmd := command{
		Name: "init",
		Desc: "Create a new declaration file to get started",
	}
	cmd.Run = func(args []string) error {
		fs := flag.NewFlagSet(cmd.Name, flag.ExitOnError)
		fileFlag := defineFileFlag(fs)
		fs.Parse(args)
		reader := yml.NewDeclarationYML(*fileFlag)
		return envvars.Init(reader)
	}
	return cmd
}

func validateCmd() command {
	cmd := command{
		Name: "validate",
		Desc: "Check if the declaration file contains any error",
	}
	cmd.Run = func(args []string) error {
		fs := flag.NewFlagSet(cmd.Name, flag.ExitOnError)
		fileFlag := defineFileFlag(fs)
		fs.Parse(args)
		reader := yml.NewDeclarationYML(*fileFlag)
		return envvars.Validate(reader)
	}
	return cmd
}

func ensureCmd() command {
	cmd := command{
		Name: "ensure",
		Desc: "Verify that the environment variables comply to their declaration. It also validates the declaration file",
	}
	cmd.Run = func(args []string) error {
		fs := flag.NewFlagSet(cmd.Name, flag.ExitOnError)
		fileFlag := defineFileFlag(fs)
		tagsFlag := defineTagsFlag(fs)
		fs.Parse(args)
		reader := yml.NewDeclarationYML(*fileFlag)
		return envvars.Ensure(reader, commaSeparateTags(*tagsFlag)...)
	}
	return cmd
}

func listCmd() command {
	cmd := command{
		Name: "list",
		Desc: "Display the declaration of each environment variable",
	}
	cmd.Run = func(args []string) error {
		fs := flag.NewFlagSet(cmd.Name, flag.ExitOnError)
		fileFlag := defineFileFlag(fs)
		tagsFlag := defineTagsFlag(fs)
		fs.Parse(args)
		reader := yml.NewDeclarationYML(*fileFlag)
		collection, err := envvars.List(reader, commaSeparateTags(*tagsFlag)...)
		if err != nil {
			return err
		}
		fmt.Println("name,desc")
		for _, ev := range collection {
			line := []string{ev.Name, ev.Desc}
			fmt.Println(strings.Join(line, ","))
		}
		return nil
	}
	return cmd
}

func envfileCmd() command {
	cmd := command{
		Name: "envfile",
		Desc: "Generate (or remove) an env file based on the declaration file",
	}
	cmd.Run = func(args []string) error {
		fs := flag.NewFlagSet(cmd.Name, flag.ExitOnError)
		fileFlag := defineFileFlag(fs)
		tagsFlag := defineTagsFlag(fs)
		envFileFlag := fs.String("env-file", ".env", "envfile to be generated")
		exampleFlag := fs.Bool("example", false, "include example values to the generated envfile")
		overwriteFlag := fs.Bool("overwrite", false, "overwrite the envfile if it exists")
		removeFlag := fs.Bool("rm", false, "remove the envfile")
		fs.Parse(args)

		if *removeFlag {
			return envfile.Remove(*envFileFlag)
		}
		reader := yml.NewDeclarationYML(*fileFlag)
		writer := envfile.NewEnvfile(*envFileFlag, *exampleFlag, *overwriteFlag)
		return envvars.Envfile(reader, writer, commaSeparateTags(*tagsFlag)...)

	}
	return cmd

}

func versionCmd(versionJSONFile []byte) command {
	cmd := command{
		Name: "version",
		Desc: "Output version information",
	}
	cmd.Run = func(args []string) error {
		fs := flag.NewFlagSet(cmd.Name, flag.ExitOnError)
		fs.Parse(args)

		var versionJSON struct {
			Version   string
			BuildDate string
			GitCommit string
		}

		if err := json.Unmarshal(versionJSONFile, &versionJSON); err != nil {
			return err
		}
		log.Printf(`
Version:      %s
Built:        %s
Git commit:   %s
Go version:   %s
OS/Arch:      %s/%s`,
			versionJSON.Version,
			versionJSON.BuildDate,
			versionJSON.GitCommit,
			runtime.Version(),
			runtime.GOOS,
			runtime.GOARCH)
		return nil
	}
	return cmd
}
