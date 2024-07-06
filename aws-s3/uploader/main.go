package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

var (
	s3Client  *s3.S3
	s3Bucket  = "go-expert-bucket-example"
	waitGroup *sync.WaitGroup
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
	s3Client = s3.New(awsSession)
	waitGroup = &sync.WaitGroup{}
}

func main() {
	directory, err := os.Open("aws-s3/tmp")
	if err != nil {
		log.Fatal(err)
	}
	defer directory.Close()

	uploadControl := make(chan struct{}, 100)
	errorControl := make(chan string)

	go handleError(uploadControl, errorControl)

	for {
		files, err := directory.Readdir(1)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		waitGroup.Add(1)
		uploadControl <- struct{}{}
		go uploadFile(files[0].Name(), uploadControl, errorControl)
	}
	waitGroup.Wait()
}

func uploadFile(fileName string, uploadControl <-chan struct{}, errorControl chan string) {
	defer func() {
		waitGroup.Done()
		<-uploadControl
	}()
	completedFileName := fmt.Sprintf("aws-s3/tmp/%s", fileName)
	fmt.Printf("Uploading file %s\n", completedFileName)
	file, err := os.Open(completedFileName)
	if err != nil {
		log.Printf("Error opening file %s", completedFileName)
		errorControl <- fileName
		return
	}
	defer file.Close()
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Body:        file,
		Bucket:      aws.String(s3Bucket),
		ContentType: aws.String("text/plain"),
		Expires:     aws.Time(time.Now().Add(5 * time.Minute)),
		Key:         aws.String(completedFileName),
	})
	if err != nil {
		log.Printf("Error uploading file %s - %s", completedFileName, err)
		errorControl <- fileName
		return
	}
	log.Printf("Successfully uploaded file %s", completedFileName)
}

func handleError(uploadControl chan struct{}, errorControl chan string) {
	for {
		select {
		case fileName := <-errorControl:
			uploadControl <- struct{}{}
			waitGroup.Add(1)
			go uploadFile(fileName, uploadControl, errorControl)
		}
	}
}
