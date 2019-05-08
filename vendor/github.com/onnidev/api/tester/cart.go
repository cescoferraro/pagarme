package tester

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/types"
	"github.com/stretchr/testify/assert"
)

// GetCart is a test that uploads an image
func GetCart(t *testing.T, helper *types.CustomerLoginResponse) types.Cart {
	var all types.Cart
	cooltureID := "5b0c6a16cc922d565cd57ac9"
	function := func(t *testing.T) {
		ajax := infra.Ajax{
			Method:  "GET",
			Path:    infra.TestServer.URL + "/cart/" + cooltureID,
			Headers: helper.Headers(),
		}
		_, body := infra.NewTestRequest(t, ajax)
		err := json.Unmarshal(body, &all)
		assert.NoError(t, err)
		for _, promotion := range all.Promotions {
			log.Println("received partyProduct _id ", promotion.ID.Hex())
		}
		assert.NoError(t, err)
	}
	t.Run("list cart", function)
	return all
}
