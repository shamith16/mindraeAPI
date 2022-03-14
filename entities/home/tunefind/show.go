package tunefind

type TuneFindShowHome struct {
	Shows []struct {
		NameStub string `json:"name_stub,omitempty"`
		Name     string `json:"name,omitempty"`
		Image    struct {
			Src     string `json:"src,omitempty"`
			Width   int    `json:"width,omitempty"`
			Height  int    `json:"height,omitempty"`
			Version string `json:"version,omitempty"`
		} `json:"image,omitempty"`
		SongsCount  int `json:"songs_count,omitempty"`
		SeasonCount int `json:"season_count,omitempty"`
	} `json:"shows,omitempty"`
	AiringShows  []string `json:"airing_shows,omitempty"`
	NewShows     []string `json:"new_shows,omitempty"`
	PopularShows []string `json:"popular_shows,omitempty"`
}
