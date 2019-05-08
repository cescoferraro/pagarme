package onni

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"image"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/JKhawaja/sendinblue"
	"github.com/carlogit/phash"
	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/types"
	"github.com/spf13/viper"
)

// AllKindsConstrain TODO: NEEDS COMMENT INFO
var AllKindsConstrain = types.VoucherUseConstrain{
	Drink:  true,
	Ticket: true,
}

// AntiTheft TODO: NEEDS COMMENT INFO
func AntiTheft(ctx context.Context, id string) (types.AntiTheftResult, error) {
	repo, ok := ctx.Value(middlewares.AntiTheftRepoKey).(interfaces.AntiTheftRepo)
	if !ok {
		return types.AntiTheftResult{}, nil
	}
	result, err := repo.GetByID(id)
	if err != nil {
		return types.AntiTheftResult{}, nil
	}

	return result, nil
}

// PendingAntithefts TODO: NEEDS COMMENT INFO
func PendingAntithefts(ctx context.Context) ([]types.AntiTheftResult, error) {
	repo, ok := ctx.Value(middlewares.AntiTheftRepoKey).(interfaces.AntiTheftRepo)
	if !ok {
		return []types.AntiTheftResult{}, nil
	}
	result, err := repo.PendingAntithefts()
	if err != nil {
		return []types.AntiTheftResult{}, nil
	}

	return result, nil
}

// CreateAntiTheftRecordBan TODO: NEEDS COMMENT INFO
func CreateAntiTheftRecordBan(ctx context.Context, result types.AntiTheftResult, bans *[]types.Ban) error {
	result.Bans = bans
	return CreateAntiTheftRecord(ctx, result)
}

// CreateAntiTheftRecordInvoiceVoucherCreation TODO: NEEDS COMMENT INFO
func CreateAntiTheftRecordInvoiceVoucherCreation(ctx context.Context, result types.AntiTheftResult, invoices types.Invoices, vouchers []types.Voucher) error {
	result.Invoices = &invoices
	array := []bson.ObjectId{}
	for _, voucher := range vouchers {
		log.Print("*" + voucher.ID.Hex()[:3])
		array = append(array, voucher.ID)
	}
	log.Print("\n")
	result.Vouchers = &array
	return CreateAntiTheftRecord(ctx, result)
}

// CreateAntiTheftRecordAfterPagarME TODO: NEEDS COMMENT INFO
func CreateAntiTheftRecordAfterPagarME(
	ctx context.Context,
	result types.AntiTheftResult,
	invoices types.Invoices,
	vouchers []types.Voucher,
	pg types.PagarMeTransactionResponse,
) error {
	result.Invoices = &invoices
	array := []bson.ObjectId{}
	for _, voucher := range vouchers {
		log.Print("*" + voucher.ID.Hex()[:3])
		array = append(array, voucher.ID)
	}
	result.Vouchers = &array
	result.PG = &pg
	return CreateAntiTheftRecord(ctx, result)
}

// CreateAntiTheftRecord TODO: NEEDS COMMENT INFO
func CreateAntiTheftRecord(ctx context.Context, result types.AntiTheftResult) error {
	log.Println("attempting inserting anthitheft model")
	db, err := infra.Cloner()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	repo, err := interfaces.NewAntiTheftCollection(db)
	defer db.Session.Close()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	log.Println("inserinrg'registro")
	log.Println("before inserting anthitheft model")
	err = repo.Collection.Insert(result)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	if result.Score > 4 {
		SendAntiTheftMail(ctx, result)
	}
	return nil
}

// AntiTheftRoutine TODO: NEEDS COMMENT INFO
func AntiTheftRoutine(ctx context.Context, club types.Club, party types.Party, customer types.Customer, products types.BuyPostList, result types.AntiTheftResult) error {
	if customer.Trusted != nil {
		if *customer.Trusted == "ACTIVE" {
			return nil
		}
	}
	if customer.CreationDate.Time().Before(time.Now().AddDate(0, -3, 0)) {
		return nil
	}
	if result.Score > 7 {
		log.Println(" 4444 fraudewwww ")
		CreateAntiTheftRecord(ctx, result)
		return errors.New("caiu no antifraud")
	}
	return nil
}

// SendAntiTheftMail TODO: NEEDS COMMENT INFO
func SendAntiTheftMail(ctx context.Context, result types.AntiTheftResult) {
	sibClient, err := sib.NewClient(viper.GetString("sendinblue"))
	if err != nil {
		log.Println(err)
		return
	}
	byt, err := json.MarshalIndent(result, "     ", "")
	if err != nil {
		return
	}
	myTemplate := &sib.Template{
		Template_name: "AntiTheft",
		Html_content:  string(byt),
		Subject:       "AntiTheft " + result.CustomerName,
		From_email:    "onni@onni.live",
		Status:        1}
	createResponse, err := sibClient.CreateTemplate(myTemplate)
	if err != nil {
		log.Println(err)
		return
	}
	options := sib.NewEmailOptions("", "", []string{""}, []string{""})
	userList := []string{"francescoaferraro@gmail.com", "onni.kerpen@gmail.com", "guilhermelongonilara@gmail.com", "scomazzond@gmail.com", "thiagobinscidade@gmail.com"}
	_, err = sibClient.SendTemplateEmail(createResponse.Data.ID, userList, options)
	if err != nil {
		log.Println(err)
		return
	}

}

// CalculateAntiFraud TODO: NEEDS COMMENT INFO
func CalculateAntiFraud(ctx context.Context, club types.Club, party types.Party, customer types.Customer, products types.BuyPostList) (types.AntiTheftResult, error) {
	totalTicket := products.SumTickets(party, club)
	totalDrink := products.SumDrinks(party, club)
	model := types.AntiTheftModelONNi
	scoreCard, err := AntiTheftScoreCard(ctx, club, party, customer, totalTicket, totalDrink, model)
	if err != nil {
		return types.AntiTheftResult{}, err
	}
	initialScore := scoreCard
	cards, err := CustomerCards(ctx, customer.ID.Hex())
	if err != nil {
		return types.AntiTheftResult{}, err
	}
	vouchers, err := PartyCustomerVoucher(ctx, party, customer)
	if err != nil {
		return types.AntiTheftResult{}, err
	}
	byt, err := ImageBytesFromFB(customer.FacebookID)
	if err != nil {
		return types.AntiTheftResult{}, err
	}
	scorePicture := 0
	log.Println("before image docode")
	_, _, err = image.Decode(bytes.NewReader(byt))
	if err != nil {
		scorePicture = scorePicture + 3
	}
	log.Println("afte image docode")
	if err == nil {
		fresult, err := IsFacebookImage(byt)
		if err != nil {
			return types.AntiTheftResult{}, err
		}
		if fresult {
			scorePicture = types.AntiTheftModelONNi.Picture
		}
	}
	initialScore = initialScore + scorePicture
	tickets := TicketScore(party, club, vouchers, totalTicket, model)
	drinks := DrinkScore(party, club, vouchers, totalDrink, model)
	horario := types.Timestamp(time.Now())
	result := types.AntiTheftResult{
		ID:                 bson.NewObjectId(),
		PartyID:            party.ID,
		ClubID:             club.ID,
		CreationDate:       &horario,
		CustomerID:         customer.ID,
		Cards:              len(cards),
		CustomerDate:       customer.CreationDate.Time().String(),
		CustomerName:       customer.Name(),
		CustomerMail:       customer.Mail,
		Total:              products.Sum(party, club),
		ScoreCard:          scoreCard,
		ScoreDrinks:        drinks,
		ScorePicture:       scorePicture,
		ScoreTickets:       tickets,
		Reviewed:           false,
		Score:              AntiTheftArithmetics(totalTicket, totalDrink, tickets, drinks, initialScore),
		AccoutableDrinks:   types.VouchersList(vouchers).Accountable().FilterDrinks().Sum(),
		AccoutableTickets:  types.VouchersList(vouchers).Accountable().FilterTickets().Sum(),
		ClubTicketsAverage: ClubTicketsAverage(club),
		ClubDrinksAverage:  ClubDrinksAverage(club),
	}
	return result, nil
}

// AntiTheftScoreCard TODO: NEEDS COMMENT INFO
func AntiTheftScoreCard(ctx context.Context, club types.Club, party types.Party, customer types.Customer, sumtickets, sumdrinks float64, model types.AntiTheftModel) (int, error) {
	score := 0
	cards, err := CustomerCards(ctx, customer.ID.Hex())
	if err != nil {
		return score, err
	}
	switch count := len(cards); {
	case count >= model.CardL3:
		score = score + 8
	case count == model.CardL2:
		score = score + 4
	case count == model.CardL1:
		score = score + 2
	default:
	}
	return score, nil
}

// AntiTheftArithmetics TODO: NEEDS COMMENT INFO
func AntiTheftArithmetics(sumtickets, sumdrinks float64, ticketScore, drinkScore, initialScore int) int {
	if sumtickets > float64(0) && sumdrinks > float64(0) {
		initialScore = initialScore + ticketScore + drinkScore
		log.Println("#### initialScore", initialScore)
		return initialScore
	}
	if sumtickets > float64(0) {
		initialScore = initialScore + ticketScore
		return initialScore
	}
	if sumdrinks > float64(0) {
		initialScore = initialScore + drinkScore
		return initialScore
	}
	return initialScore
}

// PartyCustomerVoucher TODO: NEEDS COMMENT INFO
func PartyCustomerVoucher(ctx context.Context, party types.Party, customer types.Customer) ([]types.CompleteVoucher, error) {
	vouchersCollection, ok := ctx.Value(middlewares.VouchersRepoKey).(interfaces.VouchersRepo)
	if !ok {
		err := errors.New("bug assert")
		return []types.CompleteVoucher{}, err
	}
	accountables, err := vouchersCollection.GetCompleteVouchersByPartyAndCustomer(party.ID.Hex(), customer.ID.Hex())
	if err != nil {
		return accountables, err
	}
	return accountables, nil
}

// TicketScore TODO: NEEDS COMMENT INFO
func TicketScore(party types.Party, club types.Club, vouchers []types.CompleteVoucher, sumtickets float64, model types.AntiTheftModel) int {
	all := types.VouchersList(vouchers)
	score := 0
	clubTicketAvg := ClubTicketsAverage(club)
	switch sum := sumtickets + all.Accountable().FilterTickets().Sum(); {
	case model.TicketL2*clubTicketAvg > sum && sum > model.TicketL1*clubTicketAvg:
		score = score + 2
	case model.TicketL3*clubTicketAvg > sum && sum >= model.TicketL2*clubTicketAvg:
		score = score + 3
	case model.TicketL4*clubTicketAvg > sum && sum >= model.TicketL3*clubTicketAvg:
		score = score + 4
	case model.TicketL5*clubTicketAvg > sum && sum >= model.TicketL4*clubTicketAvg:
		score = score + 6
	case sum >= model.TicketL5*clubTicketAvg:
		score = score + 8
	default:
	}
	return score
}

// DrinkScore TODO: NEEDS COMMENT INFO
func DrinkScore(party types.Party, club types.Club, vouchers []types.CompleteVoucher, sumdrinks float64, model types.AntiTheftModel) int {
	all := types.VouchersList(vouchers)
	scoreD := 0
	clubDrinkAvg := ClubDrinksAverage(club)
	switch sum := sumdrinks + all.Accountable().FilterDrinks().Sum(); {
	case model.DrinkL2*clubDrinkAvg > sum && sum > model.DrinkL1*clubDrinkAvg:
		scoreD = scoreD + 2
	case model.DrinkL3*clubDrinkAvg > sum && sum >= model.DrinkL2*clubDrinkAvg:
		scoreD = scoreD + 3
	case model.DrinkL4*clubDrinkAvg > sum && sum >= model.DrinkL3*clubDrinkAvg:
		scoreD = scoreD + 4
	case model.DrinkL5*clubDrinkAvg > sum && sum >= model.DrinkL4*clubDrinkAvg:
		scoreD = scoreD + 6
	case sum >= model.DrinkL5*clubDrinkAvg:
		scoreD = scoreD + 8
	default:
	}
	return scoreD
}

// AntiTheftScoreVouchers TODO: NEEDS COMMENT INFO
func AntiTheftScoreVouchers(ctx context.Context, club types.Club, party types.Party, customer types.Customer, sumtickets, sumdrinks float64, model types.AntiTheftModel) (int, int, error) {
	vouchers, err := PartyCustomerVoucher(ctx, party, customer)
	if err != nil {
		return 0, 0, err
	}
	return TicketScore(party, club, vouchers, sumtickets, model), DrinkScore(party, club, vouchers, sumdrinks, model), nil
}

// ClubDrinksAverage TODO: NEEDS COMMENT INFO
func ClubDrinksAverage(club types.Club) float64 {
	limit := float64(100)
	if club.AverageExpendituresProduct != nil {
		avg := *club.AverageExpendituresProduct
		if avg > limit {
			return avg
		}
	}
	return limit
}

// ClubTicketsAverage TODO: NEEDS COMMENT INFO
func ClubTicketsAverage(club types.Club) float64 {
	limit := float64(130)
	if club.AverageExpendituresTicket != nil {
		avg := *club.AverageExpendituresTicket
		if avg > limit {
			return avg
		}
	}
	return limit
}

// IsFacebookImage TODO: NEEDS COMMENT INFO
func IsFacebookImage(byt []byte) (bool, error) {
	customerHash, err := phash.GetHash(bytes.NewReader(byt))
	if err != nil {
		return false, err
	}
	aahash := "1010100011111111011111111111111111010100111111110010111111111111"
	bbhash := "1011111100011111110001111110000011010000000000010110000111110000"
	distanceF := phash.GetDistance(customerHash, aahash)
	distanceM := phash.GetDistance(customerHash, bbhash)
	log.Println(" distance", distanceF, distanceM)
	if distanceF == 0 || distanceM == 0 {
		return true, nil
	}
	return false, nil
}

// IsCustomerFake TODO: NEEDS COMMENT INFO
func IsCustomerFake(customer types.Customer) (bool, error) {
	byt, err := ImageBytesFromFB(customer.FacebookID)
	if err != nil {
		return false, err
	}
	_, _, err = image.Decode(bytes.NewReader(byt))
	if err != nil {
		return false, nil
	}
	fresult, err := IsFacebookImage(byt)
	if err != nil {
		return false, err
	}
	return fresult, nil

}

// ImageBytesFromFB TODO: NEEDS COMMENT INFO
func ImageBytesFromFB(fbid string) ([]byte, error) {
	uri := `https://graph.facebook.com/` + fbid + `/picture?type=square&width=300`
	response, e := http.Get(uri)
	if e != nil {
		return []byte{}, e
	}
	defer response.Body.Close()
	byt, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return byt, err
	}
	return byt, nil
}
