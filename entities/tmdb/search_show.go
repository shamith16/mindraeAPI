package tmdb

type ShowSearch struct {
	Adult            bool        `json:"adult"`
	BackdropPath     string      `json:"backdrop_path"`
	CreatedBy        []CreatedBy `json:"created_by"`
	EpisodeRunTime   []int       `json:"episode_run_time"`
	FirstAirDate     string      `json:"first_air_date"`
	Genres           []Genres    `json:"genres"`
	Homepage         string      `json:"homepage"`
	Id               int         `json:"id"`
	InProduction     bool        `json:"in_production"`
	Languages        []string    `json:"languages"`
	LastAirDate      string      `json:"last_air_date"`
	LastEpisodeToAir Episodes    `json:"last_episode_to_air"`
	Name             string      `json:"name"`
	NextEpisodeToAir Episodes    `json:"next_episode_to_air,omitempty"`
	//Networks         []struct {
	//	Name          string `json:"name"`
	//	Id            int    `json:"id"`
	//	LogoPath      string `json:"logo_path"`
	//	OriginCountry string `json:"origin_country"`
	//} `json:"networks"`
	NumberOfEpisodes int      `json:"number_of_episodes"`
	NumberOfSeasons  int      `json:"number_of_seasons"`
	OriginCountry    []string `json:"origin_country"`
	OriginalLanguage string   `json:"original_language"`
	OriginalName     string   `json:"original_name"`
	Overview         string   `json:"overview"`
	Popularity       float64  `json:"popularity"`
	PosterPath       string   `json:"poster_path"`
	//ProductionCompanies []struct {
	//	Id            int    `json:"id"`
	//	LogoPath      string `json:"logo_path"`
	//	Name          string `json:"name"`
	//	OriginCountry string `json:"origin_country"`
	//} `json:"production_companies"`
	//ProductionCountries []struct {
	//	Iso31661 string `json:"iso_3166_1"`
	//	Name     string `json:"name"`
	//} `json:"production_countries"`
	Seasons         []Seasons         `json:"seasons"`
	SpokenLanguages []SpokenLanguages `json:"spoken_languages"`
	Status          string            `json:"status"`
	Tagline         string            `json:"tagline"`
	Type            string            `json:"type"`
	VoteAverage     float64           `json:"vote_average"`
	VoteCount       int               `json:"vote_count"`
	Credits         struct {
		Cast []Cast `json:"cast"`
		//Crew []struct {
		//	Adult              bool    `json:"adult"`
		//	Gender             int     `json:"gender"`
		//	Id                 int     `json:"id"`
		//	KnownForDepartment string  `json:"known_for_department"`
		//	Name               string  `json:"name"`
		//	OriginalName       string  `json:"original_name"`
		//	Popularity         float64 `json:"popularity"`
		//	ProfilePath        *string `json:"profile_path"`
		//	CreditId           string  `json:"credit_id"`
		//	Department         string  `json:"department"`
		//	Job                string  `json:"job"`
		//} `json:"crew"`
	} `json:"credits"`
	EpisodeGroups struct {
		Results []EpisodeGroupsResults `json:"results"`
	} `json:"episode_groups"`
	ExternalIds ExternalIds `json:"external_ids"`
	Images      Images      `json:"images"`
	//Keywords struct {
	//	Results []struct {
	//		Name string `json:"name"`
	//		Id   int    `json:"id"`
	//	} `json:"results"`
	//} `json:"keywords"`
	//Translations struct {
	//	Translations []struct {
	//		Iso31661    string `json:"iso_3166_1"`
	//		Iso6391     string `json:"iso_639_1"`
	//		Name        string `json:"name"`
	//		EnglishName string `json:"english_name"`
	//		Data        struct {
	//			Name     string `json:"name"`
	//			Overview string `json:"overview"`
	//			Homepage string `json:"homepage"`
	//			Tagline  string `json:"tagline"`
	//		} `json:"data"`
	//	} `json:"translations"`
	//} `json:"translations"`
	Videos struct {
		Results []VideoResults `json:"results"`
	} `json:"videos"`
}

type CreatedBy struct {
	Id          int    `json:"id"`
	CreditId    string `json:"credit_id"`
	Name        string `json:"name"`
	Gender      int    `json:"gender"`
	ProfilePath string `json:"profile_path"`
}

type Seasons struct {
	AirDate      string `json:"air_date"`
	EpisodeCount int    `json:"episode_count"`
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Overview     string `json:"overview"`
	PosterPath   string `json:"poster_path"`
	SeasonNumber int    `json:"season_number"`
}

type EpisodeGroupsResults struct {
	Description  string      `json:"description"`
	EpisodeCount int         `json:"episode_count"`
	GroupCount   int         `json:"group_count"`
	Id           string      `json:"id"`
	Name         string      `json:"name"`
	Network      interface{} `json:"network"`
	Type         int         `json:"type"`
}
