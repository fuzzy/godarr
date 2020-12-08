package main

type Tag struct {
	Label string `json:"label"`
	ID    int    `json:"id"`
}

func (sc *SonarrClient) GetTags() ([]*Tag, error) {
	retv := []*Tag{}
	return retv, nil
}

func (sc *SonarrClient) GetTag(tag int) (*Tag, error) {
	return &Tag{}, nil
}

func (sc *SonarrClient) NewTag(label string) (*Tag, error) {
	return &Tag{}, nil
}

func (tag *Tag) UpdateTag(label string) (*Tag, error) {
	return &Tag{}, nil
}
