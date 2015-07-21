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

// ServiceListCommand is a Command implementation used to
// get ECS service list.
type ServiceListCommand struct{}

// Run is Command implementation method for ServiceListCommand.
func (c *ServiceListCommand) Run(args []string) int {
	var cluster, nextToken string
	var max int64
	flags := flag.NewFlagSet("service", flag.ContinueOnError)
	flags.StringVar(&cluster, "cluster", "", "cluster the service belong to.")
	flags.StringVar(&nextToken, "next", "", "next token")
	flags.Int64Var(&max, "max", 10, "the number of service to show")
	flags.Usage = func() { c.Help() }

	if err := flags.Parse(args); err != nil {
		utils.ErrorOutputf("Error parsing CLI flags: %s", err)
		fmt.Println(c.Help())
		return 1
	}
	input := &ecs.ListServicesInput{
		Cluster:    &cluster,
		MaxResults: &max,
		NextToken:  &nextToken,
	}
	client := aws.GetClient()
	resp, err := client.ServiceList(input)
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

// Help show Command implementation method for ServiceListCommand.
func (c *ServiceListCommand) Help() string {
	return `
Usage: screwdriver service list [options]
	-cluster       The name of the cluster to show the list.
	-max           The count of the service to show.
	-next          The token to specify where to start paginating.
`
}

// Synopsis Command implementation method for ServiceListCommand.
func (c *ServiceListCommand) Synopsis() string {
	return "Show ECS service list"
}
