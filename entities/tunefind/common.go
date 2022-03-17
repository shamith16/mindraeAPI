package tunefind

type Image struct {
	Src     string `json:"src,omitempty"`
	Width   int    `json:"width,omitempty"`
	Height  int    `json:"height,omitempty"`
	Version string `json:"version,omitempty"`
}

type Games struct {
	ReleaseDate string `json:"release_date,omitempty"`
	SongsCount  int    `json:"songs_count,omitempty"`
	Name        string `json:"name,omitempty"`
	NameStub    string `json:"name_stub,omitempty"`
	Image       Image  `json:"image,omitempty"`
}

type ShowMovies struct {
	Name        string `json:"name"`
	NameStub    string `json:"name_stub"`
	Image       Image  `json:"image"`
	ReleaseDate string `json:"release_date"`
	SongsCount  int    `json:"songs_count"`
}
