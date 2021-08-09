package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"text/template"

	"github.com/flemay/envvars/pkg/envvars"
	"github.com/flemay/envvars/pkg/yml"
)

func main() {
	const appName = "envvars"

	log.SetFlags(0)
	log.SetPrefix(appName + ": ")
	if err := run(appName); err != nil {
		log.Fatalf("%s\n", err)
	}
}

func run(appName string) error {
	cmds := commands{
		ensureCmd(),
		initCmd(),
		listCmd(),
		validateCmd(),
		versionCmd("1.1", "some date", "2343243"),
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
    {{.AppName}} COMMAND [OPTIONS]
    {{.AppName}} --help

Commands:
{{- range .Cmds}}
    {{.Name}}    {{.Desc}}
{{- end}}

Run "{{.AppName}} COMMAND --help" for more information on a command.

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

type commandFlag int

const (
	commandFlagFile commandFlag = iota
	commandFlagTags
)

type commandFlagSet struct {
	flagSet       *flag.FlagSet
	file          string
	tags          string
	tagsSeparated []string
}

func (cfs *commandFlagSet) Parse(args []string) {
	// Since ExitOnError is used, no need to look at a returned error
	cfs.flagSet.Parse(args)
	if cfs.tags != "" {
		cfs.tagsSeparated = strings.Split(cfs.tags, ",")
	}
}

func newCommandFlagSet(name string, cmdFlags ...commandFlag) *commandFlagSet {
	cfs := commandFlagSet{}
	cfs.flagSet = flag.NewFlagSet(name, flag.ExitOnError)
	for _, cf := range cmdFlags {
		switch cf {
		case commandFlagFile:
			cfs.flagSet.StringVar(&cfs.file, "file", "envvars.yml", "declaration file")
			cfs.flagSet.StringVar(&cfs.file, "f", "envvars.yml", "declaration file (shorthand)")
		case commandFlagTags:
			cfs.flagSet.StringVar(&cfs.tags, "tags", "", "comma-separeted list of tags")
			cfs.flagSet.StringVar(&cfs.tags, "t", "", "comma-separeted list of tags (shorthand)")
		}
	}
	return &cfs
}

func initCmd() command {
	cmd := command{
		Name: "init",
		Desc: "Create a new declaration file to get started",
	}
	cmd.Run = func(args []string) error {
		cfs := newCommandFlagSet(cmd.Name, commandFlagFile)
		cfs.Parse(args)
		reader := yml.NewDeclarationYML(cfs.file)
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
		cfs := newCommandFlagSet(cmd.Name, commandFlagFile)
		cfs.Parse(args)
		reader := yml.NewDeclarationYML(cfs.file)
		return envvars.Validate(reader)
	}
	return cmd
}

func ensureCmd() command {
	cmd := command{
		Name: "ensure",
		Desc: "Verify that the environment variables comply to their declaration",
	}
	cmd.Run = func(args []string) error {
		cfs := newCommandFlagSet(cmd.Name, commandFlagFile, commandFlagTags)
		cfs.Parse(args)
		reader := yml.NewDeclarationYML(cfs.file)
		return envvars.Ensure(reader, cfs.tagsSeparated...)
	}
	return cmd
}

func listCmd() command {
	cmd := command{
		Name: "list",
		Desc: "Display the declaration of each environment variable",
	}
	cmd.Run = func(args []string) error {
		cfs := newCommandFlagSet(cmd.Name, commandFlagFile, commandFlagTags)
		cfs.Parse(args)
		reader := yml.NewDeclarationYML(cfs.file)
		collection, err := envvars.List(reader, cfs.tagsSeparated...)
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

func versionCmd(version string, buildDate string, gitCommit string) command {
	cmd := command{
		Name: "version",
		Desc: "Show version information",
	}
	cmd.Run = func(args []string) error {
		cfs := newCommandFlagSet(cmd.Name)
		cfs.Parse(args)

		log.Printf(`
Version:      %s
Built:        %s
Git commit:   %s
Go version:   %s
OS/Arch:      %s/%s`, version, buildDate, gitCommit, runtime.Version(), runtime.GOOS, runtime.GOARCH)
		return nil
	}
	return cmd
}
