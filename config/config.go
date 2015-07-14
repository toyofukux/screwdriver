package config

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

var (
	// Config is AWS Config has AWS_ACCESS_KEY_ID, and AWS_SECRET_ACCESS_KEY
	Config *aws.Config
)

func init() {
	Config = &aws.Config{
		Credentials: credentials.NewEnvCredentials(),
	}
}
