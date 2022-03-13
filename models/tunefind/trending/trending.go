package trending

type TrendingSongs struct {
	Songs []struct {
		Rank  int    `json:"rank"`
		Trend string `json:"trend"`
		Song  struct {
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
			Artists []struct {
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
			Amazon     string `json:"amazon"`
			Applemusic string `json:"applemusic"`
			Itunes     string `json:"itunes"`
			Spotify    string `json:"spotify"`
			Youtube    string `json:"youtube"`
		} `json:"song"`
		Events []struct {
			Rank  int `json:"rank"`
			Event struct {
				Id                int         `json:"id"`
				EventGroupId      int         `json:"event_group_id"`
				Name              string      `json:"name"`
				Number            *int        `json:"number"`
				AirDate           int         `json:"air_date"`
				PublicationDate   interface{} `json:"publication_date"`
				Locked            bool        `json:"locked"`
				Description       *string     `json:"description"`
				AirdateDay        string      `json:"airdate_day"`
				AirdateMonth      string      `json:"airdate_month"`
				AirdateMonthShort string      `json:"airdate_month_short"`
				AirdateYear       string      `json:"airdate_year"`
				IsAired           bool        `json:"is_aired"`
				Tombstone         interface{} `json:"tombstone"`
				TombstoneConflict bool        `json:"tombstone_conflict"`
				HasRightTombstone bool        `json:"has_right_tombstone"`
				Embargo           bool        `json:"embargo"`
				EventGroup        struct {
					Id            int    `json:"id"`
					Name          string `json:"name"`
					NameStub      string `json:"name_stub"`
					GroupName     string `json:"group_name"`
					GroupSequence int    `json:"group_sequence"`
					Type          string `json:"type"`
					SongId        *int   `json:"song_id"`
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
			} `json:"event"`
		} `json:"events"`
	} `json:"songs"`
}
