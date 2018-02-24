package main

import (
	"github.com/flemay/envvars/cmd"
)

var (
	version    = "devel"
	buildDate  string
	commitHash string
)

func main() {
	cmd.Execute(version, commitHash, buildDate)
}
