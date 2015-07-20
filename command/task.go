package command

import "github.com/takasing/screwdriver/utils"

// TaskCommand is a Command implementation used to
// operate ECS task.
type TaskCommand struct {
}

// Run is a Command implementation method for TaskCommand.
func (c *TaskCommand) Run(argsRow []string) int {
	args := make([]string, len(argsRow))
	copy(args, argsRow)

	if len(args) == 0 {
		utils.ErrorOutput(c.Help())
		return 1
	}

	switch args[0] {
	case "defs":
		cmd := &TaskDefsCommand{}
		return cmd.Run(args[1:])
	case "list":
		cmd := &TaskListCommand{}
		return cmd.Run(args[1:])
	case "register":
		cmd := &TaskRegisterCommand{}
		return cmd.Run(args[1:])
	default:
		utils.ErrorOutput(c.Help())
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
