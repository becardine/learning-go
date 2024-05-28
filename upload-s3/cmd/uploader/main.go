package main

import (
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	s3Client *s3.S3
	s3Bucket string
)

func init() {
	// This will be executed before main() function
	session, err := session.NewSession(
		&aws.Config{
			Region: aws.String("us-west-2"),
			Credentials: credentials.NewStaticCredentials(
				"AKID",
				"SECRET",
				"TOKEN",
			),
		},
	)
	if err != nil {
		panic(err)
	}

	s3Client = s3.New(session)
	s3Bucket = "my-bucket-example"
}

func main() {
	// This will be executed when the program is run
	dir, err := os.Open("./tmp")
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	for {
		files, err := dir.Readdir(1)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("Failed to read directory: %s\n", err.Error())
			continue
		}

		for _, file := range files {
			uploadFile(file.Name())
		}
	}
}

func uploadFile(filename string) {
	completeFilename := fmt.Sprintf("./tmp/%s", filename)
	fmt.Printf("Uploading file: %s\n", completeFilename)
	file, err := os.Open(completeFilename)
	if err != nil {
		fmt.Printf("Failed to open file: %s\n", completeFilename)
		return
	}
	defer file.Close()

	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(filename),
		Body:   file,
	})
	if err != nil {
		fmt.Printf("Failed to upload file: %s\n", completeFilename)
		return
	}

	fmt.Printf("File uploaded: %s\n", completeFilename)
}
