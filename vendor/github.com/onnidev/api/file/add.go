package file

import (
	"bytes"
	"errors"
	"image"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
	"gopkg.in/mgo.v2/bson"
)

// Add TODO: NEEDS COMMENT INFO
func Add(w http.ResponseWriter, r *http.Request) {
	fileHeader := r.Context().Value(middlewares.FileHeaderKey).([]byte)
	grid := r.Context().Value(middlewares.GridFSRepoKey).(interfaces.GridFSRepo)
	creationTime := time.Now()
	file, err := grid.FS.Create(shared.RandStringBytesRmndr(3))
	defer file.Close()
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	_, err = file.Write(fileHeader)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	img, _, err := image.Decode(bytes.NewReader(fileHeader))
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	b := img.Bounds()
	if b.Max.X < 400 {
		err := errors.New("Image muito pequena")
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	log.Println(http.DetectContentType(fileHeader))
	file.SetContentType(strings.ToUpper(http.DetectContentType(fileHeader)))
	session, err := shared.NewAwsSession()
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
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
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	horario := types.Timestamp(creationTime)
	render.Status(r, http.StatusOK)
	render.JSON(w, r, types.Image{
		FileID:       file.Id().(bson.ObjectId),
		MimeType:     strings.ToUpper(http.DetectContentType(fileHeader)),
		CreationDate: &horario,
	})
}
