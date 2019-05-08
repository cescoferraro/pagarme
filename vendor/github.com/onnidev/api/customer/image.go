package customer

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/shared"
)

// Photo is the shit
func Photo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "customerId")
	customersRepo := r.Context().Value(middlewares.CustomersRepoKey).(interfaces.CustomersRepo)
	customer, err := customersRepo.GetByID(id)
	if err != nil {
		shared.MakeONNiError(w, r, 400, err)
		return
	}
	log.Println(customer.FacebookID)
	if customer.FacebookID == "" {
		avatart := []string{
			"https://user-images.githubusercontent.com/36003926/39221162-d154a218-480b-11e8-90c2-490cb5993985.png",
			"https://user-images.githubusercontent.com/36003926/39221163-d196e24a-480b-11e8-8023-ef9280f19532.png",
			"https://user-images.githubusercontent.com/36003926/39221164-d1b78540-480b-11e8-9e1c-470a6ac62a88.png",
			"https://user-images.githubusercontent.com/36003926/39221165-d1dc7012-480b-11e8-84a6-54725f047ec3.png",
			"https://user-images.githubusercontent.com/36003926/39221166-d201eeaa-480b-11e8-9ea0-516eb27e1b7a.png",
			"https://user-images.githubusercontent.com/36003926/39221167-d2258216-480b-11e8-9195-794f76ded47e.png",
			"https://user-images.githubusercontent.com/36003926/39221168-d2474af4-480b-11e8-9e29-85aa83d90ea0.png",
			"https://user-images.githubusercontent.com/36003926/39221169-d266633a-480b-11e8-8bf1-c466a3cc3cf7.png",
			"https://user-images.githubusercontent.com/36003926/39221170-d28827d6-480b-11e8-8eed-577225b2b67e.png",
			"https://user-images.githubusercontent.com/36003926/39221171-d2a8bd98-480b-11e8-9784-190a475be8e1.png",
			"https://user-images.githubusercontent.com/36003926/39221172-d2cc74e0-480b-11e8-9c8a-d509ea848f3f.png",
			"https://user-images.githubusercontent.com/36003926/39221173-d30c08ee-480b-11e8-9614-c91eac5ab589.png",
			"https://user-images.githubusercontent.com/36003926/39221174-d32d29c0-480b-11e8-9e62-a7de2f4ad7c7.png",
			"https://user-images.githubusercontent.com/36003926/39221175-d352dea4-480b-11e8-8768-6e6dd1c23c07.png",
			"https://user-images.githubusercontent.com/36003926/39221176-d380b676-480b-11e8-81b3-11b450bddfbd.jpg",
			"https://user-images.githubusercontent.com/36003926/39221177-d3a6dd4c-480b-11e8-956b-c894e7c6c01e.png",
			"https://user-images.githubusercontent.com/36003926/39221427-418ac30e-480d-11e8-8a52-ab4cf187a50a.png",
			"https://user-images.githubusercontent.com/36003926/39221178-d3cb008c-480b-11e8-83f5-33b43d0ed770.png",
			"https://user-images.githubusercontent.com/36003926/39221428-41b06a5a-480d-11e8-9840-5f8d7fd39467.png",
			"https://user-images.githubusercontent.com/36003926/39221429-41d1b0e8-480d-11e8-804c-79ab15091a35.png",
			"https://user-images.githubusercontent.com/36003926/39221430-41f197aa-480d-11e8-9d69-6ef95984e1a0.png",
			"https://user-images.githubusercontent.com/36003926/39221431-4213ecc4-480d-11e8-8c97-654b97c9d05c.png",
			"https://user-images.githubusercontent.com/36003926/39221432-42370196-480d-11e8-83ce-31b61db14dc5.png",
			"https://user-images.githubusercontent.com/36003926/39221433-42788ddc-480d-11e8-9d5f-a7d0ce614a03.png",
			"https://user-images.githubusercontent.com/36003926/39221434-429cb9b4-480d-11e8-8695-c69de8e6abee.png",
			"https://user-images.githubusercontent.com/36003926/39221435-42bdcd7a-480d-11e8-881a-70ecdba18ad4.png",
			"https://user-images.githubusercontent.com/36003926/39221436-42de836c-480d-11e8-9358-91cd649c372e.png",
			"https://user-images.githubusercontent.com/36003926/39221437-42ffe228-480d-11e8-84db-2da509ea3d57.png",
			"https://user-images.githubusercontent.com/36003926/39221438-43230f14-480d-11e8-96b1-3d30aac7bb26.png",
			"https://user-images.githubusercontent.com/36003926/39221439-43442500-480d-11e8-8ace-37c8653ff31f.png",
			"https://user-images.githubusercontent.com/36003926/39221440-436417d4-480d-11e8-9213-3c9c397b4e2b.png",
		}
		rand.Seed(time.Now().UTC().UnixNano())
		photo := avatart[rand.Intn(len(avatart))]
		http.Redirect(w, r, photo, 301)
		return
	}
	uri := `https://graph.facebook.com/` + customer.FacebookID + `/picture?type=square&width=300`
	log.Println(uri)
	http.Redirect(w, r, uri, 301)
}
