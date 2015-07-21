package command

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/takasing/screwdriver/aws"
	"github.com/takasing/screwdriver/utils"
)

// TaskRegisterCommand is a Command implementation used to
// register ECS task definition.
type TaskRegisterCommand struct{}

// Run is Command implementation method for TaskRegisterCommand.
func (c *TaskRegisterCommand) Run(args []string) int {
	var family, path string
	flags := flag.NewFlagSet("regsiter", flag.ContinueOnError)
	flags.StringVar(&family, "family", "", "task definition family name.")
	flags.StringVar(&path, "path", "task.yml", "the path of configuration file.")
	flags.Usage = func() { c.Help() }

	if err := flags.Parse(args); err != nil {
		utils.ErrorOutputf("Error parsing CLI flags: %s", err)
		fmt.Println(c.Help())
		return 1
	}

	input, err := generateRegisterTaskDefinitionInput(family, path)
	if err != nil {
		utils.ErrorOutput(err.Error())
		return 1
	}

	client := aws.GetClient()
	resp, err := client.TaskRegister(input)
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

	fmt.Printf("Success register task %s.", family)
	fmt.Println(awsutil.StringValue(resp))
	return 0
}

func generateRegisterTaskDefinitionInput(family string, path string) (*ecs.RegisterTaskDefinitionInput, error) {
	envs := utils.LoadScrewEnvs()
	data, err := ioutil.ReadFile(path)
	if err != nil {
		utils.ErrorOutput(err.Error())
		e := fmt.Errorf("Error parsing configuration file(%s)", path)
		return nil, e
	}

	b, err := utils.ExpandTemplate(data, envs)
	if err != nil {
		return nil, err
	}

	containers, err := utils.BindYml(b)
	if err != nil {
		return nil, err
	}

	if family == "" {
		return nil, errors.New("Missing family. Show help of 'task register' command.")
	}

	input := &ecs.RegisterTaskDefinitionInput{
		ContainerDefinitions: containers,
		Family:               &family,
	}

	return input, nil
}

// Help show Command implementation method for TaskRegisterCommand.
func (c *TaskRegisterCommand) Help() string {
	return `
Usage: screwdriver task def [options]
Options:
	-family(string)    The family name you want to register.
	-path(string)      The file path you want to register. Now support Yaml.
`
}

// Synopsis Command implementation method for TaskRegisterCommand.
func (c *TaskRegisterCommand) Synopsis() string {
	return "Register ECS task definitions"
}
