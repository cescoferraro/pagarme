package onni

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
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
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2/bson"
)

// ImageONNi TODO: NEEDS COMMENT INFO
func ImageONNi(ctx context.Context, id string) (types.Image, error) {
	fs, ok := ctx.Value(middlewares.GridFSRepoKey).(interfaces.GridFSRepo)
	if !ok {
		err := errors.New("bug assert")
		return types.Image{}, err
	}
	file, err := fs.FS.OpenId(bson.ObjectIdHex(id))
	if err != nil {
		return types.Image{}, err

	}
	defer file.Close()
	log.Println("file opened")
	fileHeader := make([]byte, file.Size())
	_, err = file.Read(fileHeader)
	if err != nil {
		return types.Image{}, err
	}
	horario := types.Timestamp(time.Now())
	return types.Image{
		FileID:       file.Id().(bson.ObjectId),
		MimeType:     strings.Replace(strings.ToUpper(http.DetectContentType(fileHeader)), "/", "_", -1),
		CreationDate: &horario,
	}, nil
}

// ImageONNiFromBytes TODO: NEEDS COMMENT INFO
func ImageONNiFromBytes(ctx context.Context, fileHeader []byte) (types.Image, error) {
	image := types.Image{}
	grid := ctx.Value(middlewares.GridFSRepoKey).(interfaces.GridFSRepo)
	creationTime := time.Now()
	file, err := grid.FS.Create(shared.RandStringBytesRmndr(3))
	defer file.Close()
	if err != nil {
		return image, nil
	}
	_, err = file.Write(fileHeader)
	if err != nil {
		return image, nil
	}
	file.SetContentType(strings.ToUpper(http.DetectContentType(fileHeader)))
	session, err := shared.NewAwsSession()
	if err != nil {
		return image, nil
	}
	id := file.Id().(bson.ObjectId)
	_, err = s3.New(session).PutObject(&s3.PutObjectInput{
		// TODO: Create a new bucket for tests
		Bucket:               aws.String("onni-medium-images"),
		Key:                  aws.String(id.Hex()),
		ACL:                  aws.String("public-read"),
		Body:                 bytes.NewReader(fileHeader),
		ContentLength:        aws.Int64(int64(len(fileHeader))),
		ContentType:          aws.String(http.DetectContentType(fileHeader)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})
	if err != nil {
		return image, nil
	}
	horario := types.Timestamp(creationTime)
	return types.Image{
		FileID:       id,
		MimeType:     strings.Replace(strings.ToUpper(http.DetectContentType(fileHeader)), "/", "_", -1),
		CreationDate: &horario,
	}, nil
}

// ImageONNiCombo TODO: NEEDS COMMENT INFO
func ImageONNiCombo(ctx context.Context, data string) (types.Image, error) {
	log.Println("dentro da function")
	image := types.Image{}
	if bson.IsObjectIdHex(data) {
		if viper.GetString("env") == "homolog" {
			log.Println("robando no homolog e garfiando uma foto existent")
			image, err := ImageONNi(ctx, "58c2eb58b80dff716c8a2f59")
			if err != nil {
				return image, nil
			}
			return image, nil
		}
		log.Println("lendo uma imagem existente")
		image, err := ImageONNi(ctx, data)
		if err != nil {
			return image, err
		}
		return image, nil
	}
	log.Println("tentando criar uma nova imagemmm")
	i := strings.Index(data, ",")
	sDec, err := base64.StdEncoding.DecodeString(data[i+1:])
	if err != nil {
		return types.Image{}, err
	}
	image, err = ImageONNiFromBytes(ctx, sDec)
	if err != nil {
		return image, err
	}
	return image, nil
}
