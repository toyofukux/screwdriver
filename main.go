package main

import (
	"os"

	"github.com/mitchellh/cli"
	"github.com/takasing/screwdriver/command"
	"github.com/takasing/screwdriver/utils"
)

func main() {
	os.Exit(run())
}

func run() int {
	args := os.Args[1:]

	cli := &cli.CLI{
		Args:       args,
		Commands:   command.Commands,
		HelpFunc:   cli.BasicHelpFunc("screwdriver"),
		HelpWriter: os.Stdout,
	}

	exitCode, err := cli.Run()
	if err != nil {
		utils.ErrorOutput(err.Error())
		return 1
	}

	return exitCode
}
