package command

import (
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/service/ecs"
)

// TaskDefsCommand is a Command implementation used to
// get ECS task definition list.
type TaskDefsCommand struct {
}

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
		fmt.Println(err.Error())
		fmt.Println(c.Help())
		return 1
	}

	input := generateInput(familyPrefix, max, nextToken, desc, status)

	// TODO: initialize ECS client here
	// TODO: extends Command implementation with initialize()

	resp, err := awsClient.TaskDefs(input)
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			fmt.Println(err.Error())
		}
		return 1
	}
	fmt.Println(awsutil.StringValue(resp))
	return 0
}

func generateInput(prefix string, max int64, token string, desc bool, status string) *ecs.ListTaskDefinitionsInput {
	var input *ecs.ListTaskDefinitionsInput
	sort := "ASC"
	if desc {
		sort = "DESC"
	}
	input = &ecs.ListTaskDefinitionsInput{
		MaxResults: aws.Long(max),
		NextToken:  aws.String(token),
		Sort:       aws.String(sort),
		Status:     aws.String(status),
	}
	// if set empty string to FamilyPrefix, return empty list
	if prefix != "" {
		input.FamilyPrefix = aws.String(prefix)
	}
	return input
}

// Help show Command implementation method for TaskDefsCommand.
func (c *TaskDefsCommand) Help() string {
	helpText := `
Usage: screw task def [options]
Options:
	-prefix(strig)     The full family name that you want to filter.
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
