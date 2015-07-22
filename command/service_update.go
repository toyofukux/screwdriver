package command

import (
	"errors"
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/takasing/screwdriver/aws"
	"github.com/takasing/screwdriver/utils"
)

// ServiceUpdateCommand is a Command implementation used to
// update ECS service
type ServiceUpdateCommand struct{}

// Run is Command implementation method for ServiceUpdateCommand.
func (c *ServiceUpdateCommand) Run(args []string) int {
	var cluster, name, task string
	var desire int64
	flags := flag.NewFlagSet("create", flag.ContinueOnError)
	flags.StringVar(&cluster, "cluster", "default", "cluster name you want service belong to")
	flags.Int64Var(&desire, "desire", 0, "desire count to keep the number of running task")
	flags.StringVar(&name, "name", "", "service name")
	flags.StringVar(&task, "task", "", "task definition you want to attach")
	flags.Usage = func() { c.Help() }

	if err := flags.Parse(args); err != nil {
		utils.ErrorOutputf("Error parsing CLI flags: %s", err)
		fmt.Println(c.Help())
		return 1
	}

	input, err := generateUpdateServiceInput(cluster, name, task, desire)
	if err != nil {
		utils.ErrorOutput(err.Error())
		return 1
	}

	client := aws.GetClient()
	resp, err := client.ServiceUpdate(input)

	// FIXME: utils
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

	fmt.Printf("Success update service %s.", name)
	fmt.Println(awsutil.StringValue(resp))
	return 0
}

func generateUpdateServiceInput(cluster string, name string, task string, desire int64) (*ecs.UpdateServiceInput, error) {
	if name == "" {
		return nil, errors.New("Error: Require -name option to specify service name.")
	}

	if task == "" {
		return nil, errors.New("Error: Require -task option to attach task definition.")
	}

	input := &ecs.UpdateServiceInput{
		Cluster:        &cluster,
		DesiredCount:   &desire,
		Service:        &name,
		TaskDefinition: &task,
	}

	return input, nil
}

// Help show Command implementation method for ServiceUpdateCommand.
func (c *ServiceUpdateCommand) Help() string {
	helpText := `
Usage: screwdriver service update [options]
`
	return helpText
}

// Synopsis Command implementation method for ServiceUpdateCommand.
func (c *ServiceUpdateCommand) Synopsis() string {
	return "update ECS service"
}
