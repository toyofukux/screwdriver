package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/cli"
	"github.com/takasing/screwdriver/command"
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
		fmt.Println(err)
		return 1
	}

	return exitCode
}
