package file

import (
	"bytes"
	"log"
	"net/http"
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

// Aws TODO: NEEDS COMMENT INFO
func Aws(w http.ResponseWriter, r *http.Request) {

	fs := r.Context().Value(middlewares.GridFSRepoKey).(interfaces.GridFSRepo)
	id := chi.URLParam(r, "id")
	log.Println(id)
	file, err := fs.FS.OpenId(bson.ObjectIdHex(id))
	if err != nil {
		render.Status(r, http.StatusExpectationFailed)
		render.JSON(w, r, err.Error())
		return
	}
	defer file.Close()

	session, err := shared.NewAwsSession()
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}

	fileHeader := make([]byte, file.Size())
	_, err = file.Read(fileHeader)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)

		return
	}
	id = file.Id().(bson.ObjectId).Hex()
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
	horario := types.Timestamp(time.Now())
	render.Status(r, http.StatusOK)
	render.JSON(w, r, types.Image{
		FileID:       file.Id().(bson.ObjectId),
		MimeType:     strings.ToUpper(http.DetectContentType(fileHeader)),
		CreationDate: &horario,
	})
}
