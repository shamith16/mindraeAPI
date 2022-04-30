package tmdbmodel

import jsoniter "github.com/json-iterator/go"

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type SearchMovieResult struct {
	VoteCount        int64   `json:"vote_count"`
	ID               int64   `json:"id"`
	Video            bool    `json:"video"`
	VoteAverage      float32 `json:"vote_average"`
	Title            string  `json:"title"`
	Popularity       float32 `json:"popularity"`
	PosterPath       string  `json:"poster_path"`
	OriginalLanguage string  `json:"original_language"`
	OriginalTitle    string  `json:"original_title"`
	GenreIDs         []int64 `json:"genre_ids"`
	BackdropPath     string  `json:"backdrop_path"`
	Adult            bool    `json:"adult"`
	Overview         string  `json:"overview"`
	ReleaseDate      string  `json:"release_date"`
}
