package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"log/slog"
)

var (
	secretsManagerClient *secretsmanager.SecretsManager
)

func init() {
	config := &aws.Config{
		Region:           aws.String("us-east-1"),
		Credentials:      credentials.NewStaticCredentials("test", "test", ""),
		Endpoint:         aws.String("http://localhost:4566"),
		S3ForcePathStyle: aws.Bool(true),
	}
	awsSession, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	secretsManagerClient = secretsmanager.New(awsSession)
}

func main() {
	password, err := secretsManagerClient.GetRandomPassword(nil)
	if err != nil {
		panic(err)
	}
	slog.Info("Password:", *password.RandomPassword)
	secretValue, err := secretsManagerClient.GetSecretValue(&secretsmanager.GetSecretValueInput{
		SecretId: aws.String("test-secret"),
	})
	if err != nil {
		panic(err)
	}
	slog.Info("Secret Value:", *secretValue)
}
