package search

type Movie struct {
	Movie struct {
		Id            int         `json:"id"`
		Name          string      `json:"name"`
		NameStub      string      `json:"name_stub"`
		GroupName     string      `json:"group_name"`
		GroupSequence int         `json:"group_sequence"`
		Type          string      `json:"type"`
		SongId        interface{} `json:"song_id"`
		AirDateStart  int         `json:"air_date_start"`
		AirDateEnd    int         `json:"air_date_end"`
		TypeFull      string      `json:"type_full"`
		Image         struct {
			Src     string `json:"src"`
			Width   int    `json:"width"`
			Height  int    `json:"height"`
			Version string `json:"version"`
		} `json:"image"`
		SongsCount  int    `json:"songs_count"`
		ReleaseDate string `json:"release_date"`
		Features    struct {
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
		Event struct {
			Id                int         `json:"id"`
			EventGroupId      int         `json:"event_group_id"`
			Name              string      `json:"name"`
			Number            interface{} `json:"number"`
			AirDate           int         `json:"air_date"`
			PublicationDate   interface{} `json:"publication_date"`
			Locked            bool        `json:"locked"`
			Description       interface{} `json:"description"`
			AirdateDay        string      `json:"airdate_day"`
			AirdateMonth      string      `json:"airdate_month"`
			AirdateMonthShort string      `json:"airdate_month_short"`
			AirdateYear       string      `json:"airdate_year"`
			IsAired           bool        `json:"is_aired"`
			SongsCount        int         `json:"songs_count"`
			EventGroup        struct {
				Id            int         `json:"id"`
				Name          string      `json:"name"`
				NameStub      string      `json:"name_stub"`
				GroupName     string      `json:"group_name"`
				GroupSequence int         `json:"group_sequence"`
				Type          string      `json:"type"`
				SongId        interface{} `json:"song_id"`
				AirDateStart  int         `json:"air_date_start"`
				AirDateEnd    int         `json:"air_date_end"`
				TypeFull      string      `json:"type_full"`
				Image         struct {
					Src     string `json:"src"`
					Width   int    `json:"width"`
					Height  int    `json:"height"`
					Version string `json:"version"`
				} `json:"image"`
			} `json:"event_group"`
			Features struct {
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
		} `json:"event"`
	} `json:"movie"`
	SongEvents []struct {
		Id          int    `json:"id"`
		SongId      int    `json:"song_id"`
		EventId     int    `json:"event_id"`
		DateCreated string `json:"date_created"`
		UserId      string `json:"user_id"`
		Position    int    `json:"position"`
		Features    struct {
			CanAdmin        bool `json:"canAdmin"`
			CanVote         bool `json:"canVote"`
			CanAddSceneDesc bool `json:"canAddSceneDesc"`
			CanAddFavorite  bool `json:"canAddFavorite"`
			CanDelete       bool `json:"canDelete"`
		} `json:"features"`
		DescriptionHistory interface{} `json:"description_history"`
		IsRight            bool        `json:"is_right"`
		IsWrong            bool        `json:"is_wrong"`
		IsForVote          bool        `json:"is_for_vote"`
		Link               string      `json:"link"`
		Locked             bool        `json:"locked"`
		Description        *string     `json:"description"`
		IsSeasonOnly       bool        `json:"is_season_only"`
		AlbumId            interface{} `json:"album_id"`
		AlbumSongNo        interface{} `json:"album_song_no"`
		Song               struct {
			Id         int     `json:"id"`
			Name       string  `json:"name"`
			NameStub   string  `json:"name_stub"`
			Album      string  `json:"album"`
			PreviewUrl *string `json:"preview_url"`
			Url        *struct {
				Id                 int    `json:"id"`
				SongId             int    `json:"song_id"`
				Status             string `json:"status"`
				ExternalPreviewUrl string `json:"external_preview_url"`
				ExternalArtUrl     string `json:"external_art_url"`
				ExternalDuration   int    `json:"external_duration"`
				TrackId            string `json:"track_id"`
			} `json:"url"`
			Link       string  `json:"link"`
			Amazon     string  `json:"amazon"`
			Applemusic *string `json:"applemusic"`
			Itunes     *string `json:"itunes"`
			Spotify    *string `json:"spotify"`
			Youtube    string  `json:"youtube"`
			Artists    []struct {
				Id       int    `json:"id"`
				Name     string `json:"name"`
				NameStub string `json:"name_stub"`
				Image    struct {
					Src     string      `json:"src"`
					Width   interface{} `json:"width"`
					Height  interface{} `json:"height"`
					Version string      `json:"version"`
				} `json:"image"`
			} `json:"artists"`
		} `json:"song"`
	} `json:"song_events"`
	Votes     []interface{} `json:"votes"`
	HotSongs  []interface{} `json:"hot_songs"`
	Following bool          `json:"following"`
}
