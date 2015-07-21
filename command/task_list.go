package command

import (
	"flag"
	"fmt"
)

// TaskListCommand is a Command implementation used to
// get ECS task definition list.
type TaskListCommand struct {
}

// Run is Command implementation method for TaskListCommand.
func (c *TaskListCommand) Run(argsRow []string) int {
	flags := flag.NewFlagSet("task", flag.ContinueOnError)
	flags.Usage = func() { c.Help() }
	if err := flags.Parse(argsRow); err != nil {
		fmt.Println(err.Error())
		return 1
	}
	args := flags.Args()
	fmt.Println(args)
	return 0
}

// Help show Command implementation method for TaskListCommand.
func (c *TaskListCommand) Help() string {
	helpText := `
Usage: screwdriver task list [options]
`
	return helpText
}

// Synopsis Command implementation method for TaskListCommand.
func (c *TaskListCommand) Synopsis() string {
	return "Show ECS task list"
}
