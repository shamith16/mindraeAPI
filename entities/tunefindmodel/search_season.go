package tunefindmodel

import "encoding/json"

type SeasonSearch struct {
	Season    Season         `json:"season"`
	Episodes  []Episode      `json:"episodes"`
	ThemeSong ThemeHotSong   `json:"theme_song"`
	HotSongs  []ThemeHotSong `json:"hot_songs"`
}

type Image struct {
	Src     string      `json:"src"`
	Width   json.Number `json:"width"`
	Height  json.Number `json:"height"`
	Version string      `json:"version"`
}

type Season struct {
	Id               int                         `json:"id"`
	Name             string                      `json:"name"`
	NameStub         string                      `json:"name_stub"`
	GroupName        string                      `json:"group_name"`
	GroupSequence    int                         `json:"group_sequence"`
	Type             string                      `json:"type"`
	SongId           int                         `json:"song_id"`
	AirDateStart     int                         `json:"air_date_start"`
	AirDateEnd       int                         `json:"air_date_end"`
	TypeFull         string                      `json:"type_full"`
	Image            Image                       `json:"image"`
	SongsCount       int                         `json:"songs_count"`
	EpisodesCount    int                         `json:"episodes_count,omitempty"`
	MusicSupervisors []MusicSupervisorsComposers `json:"music_supervisors,omitempty"`
	Composers        []MusicSupervisorsComposers `json:"composers,omitempty"`
	ThemeSong        ThemeHotSong                `json:"theme_song,omitempty"`
}

type Episode struct {
	Id                int        `json:"id"`
	EventGroupId      int        `json:"event_group_id"`
	Name              string     `json:"name"`
	Number            int        `json:"number"`
	AirDate           int        `json:"air_date"`
	PublicationDate   string     `json:"publication_date"`
	Locked            bool       `json:"locked"`
	Description       string     `json:"description"`
	AirdateDay        string     `json:"airdate_day"`
	AirdateMonth      string     `json:"airdate_month"`
	AirdateMonthShort string     `json:"airdate_month_short"`
	AirdateYear       string     `json:"airdate_year"`
	IsAired           bool       `json:"is_aired"`
	EventGroup        EventGroup `json:"event_group"`
	SongsCount        int        `json:"songs_count"`
}

type ThemeHotSong struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	NameStub   string `json:"name_stub"`
	Album      string `json:"album"`
	PreviewUrl string `json:"preview_url"`
	Url        URL    `json:"url"`
	//	Amazon     string   `json:"amazon"`
	Applemusic string   `json:"applemusic"`
	Itunes     string   `json:"itunes"`
	Spotify    string   `json:"spotify"`
	Youtube    string   `json:"youtube"`
	Artists    []Artist `json:"artists"`
}
