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

	// envs := utils.LoadScrewEnvs()
	//
	// // TODO: enable to specify
	// data, err := ioutil.ReadFile("task.yml")
	// if err != nil {
	// 	panic(err)
	// }
	//
	// b, err := task.ExpandTemplate(data, envs)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(b))
	//
	// _, err = task.BindYml(b)
	// if err != nil {
	// 	panic(err)
	// }

	exitCode, err := cli.Run()
	if err != nil {
		fmt.Println(err)
		return 1
	}

	return exitCode
}
