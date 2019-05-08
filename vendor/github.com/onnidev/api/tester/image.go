package tester

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"os"
	"testing"

	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/types"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
)

// UploadImage is a test that uploads an image
func UploadImage(t *testing.T) *types.Image {
	var image types.Image
	function := func(t *testing.T) {
		var b bytes.Buffer
		w := multipart.NewWriter(&b)
		fImg1, err := os.Open("../file/onni.png")
		defer fImg1.Close()
		assert.Nil(t, err)
		fw, err := w.CreateFormFile("file", "2392u9eh8u23eh823e")
		assert.Nil(t, err)
		_, err = io.Copy(fw, fImg1)
		assert.Nil(t, err)
		w.Close()
		ajax := infra.Ajax{
			Method: "PUT",
			Body:   &b,
			Path:   infra.TestServer.URL + "/file",
			Headers: map[string]string{
				"Content-Type": w.FormDataContentType()}}
		_, body := infra.NewTestRequest(t, ajax)
		err = json.Unmarshal(body, &image)
		assert.NoError(t, err)
		assert.True(t, bson.IsObjectIdHex(image.FileID.Hex()))
	}
	t.Run("Upload image", function)
	return &image
}

// PublishtoS3Image is a test that uploads an image
func PublishtoS3Image(t *testing.T, helper *types.UserClubLoginResponse, img *types.Image) *types.Image {
	var image types.Image
	function := func(t *testing.T) {
		ajax := infra.Ajax{
			Method:  "GET",
			Path:    infra.TestServer.URL + "/file/s3/" + img.FileID.Hex(),
			Headers: helper.Headers()}
		_, body := infra.NewTestRequest(t, ajax)
		err := json.Unmarshal(body, &image)
		assert.NoError(t, err)
		assert.True(t, bson.IsObjectIdHex(image.FileID.Hex()))
	}
	t.Run("Publish image", function)
	return &image
}
