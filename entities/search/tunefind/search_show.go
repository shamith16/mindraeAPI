package tunefind

type TuneFindShowSearch struct {
	Show struct {
		NameStub string `json:"name_stub"`
		Name     string `json:"name"`
		Image    struct {
			Src     string `json:"src"`
			Width   int    `json:"width"`
			Height  int    `json:"height"`
			Version string `json:"version"`
		} `json:"image"`
		Features struct {
			CanFollow bool `json:"canFollow"`
		} `json:"features"`
	} `json:"show"`
	Seasons []struct {
		Id            int    `json:"id"`
		Name          string `json:"name"`
		NameStub      string `json:"name_stub"`
		GroupName     string `json:"group_name"`
		GroupSequence int    `json:"group_sequence"`
		Type          string `json:"type"`
		SongId        int    `json:"song_id"`
		AirDateStart  int    `json:"air_date_start"`
		AirDateEnd    int    `json:"air_date_end"`
		TypeFull      string `json:"type_full"`
		Image         struct {
			Src     string `json:"src"`
			Width   int    `json:"width"`
			Height  int    `json:"height"`
			Version string `json:"version"`
		} `json:"image"`
		SeasonFeatures struct {
			CanAdmin bool `json:"canAdmin"`
		} `json:"season_features"`
		SongsCount       int           `json:"songs_count"`
		EpisodesCount    int           `json:"episodes_count"`
		MusicSupervisors []interface{} `json:"music_supervisors"`
		Composers        []interface{} `json:"composers"`
		ThemeSong        struct {
			Id         int    `json:"id"`
			Name       string `json:"name"`
			NameStub   string `json:"name_stub"`
			Album      string `json:"album"`
			PreviewUrl string `json:"preview_url"`
			Url        struct {
				Id                 int    `json:"id"`
				SongId             int    `json:"song_id"`
				Status             string `json:"status"`
				ExternalPreviewUrl string `json:"external_preview_url"`
				ExternalArtUrl     string `json:"external_art_url"`
				ExternalDuration   int    `json:"external_duration"`
				TrackId            string `json:"track_id"`
			} `json:"url"`
			Amazon     string `json:"amazon"`
			Applemusic string `json:"applemusic"`
			Itunes     string `json:"itunes"`
			Spotify    string `json:"spotify"`
			Youtube    string `json:"youtube"`
			Artists    []struct {
				Id       int    `json:"id"`
				Name     string `json:"name"`
				NameStub string `json:"name_stub"`
				Image    struct {
					Src     string `json:"src"`
					Width   int    `json:"width"`
					Height  int    `json:"height"`
					Version string `json:"version"`
				} `json:"image"`
			} `json:"artists"`
		} `json:"theme_song"`
	} `json:"seasons"`
	Following     bool `json:"following"`
	LatestEpisode struct {
		Id                int         `json:"id"`
		EventGroupId      int         `json:"event_group_id"`
		Name              string      `json:"name"`
		Number            int         `json:"number"`
		AirDate           int         `json:"air_date"`
		PublicationDate   interface{} `json:"publication_date"`
		Locked            bool        `json:"locked"`
		Description       string      `json:"description"`
		AirdateDay        string      `json:"airdate_day"`
		AirdateMonth      string      `json:"airdate_month"`
		AirdateMonthShort string      `json:"airdate_month_short"`
		AirdateYear       string      `json:"airdate_year"`
		IsAired           bool        `json:"is_aired"`
		SongsCount        int         `json:"songs_count"`
	} `json:"latest_episode"`
}
