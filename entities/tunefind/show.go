package tunefind

type ShowHome struct {
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
	AiringShows  []string `json:"airing_shows"`
	NewShows     []string `json:"new_shows"`
	PopularShows []string `json:"popular_shows"`
}
