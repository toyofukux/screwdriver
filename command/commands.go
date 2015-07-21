package command

import "github.com/mitchellh/cli"

// Commands is the mapping of screwdriver sub commands.
var (
	// Commands hold Command implementation of screwdriver
	Commands map[string]cli.CommandFactory
)

func init() {
	Commands = map[string]cli.CommandFactory{
		"task": func() (cli.Command, error) {
			return &TaskCommand{}, nil
		},
		"service": func() (cli.Command, error) {
			return &ServiceCommand{}, nil
		},
	}
}
