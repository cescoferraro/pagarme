package tester

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"image/color"
	"log"
	"math"
	"os"
	"testing"

	"github.com/carlogit/phash"
	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/onni"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/types"
	"github.com/stretchr/testify/assert"
)

// IsFacebookImageTest TODO: NEEDS COMMENT INFO
func IsFacebookImageTest(t *testing.T, fbid string) bool {
	result := false
	function := func(t *testing.T) {
		if fbid != "" {
			log.Println(fbid)
			var err error

			byt, err := onni.ImageBytesFromFB(fbid)
			assert.NoError(t, err)
			result, err = onni.IsFacebookImage(byt)
			assert.NoError(t, err)
			return
		}
		// assert.NoError(t, errors.New("sou um erro"))
	}
	t.Run("PhaseTest", function)
	return result
}

// PhaseTest TODO: NEEDS COMMENT INFO
func PhaseTest(t *testing.T) {
	function := func(t *testing.T) {
		fbDefault, err := shared.Asset("public/fbdefault.jpg")
		assert.NoError(t, err)
		gopher, err := shared.Asset("public/gopher.jpg")
		assert.NoError(t, err)
		ahash, err := phash.GetHash(bytes.NewReader(fbDefault))
		assert.NoError(t, err)
		bhash, err := phash.GetHash(bytes.NewReader(gopher))
		assert.NoError(t, err)
		distance := phash.GetDistance(ahash, bhash)
		fmt.Printf("%d\n", distance)
		distance = phash.GetDistance(ahash, ahash)
		fmt.Printf("%d\n", distance)
	}
	t.Run("PhaseTest", function)
}

// FacebookPictureTest TODO: NEEDS COMMENT INFO
func FacebookPictureTest(t *testing.T) {
	function := func(t *testing.T) {
		fbDefault, err := shared.Asset("public/fbdefault.jpg")
		assert.NoError(t, err)
		source, _, err := image.Decode(bytes.NewReader(fbDefault))
		gopher, err := shared.Asset("public/gopher.jpg")
		assert.NoError(t, err)
		target, _, err := image.Decode(bytes.NewReader(gopher))
		diff, err := CompareImages(target, source)
		fmt.Printf("%d\n", diff)
		diff, err = CompareImages(source, source)
		fmt.Printf("%d\n", diff)
	}
	t.Run("FacebookPictureTest", function)
}

// CompareImages TODO: NEEDS COMMENT INFO
func CompareImages(target, source image.Image) (int64, error) {
	targetBounds := target.Bounds()
	sourceBounds := source.Bounds()
	if !boundsMatch(targetBounds, sourceBounds) {
		return int64(0), errors.New("facebook")
	}

	var diff int64
	for y := targetBounds.Min.Y; y < targetBounds.Max.Y; y++ {
		for x := targetBounds.Min.X; x < targetBounds.Max.X; x++ {
			diff += compareColor(target.At(x, y), source.At(x, y))
		}
	}
	return diff, nil
}
func compareColor(a, b color.Color) (diff int64) {
	r1, g1, b1, a1 := a.RGBA()
	r2, g2, b2, a2 := b.RGBA()

	diff += int64(math.Abs(float64(r1 - r2)))
	diff += int64(math.Abs(float64(g1 - g2)))
	diff += int64(math.Abs(float64(b1 - b2)))
	diff += int64(math.Abs(float64(a1 - a2)))
	return diff
}

func boundsMatch(a, b image.Rectangle) bool {
	return a.Min.X == b.Min.X && a.Min.Y == b.Min.Y && a.Max.X == b.Max.X && a.Max.Y == b.Max.Y
}

// SmsTest TODO: NEEDS COMMENT INFO
func SmsTest(t *testing.T) {
	if os.Getenv("CIRCLECI") != "true" {
		function := func(t *testing.T) {
			ajax := infra.Ajax{
				// 		// dani
				// 		// Path:    infra.TestServer.URL + "/auth/sms/51991640012",
				// kerpen
				// Path:    infra.TestServer.URL + "/auth/sms/51985956935",
				Path:    infra.TestServer.URL + "/auth/sms/51984457005",
				Headers: map[string]string{"Content-Type": "application/json"},
			}
			if !true {
				response, body := infra.NewTestRequest(t, ajax)
				assert.Equal(t, response.StatusCode, 200)
				log.Println(string(body))
			}
		}
		t.Run("SmsTest", function)
	}

}

// RefreshToken is a test that uploads an image
func RefreshToken(t *testing.T, helper *types.CustomerLoginResponse) types.JWTRefresh {
	var tokenResponse types.JWTRefresh
	t.Run("Send SMS", func(t *testing.T) {
		token := types.JWTRefresh{Token: helper.Token}
		bolB, _ := json.Marshal(token)
		ajax := infra.Ajax{
			Method:  "POST",
			Path:    infra.TestServer.URL + "/auth",
			Body:    bytes.NewBuffer(bolB),
			Headers: map[string]string{"Content-Type": "application/json"},
		}
		response, body := infra.NewTestRequest(t, ajax)
		assert.Equal(t, response.StatusCode, 200)
		err := json.Unmarshal(body, &tokenResponse)
		assert.NoError(t, err)
		// assert.NotEqual(t, helper.Token, tokenResponse.Token)
		log.Println(tokenResponse.Token)
	})
	return tokenResponse
}

// SoftLogin is a test that uploads an image
func SoftLogin(t *testing.T, helper *types.CustomerLoginResponse) types.CustomerLoginResponse {
	var softCustomer types.CustomerLoginResponse
	t.Run("Login Soft", func(t *testing.T) {
		user := types.SoftLoginRequest{Email: "otestedoteste@gmail.com", Password: "abc123"}
		bolB, _ := json.Marshal(user)
		ajax := infra.Ajax{
			Method: "POST",
			Path:   "https://api.onnictrlmusic.com/app/v4/customer/login",
			Body:   bytes.NewBuffer(bolB),
			Headers: map[string]string{
				"X-AUTH-APPLICATION-TOKEN": "mYX5a43As?V7LGhTbtJ_KHpE4;:xGl;P=QvM0iJd2oPH5V<FIgB[hy67>u_3@[pc",
				"Content-Type":             "application/json",
			},
		}
		_, body := infra.NewTestRequest(t, ajax)
		err := json.Unmarshal(body, &softCustomer)
		assert.NoError(t, err)
		log.Println(softCustomer)
	})
	return softCustomer
}
