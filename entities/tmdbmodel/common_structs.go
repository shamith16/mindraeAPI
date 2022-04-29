package tmdbmodel

import jsoniter "github.com/json-iterator/go"

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Genres struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type MovieExternalIDs struct {
	IMDbID      string `json:"imdb_id,omitempty"`
	FacebookID  string `json:"facebook_id,omitempty"`
	InstagramID string `json:"instagram_id,omitempty"`
	TwitterID   string `json:"twitter_id,omitempty"`
	ID          int64  `json:"id,omitempty,omitempty"`
}

type Images struct {
	ID        int64 `json:"id,omitempty"`
	Backdrops []struct {
		AspectRatio float32 `json:"aspect_ratio"`
		FilePath    string  `json:"file_path"`
		Height      int     `json:"height"`
		Iso639_1    string  `json:"iso_639_1"`
		VoteAverage float32 `json:"vote_average"`
		VoteCount   int64   `json:"vote_count"`
		Width       int     `json:"width"`
	} `json:"backdrops"`
	Logos []struct {
		AspectRatio float32 `json:"aspect_ratio"`
		FilePath    string  `json:"file_path"`
		Height      int     `json:"height"`
		Iso639_1    string  `json:"iso_639_1"`
		VoteAverage float32 `json:"vote_average"`
		VoteCount   int64   `json:"vote_count"`
		Width       int     `json:"width"`
	} `json:"logos"`
	Posters []struct {
		AspectRatio float32 `json:"aspect_ratio"`
		FilePath    string  `json:"file_path"`
		Height      int     `json:"height"`
		Iso639_1    string  `json:"iso_639_1"`
		VoteAverage float32 `json:"vote_average"`
		VoteCount   int64   `json:"vote_count"`
		Width       int     `json:"width"`
	} `json:"posters"`
}

type Videos struct {
	VideoResults []struct {
		Iso_639_1   string          `json:"iso_639_1,omitempty"`
		Iso_3166_1  string          `json:"iso_3166_1,omitempty"`
		Name        string          `json:"name,omitempty"`
		Key         string          `json:"key,omitempty"`
		Site        string          `json:"site,omitempty"`
		Size        jsoniter.Number `json:"size,omitempty"`
		Type        string          `json:"type,omitempty"`
		Official    bool            `json:"official,omitempty"`
		PublishedAt string          `json:"publishedAt,omitempty"`
		Id          string          `json:"id,omitempty"`
	} `json:"videos"`
}

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
