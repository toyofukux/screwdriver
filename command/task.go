package command

import "fmt"

// TaskCommand is a Command implementation used to
// operate ECS task.
type TaskCommand struct {
}

// Run is a Command implementation method for TaskCommand.
func (c *TaskCommand) Run(argsRow []string) int {
	args := make([]string, len(argsRow))
	copy(args, argsRow)

	if len(args) == 0 {
		fmt.Println(c.Help())
		return 1
	}

	switch args[0] {
	case "defs":
		cmd := &TaskDefsCommand{}
		return cmd.Run(args[1:])
	case "list":
		cmd := &TaskListCommand{}
		return cmd.Run(args[1:])
	default:
		fmt.Println(c.Help())
		return 1
	}
}

// Help show how to use command.
func (c *TaskCommand) Help() string {
	helpText := `
Usage: screw task <subcommand> [options]
`
	return helpText
}

// Synopsis describe task command overview.
func (c *TaskCommand) Synopsis() string {
	return "Operate ECS task"
}
