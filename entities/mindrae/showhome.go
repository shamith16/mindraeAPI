package mindrae

import (
	"github.com/shamith16/mindraeAPI/entities/tmdb"
	"github.com/shamith16/mindraeAPI/entities/tunefind"
)

type MovieHome struct {
	Slider   []MergeStructure `json:"slider"`
	Featured []MergeStructure `json:"featured,omitempty"`
	Recent   []MergeStructure `json:"recent,omitempty"`
}

type MergeStructure struct {
	TuneFindMovie tunefind.MovieSearch `json:"tunefindmovie,omitempty"`
	TmdbMovie     tmdb.MovieSearch     `json:"tmdbmovie,omitempty"`
}
