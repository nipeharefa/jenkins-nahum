package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/sirupsen/logrus"
)

func push() {
	c := compress()

	bucketName := os.Getenv("BUCKET_NAME")

	zipReader := bytes.NewReader(c.Bytes())

	fileName := fmt.Sprintf("otten_web_%s.zip", checksumFunc("./yarn.lock"))

	input := &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
		Body:   aws.ReadSeekCloser(zipReader),
		ACL:    aws.String(s3.ObjectCannedACLPublicRead),
	}

	svc := s3.New(sess)

	result, err := svc.PutObject(input)
	if err != nil {
		logrus.Error(err)
	}
	fmt.Println(result)
}
