package tunefind

type EpisodeSearch struct {
	Episode EpisodeSearchEpisode `json:"episode"`
}

type EpisodeSearchEpisode struct {
	ID                int64  `json:"id"`
	EventGroupID      int64  `json:"event_group_id"`
	Name              string `json:"name"`
	Number            int64  `json:"number"`
	AirDate           int64  `json:"air_date"`
	Locked            bool   `json:"locked"`
	Description       string `json:"description"`
	AirdateDay        string `json:"airdate_day"`
	AirdateMonth      string `json:"airdate_month"`
	AirdateMonthShort string `json:"airdate_month_short"`
	AirdateYear       string `json:"airdate_year"`
	IsAired           bool   `json:"is_aired"`

	EventGroup EventGroup    `json:"event_group"`
	SongEvents []SongEvent   `json:"song_events"`
	Next       Next          `json:"next"`
	Previous   Next          `json:"previous"`
	Votes      []interface{} `json:"votes"`
}

type EventGroup struct {
	ID            int64       `json:"id"`
	Name          string      `json:"name"`
	NameStub      string      `json:"name_stub"`
	GroupName     string      `json:"group_name"`
	GroupSequence int64       `json:"group_sequence"`
	Type          string      `json:"type"`
	SongID        interface{} `json:"song_id"`
	AirDateStart  int64       `json:"air_date_start"`
	AirDateEnd    int64       `json:"air_date_end"`
	TypeFull      string      `json:"type_full"`
}

type Next struct {
	ID                int64       `json:"id"`
	EventGroupID      int64       `json:"event_group_id"`
	Name              string      `json:"name"`
	Number            int64       `json:"number"`
	AirDate           int64       `json:"air_date"`
	PublicationDate   interface{} `json:"publication_date"`
	Locked            bool        `json:"locked"`
	Description       string      `json:"description"`
	AirdateDay        string      `json:"airdate_day"`
	AirdateMonth      string      `json:"airdate_month"`
	AirdateMonthShort string      `json:"airdate_month_short"`
	AirdateYear       string      `json:"airdate_year"`
	IsAired           bool        `json:"is_aired"`
}

type SongEvent struct {
	ID                 int64       `json:"id"`
	SongID             int64       `json:"song_id"`
	EventID            int64       `json:"event_id"`
	DateCreated        string      `json:"date_created"`
	UserID             string      `json:"user_id"`
	Position           int64       `json:"position"`
	DescriptionHistory interface{} `json:"description_history"`
	IsRight            bool        `json:"is_right"`
	IsWrong            bool        `json:"is_wrong"`
	IsForVote          bool        `json:"is_for_vote"`
	Link               string      `json:"link"`
	Locked             bool        `json:"locked"`
	Description        *string     `json:"description"`
	IsSeasonOnly       bool        `json:"is_season_only"`
	AlbumID            interface{} `json:"album_id"`
	AlbumSongNo        interface{} `json:"album_song_no"`
	Song               Song        `json:"song"`
}

type Song struct {
	ID         int64    `json:"id"`
	Name       string   `json:"name"`
	NameStub   string   `json:"name_stub"`
	Album      string   `json:"album"`
	PreviewURL string   `json:"preview_url"`
	URL        URL      `json:"url"`
	Link       string   `json:"link"`
	Amazon     string   `json:"amazon"`
	Applemusic string   `json:"applemusic"`
	Itunes     string   `json:"itunes"`
	Spotify    string   `json:"spotify"`
	Youtube    string   `json:"youtube"`
	Artists    []Artist `json:"artists"`
}

type Artist struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	NameStub string `json:"name_stub"`
}

type URL struct {
	ID                 int64  `json:"id"`
	SongID             int64  `json:"song_id"`
	Status             string `json:"status"`
	ExternalPreviewURL string `json:"external_preview_url"`
	ExternalArtURL     string `json:"external_art_url"`
	ExternalDuration   int64  `json:"external_duration"`
	TrackID            string `json:"track_id"`
}
