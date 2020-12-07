package godarr

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
)

type ApiClient struct {
	address string
	apiKey  string
	Address *url.URL
	Client  *http.Client
}

func (api *ApiClient) connSetup() error {
	if api.address == "" {
		return errors.New("No address specified")
	} else if api.apiKey == "" {
		return errors.New("No API Key specified")
	}

	addressUrl, err := url.Parse(api.address)
	if err != nil {
		return err
	}

	if !strings.HasSuffix(addressUrl.Path, "/") {
		addressUrl.Path += "/"
	}
	if !strings.HasSuffix(addressUrl.Path, "api/") {
		addressUrl.Path += "api/"
	}
	api.address = addressUrl.String()
	api.Address = addressUrl

	api.Client = http.DefaultClient
	return nil
}

type GodarrClient struct {
	Sonarr *ApiClient
	Radarr *ApiClient
	Lidarr *ApiClient
	Mylarr *ApiClient
}

func NewClient() (*GodarrClient, error) {
	return &GodarrClient{
		Sonarr: &ApiClient{},
		Radarr: &ApiClient{},
		Lidarr: &ApiClient{},
		Mylarr: &ApiClient{},
	}, nil
}

func (gd *GodarrClient) SonarrSetup(addr, akey string) error {
	gd.Sonarr = &ApiClient{
		address: addr,
		apiKey:  akey,
	}

	err := gd.Sonarr.connSetup()
	if err != nil {
		return err
	}

	return nil
}

func (gd *GodarrClient) RadarrSetup(addr, akey string) error {
	return nil
}

func (gd *GodarrClient) LidarrSetup(addr, akey string) error {
	return nil
}

func (gd *GodarrClient) MylarrSetup(addr, akey string) error {
	return nil
}
