package getnet

import (
	"net/http"
	"net/url"
	"time"
)

// API TODO: NEEDS COMMENT INFO
type API struct {
	Key     string
	Verbose bool
}

// New creates ane
func New(key string) *API {
	return &API{Key: key, Verbose: true}
}

func (api *API) getURL() (*url.URL, error) {
	var URL *url.URL
	URL, err := url.Parse("https://api-homologacao.getnet.com.br/")
	if err != nil {
		return URL, err
	}
	return URL, nil
}

func (api *API) defaultHTTPClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 400,
	}
}
