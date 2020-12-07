package godarr

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
