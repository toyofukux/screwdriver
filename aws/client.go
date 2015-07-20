package aws

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/ecs"
)

var (
	client *Client
)

// Client is the client to control AWS.
type Client struct {
	ecs *ecs.ECS
}

// GetClient configure and return initialized client.
func GetClient() *Client {
	if client != nil {
		return client
	}
	var c Client
	fmt.Println("[INFO] Initializing ECS Connection")
	c.ecs = ecs.New(&aws.Config{
		Credentials: credentials.NewEnvCredentials(),
	})
	client = &c
	return client
}

// TaskDefs return the list of task definitions
func (c *Client) TaskDefs(input *ecs.ListTaskDefinitionsInput) (*ecs.ListTaskDefinitionsOutput, error) {
	return c.ecs.ListTaskDefinitions(input)
}

// TaskRegister regsiter ECS task definitions
func (c *Client) TaskRegister(input *ecs.RegisterTaskDefinitionInput) (*ecs.RegisterTaskDefinitionOutput, error) {
	return c.ecs.RegisterTaskDefinition(input)
}
