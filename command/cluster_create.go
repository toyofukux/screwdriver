package command

import (
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/takasing/screwdriver/aws"
	"github.com/takasing/screwdriver/utils"
)

// ClusterCreateCommand is a Command implementation used to
// create ECS cluster.
type ClusterCreateCommand struct{}

// Run is Command implementation method for ClusterCreateCommand.
func (c *ClusterCreateCommand) Run(args []string) int {
	var name string
	flags := flag.NewFlagSet("cluster", flag.ContinueOnError)
	flags.StringVar(&name, "name", "", "cluster name")
	flags.Usage = func() { c.Help() }

	if err := flags.Parse(args); err != nil {
		utils.ErrorOutputf("Error parsing CLI flags: %s", err)
		fmt.Println(c.Help())
		return 1
	}

	if name == "" {
		utils.ErrorOutput("Error: -name option required.")
		return 1
	}

	input := &ecs.CreateClusterInput{
		ClusterName: &name,
	}

	client := aws.GetClient()
	resp, err := client.ClusterCreate(input)
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			utils.ErrorOutput(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				utils.ErrorOutput(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			utils.ErrorOutput(err.Error())
		}
		return 1
	}
	fmt.Println(awsutil.StringValue(resp))
	return 0
}

// Help show Command implementation method for ClusterCreateCommand.
func (c *ClusterCreateCommand) Help() string {
	return `
Usage: screwdriver cluster create [options]
	-name       The name of the ECS cluster.
`
}

// Synopsis Command implementation method for ClusterCreateCommand.
func (c *ClusterCreateCommand) Synopsis() string {
	return "create ECS cluster"
}
