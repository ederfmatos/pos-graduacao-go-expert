package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"log/slog"
)

var (
	sesClient *ses.SES
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
	sesClient = ses.New(awsSession)
}

func main() {
	email, err := sesClient.SendEmail(&ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{aws.String("ederfmatos@gmail.com")},
		},
		Message: &ses.Message{
			Body: &ses.Body{Text: &ses.Content{
				Charset: aws.String("text/plain"),
				Data:    aws.String("Hello, World!"),
			}},
			Subject: &ses.Content{
				Charset: aws.String("text/plain"),
				Data:    aws.String("Subject!"),
			},
		},
		Source:    aws.String("hello@example.com"),
		SourceArn: nil,
		Tags:      nil,
	})
	if err != nil {
		panic(err)
	}
	slog.Info("Email sent:", email)
}
