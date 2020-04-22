package main

import (
	"fmt"
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
var build string

func main() {

	fmt.Println("Jenkins tool ", build)
	cfg := aws.NewConfig()
	cfg.Region = aws.String("ap-southeast-1")

	sess, _ = session.NewSession(cfg)
	arg := os.Args[1]

	switch arg {
	case "pull":
		pull()
	case "push":
		push()
	}
}
