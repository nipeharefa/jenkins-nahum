package main

import (
	"bytes"
	"compress/flate"
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
	log.Info("Pull")
	bucketName := os.Getenv("BUCKET_NAME")

	filename := fmt.Sprintf("otten_web_%s.zip", checksumFunc("./yarn.lock"))
	svc := s3.New(sess)

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

	z := archiver.Zip{
		CompressionLevel:       flate.DefaultCompression,
		MkdirAll:               true,
		SelectiveCompression:   true,
		ContinueOnError:        false,
		OverwriteExisting:      false,
		ImplicitTopLevelFolder: true,
	}

	err = z.Unarchive(filename, ".")
	if err != nil {
		log.Error(err)
	}
}
