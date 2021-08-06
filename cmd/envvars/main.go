package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
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
		initCmd(),
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
	flag.Parse()

	if helpFlag {
		flag.Usage()
		// same behavior as COMMAND --help, so no error
		return nil
	}

	if len(os.Args) < 2 {
		return fmt.Errorf("a command is required. See \"%s --help\".", appName)
	}

	cmd, exists := cmds.Get(os.Args[1])
	if !exists {
		return fmt.Errorf("%q is not an %s command. See \"%s --help\"", os.Args[1], appName, appName)
	}

	return cmd.Run(os.Args[2:])
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

Commands:
{{range .Cmds}}
  {{.Name}}    {{.Desc}}
{{end}}

Run "{{.AppName}} COMMAND --help" for more information on a command.`

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

func initCmd() command {
	cmd := command{
		Name: "init",
		Desc: "Create a new declaration file to get started",
	}
	cmd.Run = func(args []string) error {
		fs := flag.NewFlagSet(cmd.Name, flag.ExitOnError)
		var flagFile string
		fs.StringVar(&flagFile, "file", "envvars.yml", "declaration file")
		fs.StringVar(&flagFile, "f", "envvars.yml", "declaration file (shorthand)")

		// Since ExitOnError is used, no need to look at a returned error
		fs.Parse(args)

		reader := yml.NewDeclarationYML(flagFile)
		return envvars.Init(reader)
	}
	return cmd
}
