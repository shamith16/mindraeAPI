package tunefind

type GameBrowse struct {
	Games []struct {
		ReleaseDate string `json:"release_date"`
		SongsCount  int    `json:"songs_count"`
		Name        string `json:"name"`
		NameStub    string `json:"name_stub"`
		Image       struct {
			Src     string `json:"src"`
			Width   int    `json:"width"`
			Height  int    `json:"height"`
			Version string `json:"version"`
		} `json:"image"`
	} `json:"games"`
}
