// Package main ONNi API.
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
//
// This should demonstrate all routes on https://api.onni.live
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: https
//     Host: api.onni.live
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Francesco Ferraro<francescoaferraro@gmail.com> https://www.cescoferraro.xyz
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//       JWT_TOKEN
//
//     SecurityDefinitions:
//     JWT_TOKEN:
//          type: apiKey
//          name: JWT_TOKEN
//          in: header
//
// swagger:meta
//go:generate swagger generate spec -o swagger.json
//go:generate go-bindata -pkg shared -o shared/file.go templates/... public/... report/templates voucher/templates/... userClub/templates/... promotionalCustomer/templates/...
package main

import (
	"log"

	"github.com/onnidev/api/cmd"
	"github.com/onnidev/api/config"
)

func main() {
	command := cmd.FullServer()
	config.RunServerFlags.Register(command)
	if err := command.Execute(); err != nil {
		log.Fatal(err)
	}
}
