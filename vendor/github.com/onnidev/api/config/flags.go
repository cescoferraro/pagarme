package config

import (
	"github.com/cescoferraro/tools/venom"
	"github.com/spf13/viper"
)

// RunServerFlags makes it a 12-factor app
var RunServerFlags = venom.CommandFlag{
	venom.Flag{
		Name:        "PAGARME",
		Description: "Pagarme API Key",
		Safe:        true,
		Value:       pagarmeToken()},
	venom.Flag{
		Name:        "GOOGLEMAPSTOKEN",
		Description: "Google Maps",
		Safe:        true,
		Value:       "AIzaSyCe802OoEmmQtumF9WUxLRvipyfQz_pEs0"},
	venom.Flag{
		Name:        "verbose",
		Short:       "v",
		Description: "A descriptio about this cool flag",
		Value:       false},
	venom.Flag{
		Name:        "db",
		Short:       "d",
		Description: "The database to connect to",
		Value:       "onni"},
	venom.Flag{
		Name:        "loopport",
		Description: "A descriptio about this cool flag",
		Value:       9000},
	venom.Flag{
		Name:        "oplogport",
		Description: "A descriptio about this cool flag",
		Value:       8000},
	venom.Flag{
		Name:        "port",
		Short:       "p",
		Description: "A descriptio about this cool flag",
		Value:       7000},
	venom.Flag{
		Name:        "redishost",
		Description: "A descriptio about this cool flag",
		Value:       "localhost"},
	venom.Flag{
		Name:        "redisport",
		Description: "A descriptio about this cool flag",
		Value:       6379},
	venom.Flag{
		Name:        "mongohost",
		Description: "A descriptio about this cool flag",
		Value:       "localhost"},
	venom.Flag{
		Name:        "mongoport",
		Description: "A descriptio about this cool flag",
		Value:       27017},
	venom.Flag{
		Name:        "mongouser",
		Description: "A descriptio about this cool flag",
		Value:       ""},
	venom.Flag{
		Name:        "mongopass",
		Description: "A descriptio about this cool flag",
		Value:       ""},
	venom.Flag{
		Name:        "env",
		Description: "A validation flag",
		Value:       "dev"},
	venom.Flag{
		Name:        "X-AUTH-APPLICATION-TOKEN",
		Safe:        true,
		Description: "A descriptio about this cool flag",
		Value:       "mYX5a43As?V7LGhTbtJ_KHpE4;:xGl;P=QvM0iJd2oPH5V<FIgB[hy67>u_3@[pc"},
	venom.Flag{
		Name:        "sendinblue",
		Safe:        true,
		Description: "SendInBlue Token",
		Value:       "sh1RqULOF7mGzIE0"},
	venom.Flag{
		Name:        "jwtsecret",
		Description: "JWT Secret",
		Safe:        true,
		Value:       "IlIC91CpujjOf29rhXww39OKOz33ddW7c9ZeE0eBCHmanu2tKHdpFfQC1J6ykBt",
	},
	venom.Flag{
		Name:        "FBSTAFFAPPTOKEN",
		Description: "Facebook App Token",
		Safe:        true,
		Value:       facebookStaffAppToken(),
	},
	venom.Flag{
		Name:        "FBAPPTOKEN",
		Description: "Facebook App Token",
		Safe:        true,
		Value:       FacebookAppToken(),
	},
}

func pagarmeToken() string {
	if viper.GetString("env") == "homolog" {
		return "ak_test_63dAXI6XQ1duXFltMbhr2PqiVHup4O"
	}
	return "ak_live_iSZM4oGkTcBmVhzGysL9BE2QP6ZAIz"
}

// FacebookAppToken TODO: NEEDS COMMENT INFO
func FacebookAppToken() string {
	if viper.GetString("env") == "homolog" {
		return "450941118572249|2DCuGcC0vXcTp8TWvXDkMJ34e28"
	}
	return "450924951907199|kPCEUcDflsXYHkviq6QXSmZDuKE"
}

func facebookStaffAppToken() string {
	return "727259887478727|4zLeI8WceR4twDWIGVKLYqSBH04"
}
