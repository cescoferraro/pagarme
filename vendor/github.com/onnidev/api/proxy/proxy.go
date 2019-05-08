package proxy

import (
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/onnidev/api/shared"
	"github.com/pressly/chi/render"
	"github.com/spf13/viper"
)

// Proxy TODO: NEEDS COMMENT INFO
func Proxy(w http.ResponseWriter, r *http.Request) {
	rest := strings.Replace(r.URL.RequestURI(), "/proxy/", "", -1)
	// url := "https://backend.onni.live/" + rest
	url := "https://api.onnictrlmusic.com/" + rest
	if viper.GetString("env") == "homolog" {
		url = "http://portal.softdesign-rs.com.br:9107/" + rest

	}
	req, err := http.NewRequest(r.Method, url, r.Body)
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		DisableCompression: true,
	}
	// req.Header = r.Header
	for key, value := range r.Header {
		newKey := key
		if shared.Contains([]string{"Jwt_token", "X-Auth-Application-Token"}, key) {
			newKey = strings.ToUpper(key)
		}
		if !shared.Contains([]string{"Accept-Encoding"}, key) {
			req.Header.Set(newKey, value[0])
		}
	}
	log.Println(r.Header["Jwt_token"])
	log.Println(r.Header["X-Auth-Application-Token"])
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		render.Status(r, http.StatusUnauthorized)
		render.JSON(w, r, err.Error())
		return

	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
