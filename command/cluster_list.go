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

// ClusterListCommand is a Command implementation used to
// show the list of ECS cluster.
type ClusterListCommand struct{}

// Run is Command implementation method for ClusterListCommand.
func (c *ClusterListCommand) Run(args []string) int {
	var nextToken string
	var max int64
	flags := flag.NewFlagSet("cluster", flag.ContinueOnError)
	flags.StringVar(&nextToken, "next", "", "next token")
	flags.Int64Var(&max, "max", 10, "the number of cluster to show")
	flags.Usage = func() { c.Help() }

	if err := flags.Parse(args); err != nil {
		utils.ErrorOutputf("Error parsing CLI flags: %s", err)
		fmt.Println(c.Help())
		return 1
	}

	input := &ecs.ListClustersInput{
		NextToken:  &nextToken,
		MaxResults: &max,
	}

	client := aws.GetClient()
	resp, err := client.ClusterList(input)
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

// Help show Command implementation method for ClusterListCommand.
func (c *ClusterListCommand) Help() string {
	return `
Usage: screwdriver cluster list [options]
	-cluster       The name of the cluster to show the list.
	-max           The count of the cluster to show.
	-next          The token to specify where to start paginating.
`
}

// Synopsis Command implementation method for ClusterListCommand.
func (c *ClusterListCommand) Synopsis() string {
	return "Show ECS cluster list"
}
