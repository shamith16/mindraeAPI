package search

type Season struct {
	Season struct {
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
		SongsCount int `json:"songs_count"`
	} `json:"season"`
	Episodes []struct {
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
		EventGroup        struct {
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
		} `json:"event_group"`
		Tombstone *struct {
			Id          int    `json:"id"`
			EventId     int    `json:"event_id"`
			DateCreated string `json:"date_created"`
			UserId      string `json:"user_id"`
			IsRight     bool   `json:"is_right"`
			IsWrong     bool   `json:"is_wrong"`
			IsForVote   bool   `json:"is_for_vote"`
			Locked      bool   `json:"locked"`
			Features    struct {
				CanAdmin  bool `json:"canAdmin"`
				CanVote   bool `json:"canVote"`
				CanDelete bool `json:"canDelete"`
			} `json:"features"`
		} `json:"tombstone"`
		TombstoneConflict bool `json:"tombstone_conflict"`
		HasRightTombstone bool `json:"has_right_tombstone"`
		Embargo           bool `json:"embargo"`
		QuestionCount     int  `json:"question_count"`
		SongsCount        int  `json:"songs_count"`
		Features          struct {
			CanView               bool `json:"canView"`
			CanAddSong            bool `json:"canAddSong"`
			CanAddAlbum           bool `json:"canAddAlbum"`
			CanAddSongDescription bool `json:"canAddSongDescription"`
			CanDisableAuthSong    bool `json:"canDisableAuthSong"`
			CanAddQuestion        bool `json:"canAddQuestion"`
			CanSortSongs          bool `json:"canSortSongs"`
			CanFollow             bool `json:"canFollow"`
			CanAdmin              bool `json:"canAdmin"`
			CanLock               bool `json:"canLock"`
			CanAddTombstone       bool `json:"canAddTombstone"`
		} `json:"features"`
	} `json:"episodes"`
	ThemeSong struct {
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
	HotSongs []struct {
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
	} `json:"hot_songs"`
}
