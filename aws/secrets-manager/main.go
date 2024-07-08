package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"log/slog"
	"pos-graduacao-go-lang/aws/session"
)

var (
	secretsManagerClient *secretsmanager.SecretsManager
)

func init() {
	secretsManagerClient = secretsmanager.New(session.AwsSession)
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
