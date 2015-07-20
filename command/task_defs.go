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

// TaskDefsCommand is a Command implementation used to
// get ECS task definition list.
type TaskDefsCommand struct{}

// Run is Command implementation method for TaskDefsCommand.
func (c *TaskDefsCommand) Run(args []string) int {
	var familyPrefix, nextToken, status string
	var max int64
	var desc bool
	flags := flag.NewFlagSet("defs", flag.ContinueOnError)
	flags.StringVar(&familyPrefix, "prefix", "", "task definition family prefix string(require full family name)")
	flags.Int64Var(&max, "max", 10, "the number of task definition to show")
	flags.StringVar(&nextToken, "next", "", "next token")
	flags.BoolVar(&desc, "desc", false, "sort task definitions ASC or DESC")
	flags.StringVar(&status, "status", "ACTIVE", "task definition status")
	flags.Usage = func() { c.Help() }

	if err := flags.Parse(args); err != nil {
		utils.ErrorOutputf("Error parsing CLI flags: %s", err)
		fmt.Println(c.Help())
		return 1
	}

	input := generateListTaskDefinitionsInput(familyPrefix, max, nextToken, desc, status)

	client := aws.GetClient()
	resp, err := client.TaskDefs(input)
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

func generateListTaskDefinitionsInput(prefix string, max int64, token string, desc bool, status string) *ecs.ListTaskDefinitionsInput {
	var input *ecs.ListTaskDefinitionsInput
	sort := "ASC"
	if desc {
		sort = "DESC"
	}
	input = &ecs.ListTaskDefinitionsInput{
		MaxResults: &max,
		NextToken:  &token,
		Sort:       &sort,
		Status:     &status,
	}
	// if set empty string to FamilyPrefix, return empty list
	if prefix != "" {
		input.FamilyPrefix = &prefix
	}
	return input
}

// Help show Command implementation method for TaskDefsCommand.
func (c *TaskDefsCommand) Help() string {
	helpText := `
Usage: screw task def [options]
Options:
	-prefix(string)    The full family name that you want to filter.
	-max(int)          The total number of items to return.
	-next(string)      A token to specify where to start paginating.
	-desc(bool)        The order in which to sort the results.
	                   You want to show list DESC, specify.
	-status(string)    The task definition status that you want to filter.
	                   Status is ACTIVE or INACTIVE.
`
	return helpText
}

// Synopsis Command implementation method for TaskDefsCommand.
func (c *TaskDefsCommand) Synopsis() string {
	return "Show ECS task definitions"
}
