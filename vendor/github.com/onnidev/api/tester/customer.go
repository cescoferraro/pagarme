package tester

import (
	"bytes"
	"encoding/json"
	"log"
	"testing"

	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/types"
	"github.com/stretchr/testify/assert"
)

// LogInCescoCustomer sdfjkn
func LogInCescoCustomer(t *testing.T) *types.CustomerLoginResponse {
	return CustomerLoginTestingHelperFactory(t, CescoCustomerLogin)
}

// LogInKerpenCustomer sdfjkn
func LogInKerpenCustomer(t *testing.T) *types.CustomerLoginResponse {
	return CustomerLoginTestingHelperFactory(t, KerpenCustomerLogin)
}

// CescoCustomerLogin sdkjfn
var CescoCustomerLogin = types.LoginRequest{
	Email:    "francescoaferraro@gmail.com",
	Password: "786962",
}

// KerpenCustomerLogin sdkjfn
var KerpenCustomerLogin = types.LoginRequest{
	Email:    "matheuskerpen@gmail.com",
	Password: "373245",
}

// CustomerLoginTestingHelperFactory TODO: NEEDS COMMENT INFO
func CustomerLoginTestingHelperFactory(t *testing.T, user types.LoginRequest) *types.CustomerLoginResponse {
	var User types.CustomerLoginResponse
	function := func(t *testing.T) {
		bolB, _ := json.Marshal(user)
		ajax := infra.Ajax{
			Method: "POST",
			Path:   infra.TestServer.URL + "/customer/login",
			Body:   bytes.NewBuffer(bolB),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		_, body := infra.NewTestRequest(t, ajax)
		json.Unmarshal(body, &User)
		log.Println(User.ID)
	}
	t.Run("Login to the existing customer API", function)
	return &User
}

// ReadCustomer dkjns
func ReadCustomer(t *testing.T, helper *types.CustomerLoginResponse, id string) types.Customer {
	var customer types.Customer
	t.Run("Get Customer's Info", func(t *testing.T) {
		ajax := infra.Ajax{
			Method:  "GET",
			Path:    infra.TestServer.URL + "/customer/" + id,
			Headers: helper.Headers()}
		_, body := infra.NewTestRequest(t, ajax)
		err := json.Unmarshal(body, &customer)
		assert.NoError(t, err)
	})
	return customer
}

// CustomerCheck TODO: NEEDS COMMENT INFO
func CustomerCheck(t *testing.T, helper *types.CustomerLoginResponse, check types.CustomerCheck) types.CustomerCheckResponse {
	var response types.CustomerCheckResponse
	function := func(t *testing.T) {
		p, err := json.Marshal(check)
		assert.NoError(t, err)
		ajax := infra.Ajax{
			Method:  "POST",
			Path:    infra.TestServer.URL + "/customer/check",
			Body:    bytes.NewBuffer(p),
			Headers: helper.Headers(),
		}
		_, body := infra.NewTestRequest(t, ajax)
		json.Unmarshal(body, &response)
	}
	t.Run("Check Customer key/value", function)
	return response
}
