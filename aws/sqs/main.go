package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"log/slog"
	"pos-graduacao-go-lang/aws/session"
	"sync"
)

var (
	sqsClient *sqs.SQS
	waitGroup *sync.WaitGroup
)

func init() {
	sqsClient = sqs.New(session.AwsSession)
	waitGroup = &sync.WaitGroup{}
}

func main() {
	queues, err := sqsClient.ListQueues(nil)
	if err != nil {
		panic(err)
	}
	slog.Info("Queues:", queues)
	queue := "my-queue-2"

	result, err := sqsClient.CreateQueue(&sqs.CreateQueueInput{
		QueueName: aws.String(queue),
		Attributes: map[string]*string{
			"DelaySeconds":           aws.String("20"),
			"MessageRetentionPeriod": aws.String("86400"),
		},
	})
	if err != nil {
		panic(err)
	}
	slog.Info("Created queue:", *result.QueueUrl)

	getQueueOutput, err := sqsClient.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(queue),
	})
	if err != nil {
		panic(err)
	}
	slog.Info("Get queue:", getQueueOutput)

	url := *getQueueOutput.QueueUrl

	go ReceiveMessage(url)
	err = SendMessage("hello", url)
	if err != nil {
		panic(err)
	}

	waitGroup.Wait()
}

func SendMessage(message string, url string) error {
	input := &sqs.SendMessageInput{
		MessageBody: aws.String(message),
		QueueUrl:    aws.String(url),
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"WeeksOn": {
				DataType:    aws.String("Number"),
				StringValue: aws.String("6"),
			},
		},
	}
	slog.Info("Sending message:", message)
	output, err := sqsClient.SendMessage(input)
	if err != nil {
		return err
	}
	waitGroup.Add(1)
	slog.Info("Message sent:", *output.MessageId)
	return nil
}

func ReceiveMessage(url string) {
	for {
		msgResult, err := sqsClient.ReceiveMessage(&sqs.ReceiveMessageInput{
			MessageAttributeNames: []*string{
				aws.String(sqs.QueueAttributeNameAll),
			},
			QueueUrl:            aws.String(url),
			MaxNumberOfMessages: aws.Int64(10),
			WaitTimeSeconds:     aws.Int64(6),
		})
		if err != nil {
			slog.Info("Error receiving messages:", err)
			continue
		}
		if msgResult == nil || len(msgResult.Messages) == 0 {
			slog.Info("Received no messages")
			continue
		}
		slog.Info("Message:", msgResult)
		for _, message := range msgResult.Messages {
			slog.Info("Message:", *message.Body, message.String())

			_, _ = sqsClient.DeleteMessage(&sqs.DeleteMessageInput{
				QueueUrl:      aws.String(url),
				ReceiptHandle: aws.String(*message.ReceiptHandle),
			})
		}
	}
}
