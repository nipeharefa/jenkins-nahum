package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/sirupsen/logrus"
)

func push() {

	bucketName := os.Getenv("BUCKET_NAME")

	filename := fmt.Sprintf("otten_web_%s.zip", checksumFunc("./yarn.lock"))
	cc(filename)

	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("failed to open file %q, %v", filename, err)
		return
	}

	defer f.Close()

	buff, _ := ioutil.ReadAll(f)

	reader := bytes.NewReader(buff)

	input := &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filename),
		Body:   aws.ReadSeekCloser(reader),
	}

	svc := s3.New(sess)

	result, err := svc.PutObject(input)
	if err != nil {
		logrus.Error(err)
	}

	defer os.Remove(filename)

	fmt.Println(result)
}
