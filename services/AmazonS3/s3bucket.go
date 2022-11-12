package amazons3

import (
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/joho/godotenv"
)

type S3 struct {
	res *s3manager.UploadOutput
}

func (s S3) GetBucketName() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(".env file loading error - ", err)
		return ""
	}

	BUCKET := os.Getenv("BUCKET_NAME")

	return BUCKET
}

func (s S3) UploadToS3(file *multipart.FileHeader, filename string) (string, error) {
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: "default",
		Config: aws.Config{
			Region: aws.String("ap-south-1"),
		},
	})

	if err != nil {
		fmt.Printf("Failed to initialize new session: %v", err)
		return "", errors.New("error in Initializing S3 bucket")
	}

	bucketName := s.GetBucketName()
	if bucketName == "" {
		return "", errors.New("error in getting env")
	}

	uploader := s3manager.NewUploader(sess)

	// change the fileheader to file
	data, _ := file.Open()
	s.res, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filename),
		Body:   data,
	})
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	location := s.res.Location

	return location, nil
}
