package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type rootFlags struct {
	file *string
}

func main() {
	commands := map[string]commander{
		"init": newInitCmd(),
	}
	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	cmd, exists := commands[os.Args[1]]
	if !exists {
		flag.Usage()
		os.Exit(1)
	}

	log.SetFlags(0)
	log.SetPrefix("envvars: ")

	err := cmd.Run(os.Args[2:])
	if err != nil {
		log.Fatalf("error: %s\n", err)
	}
}

type commander interface {
	Name() string
	Usage() string
	Run(args []string) error
}

type command struct {
	flagSet *flag.FlagSet
	usage   string
	run     func() error
}

func (c command) Name() string {
	return c.flagSet.Name()
}

func (c command) Usage() string {
	return c.usage
}

func (c command) Run(args []string) error {
	if err := c.flagSet.Parse(args); err != nil {
		return err
	}
	if c.run == nil {
		return fmt.Errorf("command %q does not implement the method 'run'", c.Name())
	}
	return c.run()
}

func newInitCmd() commander {
	cmd := struct{ command }{}
	cmd.flagSet = flag.NewFlagSet("init", flag.ExitOnError)
	cmd.usage = "Creates a new Declaration File to get started"
	var flagFile string
	cmd.flagSet.StringVar(&flagFile, "file", "envvars.yml", "declaration file")
	cmd.flagSet.StringVar(&flagFile, "f", "envvars.yml", "declaration file (shorthand)")
	cmd.run = func() error {
		fmt.Printf("running init %q\n", flagFile)
		return nil
	}
	return cmd
}
