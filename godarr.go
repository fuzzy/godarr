package client

import (
	"github.com/fuzzy/godarr/client"
	"github.com/fuzzy/godarr/sonarr"
)

type RadarrClient struct {
	api *client.ApiClient
}

type LidarrClient struct {
	api *client.ApiClient
}

type MylarrClient struct {
	api *client.ApiClient
}

type GodarrClient struct {
	Sonarr *sonarr.SonarrClient
	Radarr *RadarrClient
	Lidarr *LidarrClient
	Mylarr *MylarrClient
}

func NewClient() (*GodarrClient, error) {
	return &GodarrClient{
		Sonarr: &sonarr.SonarrClient{},
		Radarr: &RadarrClient{},
		Lidarr: &LidarrClient{},
		Mylarr: &MylarrClient{},
	}, nil
}

func (gd *GodarrClient) SonarrSetup(addr, akey string) error {
	gd.Sonarr.Api = &client.ApiClient{
		RawAddress: addr,
		ApiKey:     akey,
	}

	err := gd.Sonarr.Api.ConnSetup()
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
