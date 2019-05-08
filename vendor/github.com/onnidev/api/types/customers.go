package types

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2/bson"
)

// Customer type for the above middleware
type Customer struct {
	ID             bson.ObjectId `json:"id" bson:"_id,omitempty"`
	CreationDate   *Timestamp    `bson:"creationDate,omitempty" json:"creationDate,omitempty"`
	UpdateDate     *Timestamp    `bson:"updateDate,omitempty" json:"updateDate,omitempty"`
	FirstName      string        `bson:"firstName" json:"firstName"`
	LastName       string        `bson:"lastName" json:"lastName"`
	Mail           string        `json:"mail" bson:"mail"`
	Password       string        `json:"password" bson:"password"`
	Phone          string        `json:"phone" bson:"phone"`
	UserName       string        `json:"username" bson:"username"`
	DocumentNumber *string       `json:"documentNumber" bson:"documentNumber"`
	Trusted        *string       `json:"trusted" bson:"trusted"`
	BirthDate      *Timestamp    `bson:"birthDate,omitempty" json:"birthDate,omitempty"`
	// MusicStyles    []Style         `bson:"musicStyles" json:"musicStyles"`
	FacebookID    string          `bson:"facebookId" json:"facebookId"`
	FavoriteClubs []bson.ObjectId `json:"favoriteClubs" bson:"favoriteClubs"`
}

// CustomerPostRequest type for the above middleware
type CustomerPostRequest struct {
	ID             string     `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName      string     `bson:"firstName" json:"firstName"`
	LastName       string     `bson:"lastName" json:"lastName"`
	Mail           string     `json:"mail" bson:"mail"`
	Password       string     `json:"password" bson:"password"`
	Phone          string     `json:"phone" bson:"phone"`
	UserName       string     `json:"username" bson:"username"`
	DocumentNumber string     `json:"documentNumber" bson:"documentNumber"`
	BirthDate      *Timestamp `bson:"birthDate,omitempty" json:"birthDate,omitempty"`
	FacebookID     string     `bson:"facebookId" json:"facebookId"`
	FavoriteClubs  []string   `json:"favoriteClubs" bson:"favoriteClubs"`
}

// Name sdfkjn
func (customer Customer) Name() string {
	return customer.FirstName + " " + customer.LastName
}

// Ready sdfkjn
func (customer Customer) Ready() bool {
	if customer.DocumentNumber == nil {
		return false
	}
	if customer.DocumentNumber != nil {
		if *customer.DocumentNumber == "" {
			return false
		}
	}
	if customer.UserName == "" {
		return false
	}
	if customer.Phone == "" {
		return false
	}
	return true
}

// Reset ksjdfn
type Reset struct {
	Mail string `json:"mail" bson:"mail"`
}

// CustomerCheck ksjdfn
type CustomerCheck struct {
	Type    string `json:"type" bson:"type"`
	Payload string `json:"payload" bson:"payload"`
}

// CustomerCheckResponse ksjdfn
type CustomerCheckResponse struct {
	Result bool `json:"result" bson:"result"`
}

// CustomerQuery ksjdfn
type CustomerQuery struct {
	Query string `json:"query" bson:"query"`
}

// GenerateToken sdfkjn
func (customer Customer) GenerateToken() (string, error) {
	return customer.createtoken(time.Hour * 1000000)
}

// CustomerClaims whatever
type CustomerClaims struct {
	CustomerID string `json:"customerId,omitempty"`
	jwt.StandardClaims
}

func (customer Customer) createtoken(d time.Duration) (string, error) {
	mySigningKey := []byte(viper.GetString("jwtsecret"))
	claims := CustomerClaims{
		customer.ID.Hex(),
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(d).Unix(),
			Issuer:    "softdesign",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySigningKey)
}

// CustomerPatch sdfkjd
type CustomerPatch struct {
	FirstName           string     `bson:"firstName,omitempty" json:"firstName,omitempty"`
	LastName            string     `bson:"lastName,omitempty" json:"lastName,omitempty"`
	Mail                string     `json:"mail,omitempty" bson:"mail,omitempty"`
	Password            string     `json:"password,omitempty" bson:"password,omitempty"`
	Phone               string     `json:"phone,omitempty" bson:"phone,omitempty"`
	UserName            string     `json:"username,omitempty" bson:"username,omitempty"`
	DocumentNumber      string     `json:"documentNumber,omitempty" bson:"documentNumber,omitempty"`
	BirthDate           *Timestamp `bson:"birthDate,omitempty" json:"birthDate,omitempty"`
	AddFavoriteClubs    []string   `json:"addFavoriteClubs,omitempty" bson:"addFavoriteClubs,omitempty"`
	RemoveFavoriteClubs []string   `json:"removeFavoriteClubs,omitempty" bson:"removeFavoriteClubs,omitempty"`
}

// LogInCustomer sdkjfnsdf
func (customer Customer) LogInCustomer() (CustomerLoginResponse, error) {
	var response CustomerLoginResponse
	token, err := customer.GenerateToken()
	if err != nil {
		return response, err
	}
	return CustomerLoginResponse{
		ID:    customer.ID.Hex(),
		Mail:  customer.Mail,
		Name:  customer.FirstName + " " + customer.LastName,
		Token: token,
	}, nil
}
