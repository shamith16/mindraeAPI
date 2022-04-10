package tmdbmodel

import "time"

type Genres struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Cast struct {
	Adult              bool    `json:"adult"`
	Gender             int     `json:"gender"`
	Id                 int     `json:"id"`
	KnownForDepartment string  `json:"known_for_department"`
	Name               string  `json:"name"`
	OriginalName       string  `json:"original_name"`
	Popularity         float64 `json:"popularity"`
	ProfilePath        string  `json:"profile_path"`
	CastId             int     `json:"cast_id,omitempty"`
	Character          string  `json:"character"`
	CreditId           string  `json:"credit_id"`
	Order              int     `json:"order"`
}

type ExternalIds struct {
	ImdbId      string `json:"imdb_id,omitempty"`
	FreebaseMid string `json:"freebase_mid,omitempty"`
	FreebaseId  string `json:"freebase_id,omitempty"`
	TvdbId      int    `json:"tvdb_id,omitempty"`
	TvrageId    int    `json:"tvrage_id,omitempty"`
	FacebookId  string `json:"facebook_id,omitempty"`
	InstagramId string `json:"instagram_id,omitempty"`
	TwitterId   string `json:"twitter_id,omitempty"`
}

type VideoResults struct {
	Iso6391     string    `json:"iso_639_1"`
	Iso31661    string    `json:"iso_3166_1"`
	Name        string    `json:"name"`
	Key         string    `json:"key"`
	Site        string    `json:"site"`
	Size        int       `json:"size"`
	Type        string    `json:"type"`
	Official    bool      `json:"official"`
	PublishedAt time.Time `json:"published_at"`
	Id          string    `json:"id"`
}

type SpokenLanguages struct {
	EnglishName string `json:"english_name"`
	Iso6391     string `json:"iso_639_1"`
	Name        string `json:"name"`
}

type Images struct {
	Id        int           `json:"id"`
	Backdrops []imagesTypes `json:"backdrops,omitempty"`
	Posters   []imagesTypes `json:"posters,omitempty"`
	Logos     []imagesTypes `json:"logos,omitempty"`
}

type imagesTypes struct {
	AspectRatio float64 `json:"aspect_ratio"`
	FilePath    string  `json:"file_path"`
	Height      int     `json:"height"`
	Iso6391     string  `json:"iso_639_1"`
	VoteAverage float64 `json:"vote_average"`
	VoteCount   int     `json:"vote_count"`
	Width       int     `json:"width"`
}
