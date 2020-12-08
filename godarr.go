package main

import (
	"fmt"

	"github.com/kr/pretty"
)

type SonarrClient struct {
	api *ApiClient
}

type RadarrClient struct {
	api *ApiClient
}

type LidarrClient struct {
	api *ApiClient
}

type MylarrClient struct {
	api *ApiClient
}

type GodarrClient struct {
	Sonarr *SonarrClient
	Radarr *RadarrClient
	Lidarr *LidarrClient
	Mylarr *MylarrClient
}

func NewClient() (*GodarrClient, error) {
	return &GodarrClient{
		Sonarr: &SonarrClient{},
		Radarr: &RadarrClient{},
		Lidarr: &LidarrClient{},
		Mylarr: &MylarrClient{},
	}, nil
}

func (gd *GodarrClient) SonarrSetup(addr, akey string) error {
	gd.Sonarr.api = &ApiClient{
		address: addr,
		apiKey:  akey,
	}

	err := gd.Sonarr.api.connSetup()
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

func main() {
	client, _ := NewClient()
	client.SonarrSetup("http://10.0.0.65:8989", "68e63fe4a24f41879f62dc39a217309b")
	tst, err := client.Sonarr.SystemStatus()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%# v\n", pretty.Formatter(tst))
}
