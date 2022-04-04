package tunefind

type MovieSearch struct {
	Movie      Movie       `json:"movie"`
	SongEvents []SongEvent `json:"song_events"`
	HotSongs   []Song      `json:"hot_songs"`
}

type Movie struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	NameStub      string `json:"name_stub"`
	GroupName     string `json:"group_name"`
	GroupSequence int64  `json:"group_sequence"`
	Type          string `json:"type"`
	SongID        int    `json:"song_id"`
	AirDateStart  int64  `json:"air_date_start"`
	AirDateEnd    int64  `json:"air_date_end"`
	TypeFull      string `json:"type_full"`
	SongsCount    int64  `json:"songs_count"`
	ReleaseDate   string `json:"release_date"`
	Event         Event  `json:"event"`
}

type Event struct {
	ID                int64      `json:"id"`
	EventGroupID      int64      `json:"event_group_id"`
	Name              string     `json:"name"`
	Number            string     `json:"number"`
	AirDate           int64      `json:"air_date"`
	PublicationDate   string     `json:"publication_date"`
	Locked            bool       `json:"locked"`
	Description       string     `json:"description"`
	AirdateDay        string     `json:"airdate_day"`
	AirdateMonth      string     `json:"airdate_month"`
	AirdateMonthShort string     `json:"airdate_month_short"`
	AirdateYear       string     `json:"airdate_year"`
	IsAired           bool       `json:"is_aired"`
	SongsCount        int64      `json:"songs_count,omitempty"`
	EventGroup        EventGroup `json:"event_group"`
}
