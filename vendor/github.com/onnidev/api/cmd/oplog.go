package cmd

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/cescoferraro/tools/venom"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/onnidev/api/config"
	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/middlewares"
	"github.com/onnidev/api/oplog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var oplogCMD = cobra.Command{
	Use:   "oploger",
	Short: "Return the current version of the API",
	Long:  `Return the current version of the API.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(VERSION)
		venom.PrintViperConfig(config.RunServerFlags)
		infra.Mongo(viper.GetString("db"))
		r := chi.NewRouter()
		r.Use(middlewares.Cors)
		r.Use(middleware.Logger)
		go oplog.Start()
		oplog.Routes(r)
		log.Println("Listening on 0.0.0.0:" + strconv.Itoa(viper.GetInt("oplogport")))
		log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(viper.GetInt("oplogport")), r))
	},
}
