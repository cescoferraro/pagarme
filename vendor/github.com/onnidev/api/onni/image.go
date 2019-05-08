package onni

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"gopkg.in/mgo.v2/bson"
)

// CreateImage TODO: NEEDS COMMENT INFO
func CreateImage(ctx context.Context, fileHeader []byte) (types.Image, error) {
	image := types.Image{}
	grid, ok := ctx.Value(middlewares.GridFSRepoKey).(interfaces.GridFSRepo)
	if !ok {
		err := errors.New("bug assert")
		return image, err
	}
	creationTime := time.Now()
	file, err := grid.FS.Create(shared.RandStringBytesRmndr(3))
	if err != nil {
		return image, err
	}
	_, err = file.Write(fileHeader)
	if err != nil {
		return image, err
	}
	file.SetContentType(strings.ToUpper(http.DetectContentType(fileHeader)))
	file.Close()
	session, err := shared.NewAwsSession()
	if err != nil {
		return image, err
	}
	id := file.Id().(bson.ObjectId).Hex()
	_, err = s3.New(session).PutObject(&s3.PutObjectInput{
		// TODO: Create a new bucket for tests
		Bucket:               aws.String("onni-medium-images"),
		Key:                  aws.String(id),
		ACL:                  aws.String("public-read"),
		Body:                 bytes.NewReader(fileHeader),
		ContentLength:        aws.Int64(int64(len(fileHeader))),
		ContentType:          aws.String(http.DetectContentType(fileHeader)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})
	if err != nil {
		return image, err
	}
	horario := types.Timestamp(creationTime)
	return types.Image{
		FileID:       file.Id().(bson.ObjectId),
		MimeType:     strings.Replace(strings.ToUpper(http.DetectContentType(fileHeader)), "/", "_", -1),
		CreationDate: &horario,
	}, nil
}
