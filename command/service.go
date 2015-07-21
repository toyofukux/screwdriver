package command

import "fmt"

// ServiceCommand is a Command implementation used to
// operate ECS service
type ServiceCommand struct{}

// Run is a Command implementation method for ServiceCommand
func (c *ServiceCommand) Run(argsRow []string) int {
	args := make([]string, len(argsRow))
	copy(args, argsRow)

	if len(args) == 0 {
		fmt.Println(c.Help())
		return 1
	}

	switch args[0] {
	case "list":
		cmd := &ServiceListCommand{}
		return cmd.Run(args[1:])
	case "create":
		cmd := &ServiceCreateCommand{}
		return cmd.Run(args[1:])
	case "update":
		cmd := &ServiceUpdateCommand{}
		return cmd.Run(args[1:])
	default:
		fmt.Println(c.Help())
		return 1
	}
}

// Help show how to use command
func (c *ServiceCommand) Help() string {
	return `
Usage: screwdriver service <subcommand> [options]
Subcommands:
	list        show the list of ECS services
	create      create ECS service
	update      update ECS service
`
}

// Synopsis describe service command overview.
func (c *ServiceCommand) Synopsis() string {
	return "Operate ECS service"
}
