package tmdbmodel

type MovieSearch struct {
	Adult            bool              `json:"adult"`
	BackdropPath     string            `json:"backdrop_path"`
	Budget           int               `json:"budget"`
	Genres           []Genres          `json:"genres"`
	Homepage         string            `json:"homepage"`
	Id               int               `json:"id"`
	ImdbId           string            `json:"imdb_id"`
	OriginalLanguage string            `json:"original_language"`
	OriginalTitle    string            `json:"original_title"`
	Overview         string            `json:"overview"`
	Popularity       float64           `json:"popularity"`
	PosterPath       string            `json:"poster_path"`
	ReleaseDate      string            `json:"release_date"`
	Revenue          int               `json:"revenue"`
	Runtime          int               `json:"runtime"`
	SpokenLanguages  []SpokenLanguages `json:"spoken_languages"`
	Status           string            `json:"status"`
	Tagline          string            `json:"tagline"`
	Title            string            `json:"title"`
	Video            bool              `json:"video"`
	VoteAverage      float64           `json:"vote_average"`
	VoteCount        int               `json:"vote_count"`
	Credits          Credits           `json:"credits"`
	ExternalIds      ExternalIds       `json:"external_ids"`
	Images           Images            `json:"images,omitempty"`
	Videos           Videos            `json:"videos"`
}

type Videos struct {
	Results []VideoResults `json:"results"`
}

type Credits struct {
	Cast []Cast `json:"cast"`
}
