package sonarr

type RootFolder struct {
	Path            string        `json:"path"`
	FreeSpace       int64         `json:"freeSpace"`
	UnmappedFolders []interface{} `json:"unmappedFolders"`
	ID              int           `json:"id"`
}

func (sc *SonarrClient) RootFolders() ([]*RootFolder, error) {
	retv := []*RootFolder{}
	return retv, nil
}
