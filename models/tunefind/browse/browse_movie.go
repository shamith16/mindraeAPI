package browse

type TuneFindMovieBrowse struct {
	Movies []struct {
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
	} `json:"movies"`
}
