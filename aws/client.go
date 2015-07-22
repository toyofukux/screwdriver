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

// TaskDefs return the list of ECS task definitions
func (c *Client) TaskDefs(input *ecs.ListTaskDefinitionsInput) (*ecs.ListTaskDefinitionsOutput, error) {
	return c.ecs.ListTaskDefinitions(input)
}

// TaskRegister regsiter ECS task definitions
func (c *Client) TaskRegister(input *ecs.RegisterTaskDefinitionInput) (*ecs.RegisterTaskDefinitionOutput, error) {
	return c.ecs.RegisterTaskDefinition(input)
}

// ServiceList return the list of ECS services
func (c *Client) ServiceList(input *ecs.ListServicesInput) (*ecs.ListServicesOutput, error) {
	return c.ecs.ListServices(input)
}

// ServiceCreate create ECS service
func (c *Client) ServiceCreate(input *ecs.CreateServiceInput) (*ecs.CreateServiceOutput, error) {
	return c.ecs.CreateService(input)
}

// ServiceUpdate update ECS service
func (c *Client) ServiceUpdate(input *ecs.UpdateServiceInput) (*ecs.UpdateServiceOutput, error) {
	return c.ecs.UpdateService(input)
}

// ClusterList show the list of ECS clusters
func (c *Client) ClusterList(input *ecs.ListClustersInput) (*ecs.ListClustersOutput, error) {
	return c.ecs.ListClusters(input)
}

// ClusterCreate create ECS cluster
func (c *Client) ClusterCreate(input *ecs.CreateClusterInput) (*ecs.CreateClusterOutput, error) {
	return c.ecs.CreateCluster(input)
}

// ClusterDelete delete ECS cluster
func (c *Client) ClusterDelete(input *ecs.DeleteClusterInput) (*ecs.DeleteClusterOutput, error) {
	return c.ecs.DeleteCluster(input)
}
