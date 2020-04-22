package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/mholt/archiver/v3"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func pull() {
	fmt.Println("Looking up build cache...")
	bucketName := os.Getenv("BUCKET_NAME")

	filename := fmt.Sprintf("otten_web_%s.zip", checksumFunc("./yarn.lock"))
	svc := s3.New(sess)

	isExist := headObject(filename)

	if !isExist {
		fmt.Println("Build cache doest not exist...")
		return
	}

	inputObj := &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filename),
	}

	res, err := svc.GetObject(inputObj)
	if err != nil {
		logrus.Error(err)
		return
	}

	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logrus.Error(err)
	}

	reader := bytes.NewReader(buf)

	f, err := os.Create(filename)
	if err != nil {
		logrus.Error(err)
	}

	reader.WriteTo(f)

	f.Close()

	err = archiver.Unarchive(filename, ".")
	if err != nil {
		log.Error(err)
	}
}

func headObject(filename string) bool {

	bucketName := os.Getenv("BUCKET_NAME")

	inputObject := &s3.HeadObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filename),
	}

	svc := s3.New(sess)

	_, err := svc.HeadObject(inputObject)

	return err == nil
}
