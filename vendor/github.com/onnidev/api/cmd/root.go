package cmd

import (
	"log"
	"net/http"
	"strconv"

	"github.com/cescoferraro/tools/venom"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/onnidev/api/antitheft"
	"github.com/onnidev/api/appclub"
	"github.com/onnidev/api/auth"
	"github.com/onnidev/api/banner"
	"github.com/onnidev/api/bans"
	"github.com/onnidev/api/buy"
	"github.com/onnidev/api/card"
	"github.com/onnidev/api/cart"
	"github.com/onnidev/api/club"
	"github.com/onnidev/api/clubLead"
	"github.com/onnidev/api/config"
	"github.com/onnidev/api/customer"
	"github.com/onnidev/api/deeplink"
	"github.com/onnidev/api/file"
	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/invitedCustomer"
	"github.com/onnidev/api/location"
	"github.com/onnidev/api/menuProduct"
	"github.com/onnidev/api/menuTicket"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/musicStyles"
	"github.com/onnidev/api/notification"
	"github.com/onnidev/api/party"
	"github.com/onnidev/api/partyProduct"
	"github.com/onnidev/api/playground"
	"github.com/onnidev/api/product"
	"github.com/onnidev/api/promotion"
	"github.com/onnidev/api/promotionalCustomer"
	"github.com/onnidev/api/proxy"
	"github.com/onnidev/api/pushRegistry"
	"github.com/onnidev/api/recipient"
	"github.com/onnidev/api/report"
	"github.com/onnidev/api/shared"
	"github.com/onnidev/api/site"
	"github.com/onnidev/api/token"
	"github.com/onnidev/api/userClub"
	"github.com/onnidev/api/voucher"
	"github.com/onnidev/api/ws"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// FullServer sdkjfn
func FullServer() *cobra.Command {
	cmd := &cobra.Command{
		Short:            "A brief description of your command",
		Long:             `A loooooooonger description of your command.`,
		TraverseChildren: true,
		Run: func(cmd *cobra.Command, args []string) {
			venom.PrintViperConfig(config.RunServerFlags)
			infra.Mongo(viper.GetString("db"))
			done := make(chan bool)
			go ws.Redis(shared.RedisPrefixer("staff"), done)
			one := make(chan bool)
			go ws.Redis(shared.RedisPrefixer("dashboard"), one)
			ne := make(chan bool)
			go ws.Redis(shared.RedisPrefixer("app"), ne)
			r := chi.NewRouter()
			r.Use(middlewares.Cors)
			r.Use(middleware.Logger)
			token.TokenRoutes(r)
			customer.Routes(r)
			card.Routes(r)
			notification.Routes(r)
			club.Routes(r)
			party.Routes(r)
			voucher.Routes(r)
			report.Routes(r)
			product.Routes(r)
			proxy.Routes(r)
			userClub.Routes(r)
			partyProduct.Routes(r)
			deeplink.Routes(r)
			file.Routes(r)
			invitedCustomer.Routes(r)
			auth.Routes(r)
			musicStyles.Routes(r)
			site.Routes(r)
			recipient.Routes(r)
			ws.Routes(r)
			cart.Routes(r)
			antitheft.Routes(r)
			clubLead.Routes(r)
			menuProduct.Routes(r)
			menuTicket.Routes(r)
			promotion.Routes(r)
			promotionalCustomer.Routes(r)
			pushRegistry.Routes(r)
			playground.Routes(r)
			buy.Routes(r)
			banner.Routes(r)
			appclub.Routes(r)
			bans.Routes(r)
			location.Routes(r)
			version(r)
			config.Docs(r)
			log.Println("Listening on 0.0.0.0:" + strconv.Itoa(viper.GetInt("port")))
			log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(viper.GetInt("port")), r))
		},
	}
	cmd.AddCommand(&loopCMD)
	cmd.AddCommand(&oplogCMD)
	cmd.AddCommand(&versionCMD)
	cmd.AddCommand(&closePartyCMD)

	return cmd
}
