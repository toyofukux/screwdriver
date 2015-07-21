package command

// ServiceUpdateCommand is a Command implementation used to
// update ECS service
type ServiceUpdateCommand struct{}

// Run is Command implementation method for ServiceUpdateCommand.
func (c *ServiceUpdateCommand) Run(args []string) int {
	return 0
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
