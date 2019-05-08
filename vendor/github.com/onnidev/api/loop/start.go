package loop

import (
	"bytes"
	"encoding/json"
	"log"
	"time"

	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/interfaces"
	"github.com/onnidev/api/report"
	"github.com/onnidev/api/report/common"
	"github.com/onnidev/api/types"
)

// COUNTER sdkjfn
var COUNTER = 0

// Start sdkjf
func Start() {
	routine(time.Now())
	doEvery(60000*time.Millisecond, routine)
}

var routine = func(t time.Time) {
	log.Println("Starting the routine")
	db, err := infra.Cloner()
	partyCollection, err := interfaces.NewPartiesCollection(db)
	if err != nil {
		log.Println(err.Error())
		return
	}
	parties, err := partyCollection.List()
	if err != nil {
		log.Println(err.Error())
		return
	}
	var partiesTobeClosed []types.Party
	for _, party := range parties {
		if party.Status == "ACTIVE" {
			if time.Now().After(party.EndDate.Time()) {
				partiesTobeClosed = append(partiesTobeClosed, party)
			}
		}
	}
	var CLOSED = 0
	for _, party := range partiesTobeClosed {
		log.Printf("Attempting to close %s\n", party.Name)
		closedParty, err := partyCollection.CloseParty(party.ID.Hex())
		if err != nil {
			log.Println(err.Error())
			return
		}

		voucherCollection, err := interfaces.NewVoucherCollection(db)
		if err != nil {
			log.Println(err.Error())
			return
		}
		all, err := voucherCollection.GetByParty(party.ID.Hex())
		if err != nil {
			log.Println(err.Error())
			return
		}
		clubCollection, err := interfaces.NewClubsCollection(db)
		if err != nil {
			log.Println(err.Error())
			return
		}
		club, err := clubCollection.GetByID(party.ClubID.Hex())
		if err != nil {
			log.Println(err.Error())
			return
		}
		userClubCollection, err := interfaces.NewUserClubCollection(db)
		if err != nil {
			log.Println(err.Error())
			return
		}
		buf := new(bytes.Buffer)
		data := common.Data{Club: club, Party: party, Vouchers: types.VouchersList(all)}
		xlsx := report.Excel(data)
		xlsx.Write(buf)
		admins, err := userClubCollection.FindClubAdmins(party.ClubID.Hex())
		if err != nil {
			log.Println(err.Error())
			return
		}
		for _, admin := range admins {
			log.Println("hello")
			go common.MailPartyReport(admin, party, buf)
		}
		CLOSED = CLOSED + 1
		j, err := json.MarshalIndent(closedParty, "", "    ")
		if err != nil {
			log.Println(err.Error())
			return
		}
		log.Println(string(j))
	}
	log.Printf("closed %v\n", CLOSED)
	COUNTER = COUNTER + 1
	log.Println("Finishing the routine")
}

func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}
