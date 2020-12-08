package sonarr

type DiskSpace struct {
	Path       string `json:"path"`
	Label      string `json:"label"`
	FreeSpace  int64  `json:"freeSpace"`
	TotalSpace int64  `json:"totalSpace"`
}

func (sc *SonarrClient) DiskSpace() ([]*DiskSpace, error) {
	retv := []*DiskSpace{}
	return retv, nil
}
