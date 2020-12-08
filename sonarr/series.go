package sonarr

import "time"

type SeriesImage struct {
	CoverType string `json:"coverType"`
	URL       string `json:"url"`
}

type SeasonStatistics struct {
	PreviousAiring    time.Time `json:"previousAiring"`
	EpisodeFileCount  int       `json:"episodeFileCount"`
	EpisodeCount      int       `json:"episodeCount"`
	TotalEpisodeCount int       `json:"totalEpisodeCount"`
	SizeOnDisk        int       `json:"sizeOnDisk"`
	PercentOfEpisodes float64   `json:"percentOfEpisodes"`
}

type SeriesSeason struct {
	SeasonNumber int              `json:"seasonNumber"`
	Monitored    bool             `json:"monitored"`
	Statistics   SeasonStatistics `json:"statistics,omitempty"`
}

type SeriesRatings struct {
	Votes int     `json:"votes"`
	Value float64 `json:"value"`
}

type QualityValue struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Weight int    `json:"weight"`
}

type QualityProfileValue struct {
	Name    string         `json:"name"`
	Allowed []QualityValue `json:"allowed"`
	Cutoff  QualityValue   `json:"cutoff"`
	ID      int            `json:"id"`
}

type QualityProfile struct {
	Value    QualityProfileValue `json:"value"`
	IsLoaded bool                `json:"isLoaded"`
}

type Series struct {
	Title             string         `json:"title"`
	AlternateTitles   []string       `json:"alternateTitles"`
	SortTitle         string         `json:"sortTitle"`
	SeasonCount       int            `json:"seasonCount"`
	TotalEpisodeCount int            `json:"totalEpisodeCount"`
	EpisodeCount      int            `json:"episodeCount"`
	EpisodeFileCount  int            `json:"episodeFileCount"`
	SizeOnDisk        int64          `json:"sizeOnDisk"`
	Status            string         `json:"status"`
	Overview          string         `json:"overview"`
	PreviousAiring    time.Time      `json:"previousAiring"`
	Network           string         `json:"network"`
	AirTime           string         `json:"airTime"`
	Images            []SeriesImage  `json:"images"`
	Seasons           []SeriesSeason `json:"seasons"`
	Year              int            `json:"year"`
	Path              string         `json:"path"`
	ProfileID         int            `json:"profileId"`
	SeasonFolder      bool           `json:"seasonFolder"`
	Monitored         bool           `json:"monitored"`
	UseSceneNumbering bool           `json:"useSceneNumbering"`
	Runtime           int            `json:"runtime"`
	TvdbID            int            `json:"tvdbId"`
	TvRageID          int            `json:"tvRageId"`
	TvMazeID          int            `json:"tvMazeId"`
	FirstAired        time.Time      `json:"firstAired"`
	LastInfoSync      time.Time      `json:"lastInfoSync"`
	SeriesType        string         `json:"seriesType"`
	CleanTitle        string         `json:"cleanTitle"`
	ImdbID            string         `json:"imdbId"`
	TitleSlug         string         `json:"titleSlug"`
	Certification     string         `json:"certification"`
	Genres            []string       `json:"genres"`
	Tags              []Tag          `json:"tags"`
	Added             time.Time      `json:"added"`
	Ratings           SeriesRatings  `json:"ratings"`
	QualityProfileID  int            `json:"qualityProfileId"`
	ID                int            `json:"id"`
}

func (sc *SonarrClient) ListSeries() ([]*Series, error) {
	return []*Series{}, nil
}
