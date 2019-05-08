package tester

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/types"
	"github.com/stretchr/testify/assert"
)

// ReadRecipient skdjnf
func ReadRecipient(t *testing.T, helper *types.UserClubLoginResponse, id string) types.Recipient {
	var readRecipient types.Recipient
	t.Run("Read Recipient", func(t *testing.T) {
		ajax := infra.Ajax{
			Path:    infra.TestServer.URL + "/recipient/" + id,
			Headers: helper.Headers()}
		_, body := infra.NewTestRequest(t, ajax)
		log.Println(string(body))
		json.Unmarshal(body, &readRecipient)
	})
	return readRecipient
}

// ListRecipients skdjnf
func ListRecipients(t *testing.T, helper *types.UserClubLoginResponse) []types.Recipient {
	var retrievedRecipients []types.Recipient
	headers := map[string]string{
		"JWT_TOKEN":    helper.Token,
		"Content-Type": "application/json"}
	t.Run("List Recipients", func(t *testing.T) {
		ajax := infra.Ajax{
			Method:  "GET",
			Path:    infra.TestServer.URL + "/recipient",
			Headers: headers}
		_, body := infra.NewTestRequest(t, ajax)
		err := json.Unmarshal(body, &retrievedRecipients)
		assert.NoError(t, err)
	})
	return retrievedRecipients
}
