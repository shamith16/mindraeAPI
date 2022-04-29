package mindraemodel

import (
	tmdb "github.com/cyruzin/golang-tmdb"
	"github.com/shamith16/mindraeAPI/entities/tunefindmodel"
)

type MovieHome struct {
	Movies []Movie `json:"movies,omitempty"`
}

type Movie struct {
	TuneFind TuneFind          `json:"tuneFind,omitempty"`
	Tmdb     tmdb.MovieDetails `json:"tmdb,omitempty"`
}

type TuneFind struct {
	TuneFindId            int64                     `json:"Id,omitempty"`
	TuneFindName          string                    `json:"Name,omitempty"`
	TuneFindNameStub      string                    `json:"NameStub,omitempty"`
	TuneFindType          string                    `json:"Type,omitempty"`
	TuneFindTypeFull      string                    `json:"TypeFull,omitempty"`
	TuneFindSongsCount    int64                     `json:"SongsCount,omitempty"`
	TuneFindReleaseDate   string                    `json:"ReleaseDate,omitempty"`
	TuneFindAirDay        string                    `json:"AirDay,omitempty"`
	TuneFindAirMonth      string                    `json:"AirMonth,omitempty"`
	TuneFindAirMonthShort string                    `json:"AirMonthShort,omitempty"`
	TuneFindAirYear       string                    `json:"AirYear,omitempty"`
	TuneFindIsAired       bool                      `json:"IsAired,omitempty"`
	SongEvents            []tunefindmodel.SongEvent `json:"songEvents,omitempty"`
	HotSongs              []tunefindmodel.Song      `json:"hotSongs,omitempty"`
}
