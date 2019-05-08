package file

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/pressly/chi/render"
	"gopkg.in/mgo.v2/bson"
)

// AddMin TODO: NEEDS COMMENT INFO
func AddMin(w http.ResponseWriter, r *http.Request) {
	fileHeader := r.Context().Value(middlewares.FileHeaderKey).([]byte)

	img, _, err := image.Decode(bytes.NewReader(fileHeader))
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	nmin, err := strconv.Atoi(chi.URLParam(r, "min"))
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	b := img.Bounds()
	if b.Max.X < nmin {
		err := errors.New(fmt.Sprintf("Imagem menor que %d pixels.", nmin))
		shared.MakeONNiError(w, r, 400, err)
		return
	}

	grid := r.Context().Value(middlewares.GridFSRepoKey).(interfaces.GridFSRepo)
	creationTime := time.Now()
	file, err := grid.FS.Create(shared.RandStringBytesRmndr(3))
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	_, err = file.Write(fileHeader)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	log.Println(http.DetectContentType(fileHeader))
	file.SetContentType(strings.ToUpper(http.DetectContentType(fileHeader)))
	defer file.Close()
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
