package cmd

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/spf13/cobra"
)

// VERSION is the app version
var VERSION = "UNVERSIONED"

var versionCMD = cobra.Command{
	Use:   "version",
	Short: "Return the current version of the API",
	Long:  `Return the current version of the API.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(VERSION)
	},
}

func version(r chi.Router) {
	r.Get("/version", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(VERSION))
	})
}
