package mindraemodel

import (
	"github.com/shamith16/mindraeAPI/entities/tmdbmodel"
	"github.com/shamith16/mindraeAPI/entities/tunefindmodel"
)

type MovieHome struct {
	Movies []Movie `json:"movies,omitempty"`
}

type Movie struct {
	TuneFindId            int64                     `json:"tuneId,omitempty"`
	TuneFindName          string                    `json:"tuneName,omitempty"`
	TuneFindNameStub      string                    `json:"tuneNameStub,omitempty"`
	TuneFindType          string                    `json:"tuneType,omitempty"`
	TuneFindTypeFull      string                    `json:"tuneTypeFull,omitempty"`
	TuneFindSongsCount    int64                     `json:"tuneSongsCount,omitempty"`
	TuneFindReleaseDate   string                    `json:"tuneReleaseDate,omitempty"`
	TuneFindAirDay        string                    `json:"tuneAirDay,omitempty"`
	TuneFindAirMonth      string                    `json:"tuneAirMonth,omitempty"`
	TuneFindAirMonthShort string                    `json:"tuneAirMonthShort,omitempty"`
	TuneFindAirYear       string                    `json:"tuneAirYear,omitempty"`
	TuneFindIsAired       bool                      `json:"tuneIsAired,omitempty"`
	SongEvents            []tunefindmodel.SongEvent `json:"songEvents,omitempty"`
	HotSongs              []tunefindmodel.Song      `json:"hotSongs,omitempty"`
	TmdbIsAdult           bool                      `json:"tmdbIsAdult,omitempty"`
	TmdbBackDropPath      string                    `json:"tmdbBackDropPath,omitempty"`
	TmdbBudget            int64                     `json:"tmdbBudget,omitempty"`
	Genres                []tmdbmodel.Genres        `json:"genres,omitempty"`
	MovieHomePageUrl      string                    `json:"movieHomePageUrl,omitempty"`
	TmdbId                int64                     `json:"tmdbId,omitempty"`
	ImdbId                string                    `json:"imdbId,omitempty"`
	OriginalLanguage      string                    `json:"originalLanguage,omitempty"`
	OriginalTitle         string                    `json:"originalTitle,omitempty"`
	MovieOverview         string                    `json:"movieOverview,omitempty"`
	Popularity            float64                   `json:"popularity,omitempty"`
	PosterPath            string                    `json:"posterPath,omitempty"`
	TmdbReleaseDate       string                    `json:"tmdbReleaseDate,omitempty"`
	Revenue               int64                     `json:"revenue,omitempty"`
	Runtime               int64                     `json:"runtime,omitempty"`
	TmdbStatus            string                    `json:"tmdbStatus,omitempty"`
	TagLine               string                    `json:"tagLine,omitempty"`
	Title                 string                    `json:"title,omitempty"`
	Video                 bool                      `json:"video,omitempty"`
	VoteAverage           float64                   `json:"voteAverage,omitempty"`
	VoteCount             int64                     `json:"voteCount,omitempty"`
	ExternalIds           tmdbmodel.ExternalIds     `json:"externalIds,omitempty"`
	Videos                tmdbmodel.Videos          `json:"videos,omitempty"`
	Images                tmdbmodel.Images          `json:"images,omitempty"`
}
