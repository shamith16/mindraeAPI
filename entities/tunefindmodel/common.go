package tunefindmodel

type Games struct {
	ReleaseDate string `json:"release_date,omitempty"`
	SongsCount  int    `json:"songs_count,omitempty"`
	Name        string `json:"name,omitempty"`
	NameStub    string `json:"name_stub,omitempty"`
}

type ShowMovies struct {
	Name        string `json:"name"`
	NameStub    string `json:"name_stub"`
	ReleaseDate string `json:"release_date"`
	SongsCount  int    `json:"songs_count"`
}
