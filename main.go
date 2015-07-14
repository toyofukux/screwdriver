package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/takasing/screwdriver/config"
)

func main() {
	svc := ecs.New(config.Config)
	params := &ecs.ListTaskDefinitionFamiliesInput{
		FamilyPrefix: aws.String("golang"),
		MaxResults:   aws.Long(3),
		NextToken:    aws.String("golang"),
	}
	resp, err := svc.ListTaskDefinitionFamilies(params)
	if err != nil {
		panic(err)
	}
	fmt.Println(awsutil.StringValue(resp))
}
