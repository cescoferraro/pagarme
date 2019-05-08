package customer_test

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"net/http/httptest"
	"testing"

	"github.com/cescoferraro/tools/venom"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/icrowley/fake"
	"github.com/onnidev/api/club"
	"github.com/onnidev/api/config"
	"github.com/onnidev/api/customer"
	"github.com/onnidev/api/file"
	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/tester"
	"github.com/onnidev/api/types"
	"github.com/onnidev/api/userClub"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

// TestSpec isk kjsdf
func TestSpecReturn(t *testing.T) {
	config.RunServerFlags.Register(&cobra.Command{})
	venom.PrintViperConfig(config.RunServerFlags)
	viper.SetDefault("verbose", true)
	infra.Mongo("onni")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	userClub.Routes(r)
	club.Routes(r)
	file.Routes(r)
	customer.Routes(r)
	infra.TestServer = httptest.NewServer(r)
	defer infra.TestServer.Close()
	cescoUserClub := tester.LogInCescoUserClub(t)
	cescoCustomer := tester.LogInCescoCustomer(t)
	clubs := tester.ListAllClubsByToken(t, cescoUserClub.Token)
	_ = tester.RandomClubID(clubs)
	customer := tester.ReadCustomer(t, cescoCustomer, cescoCustomer.ID)

	rand.Seed(time.Now().UTC().UnixNano())
	log.Println(customer)
	newName := fake.FirstName()
	log.Println(cescoCustomer)
	// newUsername := fake.FirstName()
	// remove := customer.FavoriteClubs[rand.Intn(len(customer.FavoriteClubs))].Hex()
	// add := clubs[rand.Intn(len(clubs))].ID.Hex()
	rand.Seed(time.Now().UTC().UnixNano())
	// add2 := clubs[rand.Intn(len(clubs))].ID.Hex()
	birthDate := types.Timestamp(time.Now().Add(24 * 365 * time.Duration(random(25, 50)) * time.Hour))
	newDoc := fake.SimplePassword()
	// for add == remove {
	// 	rand.Seed(time.Now().UTC().UnixNano())
	// 	add = clubs[rand.Intn(len(clubs))].ID.Hex()
	// }
	patch := types.CustomerPatch{
		FirstName: newName,
		// Mail:                fake.EmailAddress(),
		DocumentNumber: newDoc,
		Phone:          fake.Phone(),
		BirthDate:      &birthDate,
		// AddFavoriteClubs: []string{add},
		// RemoveFavoriteClubs: []string{remove},
	}
	p, _ := json.MarshalIndent(patch, "", "    ")
	t.Run("Patch a customer", func(t *testing.T) {
		ajax := infra.Ajax{
			Method:  "PATCH",
			Path:    infra.TestServer.URL + "/customer",
			Body:    bytes.NewBuffer(p),
			Headers: cescoCustomer.Headers()}
		_, body := infra.NewTestRequest(t, ajax)

		var patchedCustomer types.Customer
		err := json.Unmarshal(body, &patchedCustomer)
		assert.NoError(t, err)
		assert.Equal(t, patch.FirstName, patchedCustomer.FirstName)
		// assert.Equal(t, patch.DocumentNumber, *patchedCustomer.DocumentNumber)
		assert.Equal(t, patch.Phone, patchedCustomer.Phone)
		// assert.Equal(t, patch.Mail, patchedCustomer.Mail)
		// assert.Equal(t, patch.UserName, patchedCustomer.UserName)
		y1, m1, d1 := patchedCustomer.BirthDate.Time().Date()
		y2, m2, d2 := birthDate.Time().Date()
		assert.Equal(t, y1, y2)
		assert.Equal(t, m1, m2)
		assert.Equal(t, d1, d2)
		filters, _ := json.MarshalIndent(patchedCustomer, "", "    ")
		log.Println(string(filters))
		// var found bool
		// var fount bool
		// for _, fav := range patchedCustomer.FavoriteClubs {
		// if fav.Hex() == add {
		// 	found = true
		// }
		// if fav.Hex() == remove {
		// 	fount = true
		// }
		// }
		// assert.True(t, found)
		// assert.False(t, fount)
		// log.Println(found, fount)
	})
}
