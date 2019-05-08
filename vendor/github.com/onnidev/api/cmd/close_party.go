package cmd

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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var closePartyCMD = cobra.Command{
	Use:              "closeparty",
	Short:            "Return the current version of the API",
	TraverseChildren: true,
	Long:             `Return the current version of the API.`,
	Run: func(cmd *cobra.Command, args []string) {
		infra.Mongo(viper.GetString("db"))
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
		if len(partiesTobeClosed) == 0 {
			log.Println("No parties to close")
			return
		}
		var CLOSED = 0
		for _, party := range partiesTobeClosed {
			log.Printf("Attempting to close %s\n", party.Name)
			closedParty, err := partyCollection.CloseParty(party.ID.Hex())
			if err != nil {
				log.Println(err.Error())
				return
			}
			j, err := json.MarshalIndent(closedParty, "", "    ")
			if err != nil {
				log.Println(err.Error())
				return
			}
			log.Println(string(j))
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
				go common.MailPartyReport(admin, party, buf)
			}
			CLOSED = CLOSED + 1
		}
		log.Printf("Closed %v parties\n", CLOSED)
	},
}
