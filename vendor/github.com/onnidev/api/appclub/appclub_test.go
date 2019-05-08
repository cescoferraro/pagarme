package appclub_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"

	"gopkg.in/mgo.v2/bson"

	"golang.org/x/net/context/ctxhttp"

	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/interfaces"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	// . "github.com/smartystreets/goconvey/convey"
)

// TestSpec isk kjsdf
func TestSpec(t *testing.T) {
	if os.Getenv("CIRCLECI") != "true" {
		t.Run("LEITOR READ REDIRECT", func(t *testing.T) {
			viper.SetDefault("env", "dev")
			infra.MongoDev("onni")
			db, err := infra.Cloner()
			assert.NoError(t, err)
			log.Println("before token collection")
			repo, err := interfaces.NewTokenCollection(db)
			assert.NoError(t, err)
			userclubID := bson.ObjectIdHex("5a578349cc922d56e5ebe45d")
			token, err := repo.Create(userclubID)
			assert.NoError(t, err)
			URL, err := url.Parse("https://api.onnictrlmusic.com")
			// URL, err := url.Parse("http://localhost:9000")
			assert.NoError(t, err)
			URL.Path += "/appclub/v1/voucher/read"
			fmt.Println(URL.String())
			j, err := json.MarshalIndent(struct {
				ClubID     string `json:"clubId"`
				VoucherID  string `json:"voucherId"`
				UserClubID string `json:"userClubId"`
			}{
				ClubID:     "592ca3b3cc922d18bdae4cd8",
				VoucherID:  "5b7335979fb84f0001c83d51",
				UserClubID: userclubID.Hex(),
			}, "", "    ")
			assert.NoError(t, err)
			log.Println("Request sent to Soft Leitor")
			log.Println(string(j))
			req, err := http.NewRequest("POST", URL.String(), bytes.NewBuffer(j))
			assert.NoError(t, err)
			req.Header["X-AUTH-TOKEN"] = []string{token.Token}
			req.Header["X-CLIENT-ID"] = []string{userclubID.Hex()}
			req.Header["X-AUTH-APPLICATION-TOKEN"] = []string{"mYX5a43As?V7LGhTbtJ_KHpE4;:xGl;P=QvM0iJd2oPH5V<FIgB[hy67>u_3@[pc"}
			j, err = json.MarshalIndent(req.Header, "", "    ")
			assert.NoError(t, err)
			log.Println("Request sent to Soft Leitor")
			log.Println(string(j))
			req.Header.Set("Content-Type", "application/json")
			client := http.Client{Timeout: time.Second * 20}
			resp, err := ctxhttp.Do(context.Background(), &client, req)
			assert.NoError(t, err)
			defer resp.Body.Close()
			byt, err := ioutil.ReadAll(resp.Body)
			assert.NoError(t, err)
			log.Println(string(byt))
			assert.NoError(t, err)
			log.Println(resp.StatusCode)
			assert.NoError(t, nil)
		})
	}
}
