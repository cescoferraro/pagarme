package tester

import (
	"bytes"
	"encoding/json"
	"log"
	"testing"

	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/types"
	"github.com/stretchr/testify/assert"
	"github.com/xuri/excelize"
)

// GetRecipientTransactionDaysBalance skdjnf
func GetRecipientTransactionDays(t *testing.T,
	helper *types.UserClubLoginResponse, query types.FinanceQuery) {
	t.Run("Get recipient day balance", func(t *testing.T) {
		bolB, _ := json.Marshal(query)
		ajax := infra.Ajax{
			Method:  "POST",
			Body:    bytes.NewBuffer(bolB),
			Path:    infra.TestServer.URL + "/recipient/transactions",
			Headers: helper.Headers()}
		_, body := infra.NewTestRequest(t, ajax)
		log.Println(string(body))
	})
}

// GetRecipientAntecipations skdjnf
func GetRecipientAntecipationsLimit(t *testing.T,
	helper *types.UserClubLoginResponse, query types.FinanceQuery) {
	t.Run("Get recipient antecipations limit", func(t *testing.T) {
		bolB, _ := json.Marshal(query)
		ajax := infra.Ajax{
			Method:  "POST",
			Body:    bytes.NewBuffer(bolB),
			Path:    infra.TestServer.URL + "/recipient/antecipation/limits",
			Headers: helper.Headers()}
		_, body := infra.NewTestRequest(t, ajax)
		log.Println(string(body))
	})
}

// GetRecipientAntecipations skdjnf
func GetRecipientAntecipations(t *testing.T,
	helper *types.UserClubLoginResponse, query types.FinanceQuery) {
	t.Run("Get recipient antecipations", func(t *testing.T) {
		bolB, _ := json.Marshal(query)
		ajax := infra.Ajax{
			Method:  "POST",
			Body:    bytes.NewBuffer(bolB),
			Path:    infra.TestServer.URL + "/recipient/antecipations",
			Headers: helper.Headers()}
		_, body := infra.NewTestRequest(t, ajax)
		log.Println(string(body))
	})
}

// GetRecipientOperationDaysXlsx skdjnf
func GetRecipientOperationDaysXlsx(t *testing.T,
	helper *types.UserClubLoginResponse, query types.FinanceQuery) {
	t.Run("Get recipient day balance", func(t *testing.T) {
		bolB, _ := json.Marshal(query)
		ajax := infra.Ajax{
			Method:  "POST",
			Body:    bytes.NewBuffer(bolB),
			Path:    infra.TestServer.URL + "/recipient/operations/xlsx",
			Headers: helper.Headers()}
		_, body := infra.NewTestRequest(t, ajax)
		_, err := excelize.OpenReader(bytes.NewReader(body))
		assert.NoError(t, err)
	})
}

// GetRecipientOperationDaysBalance skdjnf
func GetRecipientOperationDaysBalance(t *testing.T,
	helper *types.UserClubLoginResponse, query types.FinanceQuery) {
	t.Run("Get recipient day balance", func(t *testing.T) {
		bolB, _ := json.Marshal(query)
		ajax := infra.Ajax{
			Method:  "POST",
			Body:    bytes.NewBuffer(bolB),
			Path:    infra.TestServer.URL + "/recipient/operations",
			Headers: helper.Headers()}
		_, _ = infra.NewTestRequest(t, ajax)
	})
}

// GetRecipientTransactionDaysBalance skdjnf
func GetRecipientTransactionDaysBalance(t *testing.T,
	helper *types.UserClubLoginResponse, query types.FinanceQuery) {
	t.Run("Get recipient day balance", func(t *testing.T) {
		bolB, _ := json.Marshal(query)
		ajax := infra.Ajax{
			Method:  "POST",
			Body:    bytes.NewBuffer(bolB),
			Path:    infra.TestServer.URL + "/recipient/operations/timeline",
			Headers: helper.Headers()}
		_, body := infra.NewTestRequest(t, ajax)
		log.Println(string(body))
	})
}

// GetRecipientTransactionBalance skdjnf
func GetRecipientTransactionBalance(t *testing.T,
	helper *types.UserClubLoginResponse, recipientID string) types.PagarMeTransactionBalance {
	var balance types.PagarMeTransactionBalance
	query := types.FinanceQuery{RecipientID: recipientID}
	t.Run("Get recipient balance", func(t *testing.T) {
		bolB, _ := json.Marshal(query)
		ajax := infra.Ajax{
			Method:  "POST",
			Body:    bytes.NewBuffer(bolB),
			Path:    infra.TestServer.URL + "/recipient/balance",
			Headers: helper.Headers()}
		_, body := infra.NewTestRequest(t, ajax)
		err := json.Unmarshal(body, &balance)
		log.Println(string(body))
		assert.NoError(t, err)
		log.Println(balance)
	})
	return balance
}
