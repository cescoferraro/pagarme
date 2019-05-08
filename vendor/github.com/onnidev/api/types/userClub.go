package types

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2/bson"
)

// UserClub sdifu
type UserClub struct {
	ID             bson.ObjectId `json:"id" bson:"_id,omitempty"`
	CreationDate   *Timestamp    `bson:"creationDate" json:"creationDate"`
	UpdateDate     *Timestamp    `bson:"updateDate,omitempty" json:"updateDate,omitempty"`
	Name           string        `bson:"name" json:"name"`
	Mail           string        `json:"mail" bson:"mail"`
	Password       string        `bson:"password" json:"password"`
	Status         string        `bson:"status" json:"status"`
	Profile        string        `bson:"profile" json:"profile"`
	RegisterOrigin string        `bson:"registerOrigin" json:"registerOrigin"`

	Favorite *bson.ObjectId  `json:"favorite" bson:"favorite,omitempty"`
	Clubs    []bson.ObjectId `json:"clubs" bson:"clubs,omitempty"`
	Tags     *[]string       `json:"tags" bson:"tags"`
}

// SoftUserClubLoginResponse sdfkjn
type SoftUserClubLoginResponse struct {
	ID                bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Token             string        `json:"token" bson:"token"`
	Name              string        `json:"name" bson:"name"`
	Mail              string        `json:"mail" bson:"mail"`
	Profile           string        `json:"profile" bson:"profile"`
	ClubID            string        `json:"clubId" bson:"clubId"`
	ClubPercentTicket float64       `json:"clubPercentTicket" bson:"clubPercentTicket"`
	ClubPercentDrink  float64       `json:"clubPercentDrink" bson:"clubPercentDrink"`

	Tags  *[]string `json:"tags" bson:"tags"`
	Clubs []string  `json:"clubs" bson:"clubs"`
}

// UserClubLoginResponse sdfkjn
type UserClubLoginResponse struct {
	ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Token   string        `json:"token" bson:"token"`
	Name    string        `json:"name" bson:"name"`
	Mail    string        `json:"mail" bson:"mail"`
	Profile string        `json:"profile" bson:"profile"`
	ClubID  string        `json:"clubId" bson:"clubId"`

	Clubs []Club `json:"clubs" bson:"clubs"`
}

// Headers sdfkjn
func (user UserClubLoginResponse) Headers() map[string]string {
	return map[string]string{
		"JWT_TOKEN":    user.Token,
		"Content-Type": "application/json"}
}

// CustomerLoginResponse sdfkjn
type CustomerLoginResponse struct {
	ID    string `json:"id" bson:"id"`
	Token string `json:"token" bson:"token"`
	Name  string `json:"name" bson:"name"`
	Mail  string `json:"mail" bson:"mail"`
}

// SoftLoginRequest sjkdfn
type SoftLoginRequest struct {
	Email    string `json:"mail,omitempty" bson:"mail,omitempty"`
	Password string `json:"password" bson:"mail,omitempty"`
}

// LoginRequest sjkdfn
type LoginRequest struct {
	Email    string `json:"email,omitempty" bson:"email,omitempty"`
	UserName string `json:"username,omitempty" bson:"username,omitempty"`
	Password string `json:"password"`
}

// Oauth2LoginRequest sjkdfn
type Oauth2LoginRequest struct {
	AccessToken string `json:"access_token" bson:"access_token"`
}

// ChangePasswordRequest sdkfjn
type ChangePasswordRequest struct {
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

// GenerateToken sdfkjn
func (user UserClub) GenerateToken() (string, error) {
	return user.createtoken(time.Minute * 77777 * 5)
}

func (user UserClub) createtoken(d time.Duration) (string, error) {
	mySigningKey := []byte(viper.GetString("jwtsecret"))
	claims := MyClaimsType{
		user.ID.Hex(),
		user.Profile,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(d).Unix(),
			Issuer:    "softdesign",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySigningKey)
	return tokenString, err
}

// MyClaimsType whatever
type MyClaimsType struct {
	ClientID string `json:"clientId,omitempty"`
	Profile  string `json:"profile"`
	jwt.StandardClaims
}

// SwaggerLoginResponse dsfkjnsdf
// swagger:response loginResponse
type SwaggerLoginResponse struct {
	// in: body
	Body UserClubLoginResponse `json:"body"`
}

// Endpoints type
type Endpoints struct {
	Tokens    chi.Router
	Customers chi.Router
	Cards     chi.Router
	Parties   chi.Router
}

// ClubUserResponse dkjfnsd
type ClubUserResponse struct {
	ID                string `json:"id"`
	Token             string `json:"token"`
	Name              string `json:"name"`
	Profile           string `json:"profile"`
	ClubName          string `json:"clubName"`
	ClubID            string `json:"clubId"`
	ClubPercentTicket int    `json:"clubPercentTicket"`
	Mail              string `json:"mail"`
}

// UserClubPatch Request type for the above middleware
type UserClubPatch struct {
	Name    string   `bson:"name,omitempty" json:"name,omitempty"`
	Mail    string   `bson:"mail,omitempty" json:"mail,omitempty"`
	Profile string   `bson:"profile,omitempty" json:"profile,omitempty"`
	Status  string   `bson:"status,omitempty" json:"status,omitempty"`
	Clubs   []string `json:"clubs,omitempty" bson:"clubs,omitempty"`
}

// UserClubPostRequest type for the above middleware
type UserClubPostRequest struct {
	ID      string    `json:"id" bson:"id"`
	Name    string    `bson:"name" json:"name"`
	Email   string    `bson:"mail" json:"mail"`
	Profile string    `bson:"profile" json:"profile"`
	Tags    *[]string `bson:"tags,omitempty" json:"tags,omitempty"`
	Clubs   []string  `json:"clubs" bson:"clubs"`
}

// SoftHeaders sdfkjn
func (user CustomerLoginResponse) SoftHeaders() map[string]string {
	token := "8IO6q/yI14z;NnD0?H9|S60$n3M'#LWC;L2y|O47**,t&foARU]8fW14M2R^~C8"
	if viper.GetString("env") == "homolog" {
		token = "mYX5a43As?V7LGhTbtJ_KHpE4;:xGl;P=QvM0iJd2oPH5V<FIgB[hy67>u_3@[pc"
	}
	return map[string]string{
		"X-AUTH-APPLICATION-TOKEN": token,
		"JWT_TOKEN":                user.Token,
		"Content-Type":             "application/json"}
}

// Headers sdfkjn
func (user CustomerLoginResponse) Headers() map[string]string {
	return map[string]string{
		"JWT_TOKEN":    user.Token,
		"Content-Type": "application/json"}
}
