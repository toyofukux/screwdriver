package aws

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/ecs"
)

// Client is the client to control AWS.
type Client struct {
	ecs *ecs.ECS
}

// NewClient configure and return initialized client.
func NewClient() (interface{}, error) {
	var client Client
	fmt.Println("[INFO] Initializing ECS Connection")
	client.ecs = ecs.New(&aws.Config{
		Credentials: credentials.NewEnvCredentials(),
	})
	return &client, nil
}

// TaskDefs return the list of task definitions
func (c *Client) TaskDefs(input *ecs.ListTaskDefinitionsInput) (*ecs.ListTaskDefinitionsOutput, error) {
	return c.ecs.ListTaskDefinitions(input)
}
