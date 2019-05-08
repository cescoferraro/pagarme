package recipient_test

import (
	"log"
	"math/rand"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/cescoferraro/tools/venom"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/onnidev/api/config"
	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/recipient"
	"github.com/onnidev/api/tester"
	"github.com/onnidev/api/types"
	"github.com/onnidev/api/userClub"
	"github.com/spf13/cobra"
)

// TestSpec isk kjsdf
func TestSpecReturn(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	config.RunServerFlags.Register(&cobra.Command{})
	venom.PrintViperConfig(config.RunServerFlags)
	infra.Mongo("onni")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	recipient.Routes(r)
	userClub.Routes(r)
	infra.TestServer = httptest.NewServer(r)
	defer infra.TestServer.Close()
	loginHelper := tester.LogInCescoUserClub(t)
	retrievedRecipients := tester.ListRecipients(t, loginHelper)
	id := retrievedRecipients[rand.Intn(len(retrievedRecipients))].ID.Hex()
	readRecipients := tester.ReadRecipient(t, loginHelper, id)
	log.Println(readRecipients)
	// id := retrievedRecipients[rand.Intn(len(retrievedRecipients))].ID.Hex()
	// readRecipients := tester.ReadRecipient(t, loginHelper, id)
	tester.GetRecipientTransactionBalance(t, loginHelper, readRecipients.RecipientID)
	tester.GetRecipientTransactionBalance(t, loginHelper, "re_cj0ea9v3s0065sj5y19n1xwt3")
	from := types.Timestamp(time.Now().Add(-2 * 24 * time.Hour))
	till := types.Timestamp(time.Now())
	tester.GetRecipientTransactionDays(t, loginHelper,
		types.FinanceQuery{
			From:        &from,
			Till:        &till,
			RecipientID: readRecipients.RecipientID})
	tester.GetRecipientTransactionDays(t, loginHelper,
		types.FinanceQuery{
			From:        &from,
			Till:        &till,
			RecipientID: "re_cj0ea9v3s0065sj5y19n1xwt3",
		})
	tester.GetRecipientTransactionDaysBalance(t, loginHelper,
		types.FinanceQuery{
			From:        &from,
			Till:        &till,
			RecipientID: readRecipients.RecipientID})
	tester.GetRecipientTransactionDaysBalance(t, loginHelper,
		types.FinanceQuery{
			From:        &from,
			Till:        &till,
			RecipientID: "re_cj0ea9v3s0065sj5y19n1xwt3",
		})
	tester.GetRecipientOperationDaysBalance(t, loginHelper,
		types.FinanceQuery{
			From:        &from,
			Till:        &till,
			RecipientID: readRecipients.RecipientID})
	tester.GetRecipientOperationDaysBalance(t, loginHelper,
		types.FinanceQuery{
			From:        &from,
			Till:        &till,
			RecipientID: "re_cj0ea9v3s0065sj5y19n1xwt3",
		})
	tester.GetRecipientOperationDaysXlsx(t, loginHelper,
		types.FinanceQuery{
			From:        &from,
			Till:        &till,
			RecipientID: "re_cj0ea9v3s0065sj5y19n1xwt3",
		})
	tester.GetRecipientAntecipations(t, loginHelper,
		types.FinanceQuery{
			RecipientID: "re_cj0ea9v3s0065sj5y19n1xwt3",
		})
	payment := types.Timestamp(time.Now().Add(1 * 24 * time.Hour))
	tester.GetRecipientAntecipationsLimit(t, loginHelper,
		types.FinanceQuery{
			RecipientID: "re_cj0ea9v3s0065sj5y19n1xwt3",
			From:        &payment,
		})
}
