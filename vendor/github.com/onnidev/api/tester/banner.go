package tester

import (
	"bytes"
	"encoding/json"
	"log"
	"testing"

	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/types"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
)

// CreateClubBanner is a test that uploads an image
func CreateClubBanner(t *testing.T, helper *types.UserClubLoginResponse, fileID, clubID, itype string) *types.Banner {
	var banner types.Banner
	function := func(t *testing.T) {
		bannerRequest := types.BannerPostRequest{
			ID:          bson.NewObjectId().Hex(),
			Name:        "COOLTURE NA ÏNDIA",
			Type:        itype,
			Description: "a really cool description",
			Action:      clubID,
			Image:       fileID,
		}
		j, err := json.MarshalIndent(bannerRequest, "", "    ")
		assert.NoError(t, err)
		ajax := infra.Ajax{
			Method:  "POST",
			Path:    infra.TestServer.URL + "/banner",
			Headers: helper.Headers(),
			Body:    bytes.NewBuffer(j),
		}
		_, body := infra.NewTestRequest(t, ajax)
		log.Println(string(body))
		err = json.Unmarshal(body, &banner)
		assert.NoError(t, err)
		log.Println(banner)
	}
	t.Run("Create banner", function)
	return &banner
}
