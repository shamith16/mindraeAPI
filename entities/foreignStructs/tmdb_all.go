package foreignStructs

import "github.com/shamith16/mindraeAPI/entities/tmdb"

type TmdbSearchAll struct {
	TmdbMovie     tmdb.ShowMovieBrowse `json:"tmdbMovie"`
	TmdbShow      tmdb.ShowMovieBrowse `json:"tmdbShow"`
	MovieSearch   tmdb.MovieSearch     `json:"tmdbMovieSearch"`
	ShowSearch    tmdb.ShowSearch      `json:"tmdbShowSearch"`
	SeasonSearch  tmdb.SeasonSearch    `json:"tmdbSeasonSearch"`
	EpisodeSearch tmdb.Episodes        `json:"tmdbEpisodeSearch"`
}
