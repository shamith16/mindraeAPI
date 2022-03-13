package home

type TuneFindMovieGameHome struct {
	Slider        []string `json:"slider,omitempty"`
	Featured      []string `json:"featured,omitempty"`
	RecentlyAdded []string `json:"recently_added,omitempty"`
	Contribute    []string `json:"contribute,omitempty"`
	Games         []struct {
		ReleaseDate string `json:"release_date,omitempty"`
		SongsCount  int    `json:"songs_count,omitempty"`
		Name        string `json:"name,omitempty"`
		NameStub    string `json:"name_stub,omitempty"`
		Image       struct {
			Src     string `json:"src,omitempty"`
			Width   int    `json:"width,omitempty"`
			Height  int    `json:"height,omitempty"`
			Version string `json:"version,omitempty"`
		} `json:"image,omitempty"`
	} `json:"games,omitempty"`
	Movies []struct {
		ReleaseDate string `json:"release_date,omitempty"`
		SongsCount  int    `json:"songs_count,omitempty"`
		Name        string `json:"name,omitempty"`
		NameStub    string `json:"name_stub,omitempty"`
		Image       struct {
			Src     string `json:"src,omitempty"`
			Width   int    `json:"width,omitempty"`
			Height  int    `json:"height,omitempty"`
			Version string `json:"version,omitempty"`
		} `json:"image,omitempty"`
	}
}
