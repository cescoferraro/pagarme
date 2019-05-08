package shared

import (
	"math/rand"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

// NewAwsSession get an aws session for onni
func NewAwsSession() (*session.Session, error) {
	acess := "AKIAIC5SILAYD5BKEALA"
	secret := "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"
	creds := credentials.NewStaticCredentials(acess, secret, "")
	_, err := creds.Get()
	if err != nil {
		return nil, err
	}
	sess, err := session.NewSession(&aws.Config{
		Credentials: creds,
		Region:      aws.String("sa-east-1"),

		LogLevel: aws.LogLevel(aws.LogOff)})
	return sess, err
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// RandStringBytesRmndr TODO: NEEDS COMMENT INFO
func RandStringBytesRmndr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
