package command

import (
	"github.com/mitchellh/cli"
	"github.com/takasing/screwdriver/aws"
)

// Commands is the mapping of screwdriver sub commands.
var (
	// Commands hold Command implementation of screw
	Commands map[string]cli.CommandFactory
	// AWSClient call AWS API
	awsClient *aws.Client
)

func init() {
	Commands = map[string]cli.CommandFactory{
		"task": func() (cli.Command, error) {
			return &TaskCommand{}, nil
		},
	}
	// FIXME never initialize in init func
	c, err := aws.NewClient()
	if err != nil {
		panic(err)
	}
	client, ok := c.(*aws.Client)
	if !ok {
		panic("cannot get aws client")
	}
	awsClient = client
}
