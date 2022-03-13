package home

type MovieGame struct {
	Slider        []string `json:"slider"`
	Featured      []string `json:"featured"`
	RecentlyAdded []string `json:"recently_added"`
	Contribute    []string `json:"contribute"`
	Games         []struct {
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
