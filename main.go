package main

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

/*

1. Ambil nilai hash dari yarn.lock atau package.json.
2. Check ke bucket, jika tidak ditemukan makan program berhenti.
Jika ditemukan maka di pull
*/

var sess *session.Session

func main() {

	cfg := aws.NewConfig()
	cfg.Region = aws.String("ap-southeast-1")

	sess, _ = session.NewSession(cfg)
	// svc := s3.New(sess)

	// checksum := checksumFunc("./yarn.lock")

	// fileName := fmt.Sprintf("%s.zip", checksum)
	// input := &s3.HeadObjectInput{
	// 	Bucket: aws.String("s-dev.ottenstatic.com"),
	// 	Key:    aws.String(fileName),
	// }

	// _, err := svc.HeadObject(input)
	// if err != nil {
	// 	return
	// }

	// fmt.Println(checksum)

	arg := os.Args[1]

	switch arg {
	case "pull":
		pull()
	case "push":
		push()
	}
}
