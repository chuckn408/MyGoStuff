package main

import (
	"context"
	"encoding/json"

	// "encoding/json"

	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"

	"github.com/ProtonMail/gopenpgp/v2/helper"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func init() {
}

func handler(ctx context.Context, sqsEvent events.SQSEvent) {
	privateKey := (os.Getenv("PRIVATE_KEY"))
	passphrase := []byte(os.Getenv("PASSPHRASE"))
	//s3Dest := os.Getenv("S3_DEST")

	message := ""
	error := false

	if privateKey == "" {
		message += "PRIVATE_KEY env var does not exist"
		error = true
	}
	if passphrase == "" {
		message += "PASSPHRASE env var does not exist"
		error = true
	}
	if error {
		panic(message)
	}
	armor := ""
	decrypted, _ := helper.DecryptMessageArmored(privateKey, passphrase, armor)
	fmt.Println(decrypted)
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)

	downloader := s3manager.NewDownloader(sess)

	for _, sqsRecord := range sqsEvent.Records {
		s3Event := &events.S3Event{}
		err := json.Unmarshal([]byte(sqsRecord.Body), s3Event)
		panic(err)
	}

	for _, record := range s3Event.Records {
		s3RecordMetadata := record.S3

		splitKey := strings.Split(s3RecordMetadata.Object.Key, "/")
		fmt.Println(splitKey)
		file := splitKey[len(splitKey)-1]

		destinationFile, err := client.Create(file)
		if err != nil {
			panic(err)
		}
		defer destinationFile.Close()

		_, err = downloader.Download(destinationFile,
			&s3.GetObjectInput{
				Bucket: aws.String(s3RecordMetadata.Bucket.Name),
				Key:    aws.String(s3RecordMetadata.Object.Key),
			})
		if err != nil {
			panic(err)
		}

	}
}