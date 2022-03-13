package browse

type TuneFindShowBrowse struct {
	Shows []struct {
		NameStub string `json:"name_stub"`
		Name     string `json:"name"`
		Image    struct {
			Src     string `json:"src"`
			Width   int    `json:"width"`
			Height  int    `json:"height"`
			Version string `json:"version"`
		} `json:"image"`
		SongsCount  int `json:"songs_count"`
		SeasonCount int `json:"season_count"`
	} `json:"shows"`
}
