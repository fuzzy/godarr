package sonarr

import "time"

type Command struct {
	Name                string    `json:"name"`
	StartedOn           time.Time `json:"startedOn"`
	StateChangeTime     time.Time `json:"stateChangeTime"`
	SendUpdatesToClient bool      `json:"sendUpdatesToClient"`
	State               string    `json:"state"`
	ID                  int       `json:"id"`
}

func (sc *SonarrClient) ListCommands() ([]*Command, error) {
	retv := []*Command{}
	return retv, nil
}

func (sc *SonarrClient) CommandStatus(id int) (*Command, error) {
	return &Command{}, nil
}

func (sc *SonarrClient) RefreshSeries(series int) (*Command, error) {
	return &Command{}, nil
}

func (sc *SonarrClient) RescanSeries(series int) (*Command, error) {
	return &Command{}, nil
}

func (sc *SonarrClient) EpisodeSearch(eps []int) (*Command, error) {
	return &Command{}, nil
}

func (sc *SonarrClient) SeasonSearch(season, series int) (*Command, error) {
	return &Command{}, nil
}

func (sc *SonarrClient) SeriesSearch(series int) (*Command, error) {
	return &Command{}, nil
}

func (sc *SonarrClient) RssSync() (*Command, error) {
	return &Command{}, nil
}

func (sc *SonarrClient) RenameFiles(files []int) (*Command, error) {
	return &Command{}, nil
}

func (sc *SonarrClient) RenameSeries(series []int) (*Command, error) {
	return &Command{}, nil
}

func (sc *SonarrClient) Backup() (*Command, error) {
	return &Command{}, nil
}

func (sc *SonarrClient) MissingEpisodeSearch() (*Command, error) {
	return &Command{}, nil
}
