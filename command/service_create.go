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

// ServiceCreateCommand is a Command implementation used to
// create ECS service
type ServiceCreateCommand struct{}

// Run is Command implementation method for ServiceCreateCommand
func (c *ServiceCreateCommand) Run(args []string) int {
	// FIXME support ELB
	// FIXME support idempotent
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

	input, err := generateCreateServiceInput(cluster, name, task, desire)
	if err != nil {
		utils.ErrorOutput(err.Error())
		return 1
	}

	client := aws.GetClient()
	resp, err := client.ServiceCreate(input)

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

	fmt.Printf("Success create service %s.", name)
	fmt.Println(awsutil.StringValue(resp))
	return 0
}

func generateCreateServiceInput(cluster string, name string, task string, desire int64) (*ecs.CreateServiceInput, error) {
	if name == "" {
		return nil, errors.New("Error: Require -name option to specify service name.")
	}

	if task == "" {
		return nil, errors.New("Error: Require -task option to attach task definition.")
	}

	input := &ecs.CreateServiceInput{
		Cluster:        &cluster,
		DesiredCount:   &desire,
		ServiceName:    &name,
		TaskDefinition: &task,
	}

	return input, nil
}

// Help show Command implementation method for ServiceListCommand.
func (c *ServiceCreateCommand) Help() string {
	return `
Usage: screwdriver service create [options]
	-cluster        The name of cluster the service belong to.
	-name           The name of service.
	-task           The name of the task definition to attach to the service.
	-desire         The count of the running task to keep.
`
}

// Synopsis Command implementation method for ServiceCreateCommand.
func (c *ServiceCreateCommand) Synopsis() string {
	return "create ECS service"
}
