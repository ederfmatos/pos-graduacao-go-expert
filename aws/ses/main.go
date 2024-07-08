package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"
	"log/slog"
	"pos-graduacao-go-lang/aws/session"
)

var (
	sesClient *ses.SES
)

func init() {
	sesClient = ses.New(session.AwsSession)
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
