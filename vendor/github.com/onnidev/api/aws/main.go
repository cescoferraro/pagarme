package aws

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/onnidev/api/shared"
)

var allPhotos []s3.Object

// MYS3 is the global s3 bucket
var MYS3 *s3.S3

// Run is awesome
func Run() {
	done := make(chan bool, 1)
	session, err := shared.NewAwsSession()
	if err != nil {
		os.Exit(1)
	}
	MYS3 = s3.New(session)
	params := &s3.ListObjectsInput{
		Bucket: aws.String("onni-backup"),
	}
	resp, err := MYS3.ListObjects(params)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, key := range resp.Contents {
		allPhotos = append(allPhotos, *key)
		log.Println(*key.Key)
	}

	done <- true
}
