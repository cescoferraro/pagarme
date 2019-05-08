package tester

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"testing"
	"time"

	"github.com/icrowley/fake"
	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/types"
	"github.com/stretchr/testify/assert"
)

// ReadClubLead TODO: NEEDS COMMENT INFO
func ReadClubLead(t *testing.T, id string) types.ClubLead {
	var lead types.ClubLead
	t.Run("read ClubLead", func(t *testing.T) {
		ajax := infra.Ajax{
			Method: "GET",
			Path:   infra.TestServer.URL + "/clubLead/" + id}
		_, body := infra.NewTestRequest(t, ajax)
		err := json.Unmarshal(body, &lead)
		assert.NoError(t, err)
	})
	return lead
}

// ListAllClubLeads TODO: NEEDS COMMENT INFO
func ListAllClubLeads(t *testing.T) []types.ClubLead {
	var leads []types.ClubLead
	t.Run("List Parties", func(t *testing.T) {
		ajax := infra.Ajax{
			Method: "GET",
			Path:   infra.TestServer.URL + "/clubLead"}
		_, body := infra.NewTestRequest(t, ajax)
		err := json.Unmarshal(body, &leads)
		assert.NoError(t, err)
	})
	return leads
}

// RandomClubLeadID TODO: NEEDS COMMENT INFO
func RandomClubLeadID(leads []types.ClubLead) string {
	if len(leads) != 0 {
		rand.Seed(time.Now().UTC().UnixNano())
		return leads[rand.Intn(len(leads))].ID.Hex()
	}
	return ""
}

// CreateClubLeadRequest TODO: NEEDS COMMENT INFO
func CreateClubLeadRequest(t *testing.T) types.ClubLeadPostRequest {
	name := fake.FirstName()
	image := UploadImage(t)
	image2 := UploadImage(t)
	return types.ClubLeadPostRequest{
		AdminName:       name,
		AdminMail:       name + "@gmail.com",
		AdminPhone:      fake.Phone(),
		Image:           image.FileID.Hex(),
		BackgroundImage: image2.FileID.Hex(),
	}
}

// CreateClubLead sdkjfdn
func CreateClubLead(t *testing.T, lead types.ClubLeadPostRequest) types.ClubLead {
	j, err := json.Marshal(lead)
	assert.NoError(t, err)
	var createdLead types.ClubLead
	t.Run("read ClubLead", func(t *testing.T) {
		ajax := infra.Ajax{
			Method: "POST",
			Path:   infra.TestServer.URL + "/clubLead",
			Body:   bytes.NewBuffer(j),
		}
		_, body := infra.NewTestRequest(t, ajax)
		err := json.Unmarshal(body, &createdLead)
		assert.NoError(t, err)
		assert.Equal(t, lead.AdminName, createdLead.AdminName)
		assert.Equal(t, lead.AdminPhone, createdLead.AdminPhone)
		assert.Equal(t, lead.AdminMail, createdLead.AdminMail)
	})
	return createdLead
}
