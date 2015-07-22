package command

import "fmt"

// ClusterCommand is a Command implementation used to
// operate ECS cluster.
type ClusterCommand struct{}

// Run is a Command implementation method for ClusterCommand
func (c *ClusterCommand) Run(argsRow []string) int {
	args := make([]string, len(argsRow))
	copy(args, argsRow)

	if len(args) == 0 {
		fmt.Println(c.Help())
		return 1
	}

	switch args[0] {
	case "list":
		cmd := &ClusterListCommand{}
		return cmd.Run(args[1:])
	case "create":
		cmd := &ClusterCreateCommand{}
		return cmd.Run(args[1:])
	case "delete":
		cmd := &ClusterDeleteCommand{}
		return cmd.Run(args[1:])
	default:
		fmt.Println(c.Help())
		return 1
	}
}

// Help show how to use command
func (c *ClusterCommand) Help() string {
	return `
Usage: screwdriver cluster <subcommand> [options]
Subcommands:
	list        show the list of ECS cluster
	create      create ECS cluster
	delete      delete ECS service
`
}

// Synopsis describe service command overview.
func (c *ClusterCommand) Synopsis() string {
	return "Operate ECS cluster"
}
