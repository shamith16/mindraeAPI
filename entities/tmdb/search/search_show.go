package search

import "time"

type TmdbShowSearch struct {
	Adult        bool   `json:"adult,omitempty"`
	BackdropPath string `json:"backdrop_path,omitempty"`
	CreatedBy    []struct {
		Id          int    `json:"id,omitempty"`
		CreditId    string `json:"credit_id,omitempty"`
		Name        string `json:"name,omitempty"`
		Gender      int    `json:"gender,omitempty"`
		ProfilePath string `json:"profile_path,omitempty"`
	} `json:"created_by,omitempty"`
	EpisodeRunTime []int  `json:"episode_run_time,omitempty"`
	FirstAirDate   string `json:"first_air_date,omitempty"`
	Genres         []struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"genres,omitempty"`
	Homepage         string   `json:"homepage,omitempty"`
	Id               int      `json:"id,omitempty"`
	InProduction     bool     `json:"in_production,omitempty"`
	Languages        []string `json:"languages,omitempty"`
	LastAirDate      string   `json:"last_air_date,omitempty"`
	LastEpisodeToAir struct {
		AirDate        string  `json:"air_date,omitempty"`
		EpisodeNumber  int     `json:"episode_number,omitempty"`
		Id             int     `json:"id,omitempty"`
		Name           string  `json:"name,omitempty"`
		Overview       string  `json:"overview,omitempty"`
		ProductionCode string  `json:"production_code,omitempty"`
		SeasonNumber   int     `json:"season_number,omitempty"`
		StillPath      string  `json:"still_path,omitempty"`
		VoteAverage    float64 `json:"vote_average,omitempty"`
		VoteCount      int     `json:"vote_count,omitempty"`
	} `json:"last_episode_to_air,omitempty"`
	Name             string      `json:"name,omitempty"`
	NextEpisodeToAir interface{} `json:"next_episode_to_air,omitempty"`
	Networks         []struct {
		Name          string `json:"name,omitempty"`
		Id            int    `json:"id,omitempty"`
		LogoPath      string `json:"logo_path,omitempty"`
		OriginCountry string `json:"origin_country,omitempty"`
	} `json:"networks,omitempty"`
	NumberOfEpisodes    int      `json:"number_of_episodes,omitempty"`
	NumberOfSeasons     int      `json:"number_of_seasons,omitempty"`
	OriginCountry       []string `json:"origin_country,omitempty"`
	OriginalLanguage    string   `json:"original_language,omitempty"`
	OriginalName        string   `json:"original_name,omitempty"`
	Overview            string   `json:"overview,omitempty"`
	Popularity          float64  `json:"popularity,omitempty"`
	PosterPath          string   `json:"poster_path,omitempty"`
	ProductionCompanies []struct {
		Id            int    `json:"id,omitempty"`
		LogoPath      string `json:"logo_path,omitempty"`
		Name          string `json:"name,omitempty"`
		OriginCountry string `json:"origin_country,omitempty"`
	} `json:"production_companies,omitempty"`
	ProductionCountries []struct {
		Iso31661 string `json:"iso_3166_1,omitempty"`
		Name     string `json:"name,omitempty"`
	} `json:"production_countries,omitempty"`
	Seasons []struct {
		AirDate      string `json:"air_date,omitempty"`
		EpisodeCount int    `json:"episode_count,omitempty"`
		Id           int    `json:"id,omitempty"`
		Name         string `json:"name,omitempty"`
		Overview     string `json:"overview,omitempty"`
		PosterPath   string `json:"poster_path,omitempty"`
		SeasonNumber int    `json:"season_number,omitempty"`
	} `json:"seasons,omitempty"`
	SpokenLanguages []struct {
		EnglishName string `json:"english_name,omitempty"`
		Iso6391     string `json:"iso_639_1,omitempty"`
		Name        string `json:"name,omitempty"`
	} `json:"spoken_languages,omitempty"`
	Status      string  `json:"status,omitempty"`
	Tagline     string  `json:"tagline,omitempty"`
	Type        string  `json:"type,omitempty"`
	VoteAverage float64 `json:"vote_average,omitempty"`
	VoteCount   int     `json:"vote_count,omitempty"`
	Credits     struct {
		Cast []struct {
			Adult              bool    `json:"adult,omitempty"`
			Gender             int     `json:"gender,omitempty"`
			Id                 int     `json:"id,omitempty"`
			KnownForDepartment string  `json:"known_for_department,omitempty"`
			Name               string  `json:"name,omitempty"`
			OriginalName       string  `json:"original_name,omitempty"`
			Popularity         float64 `json:"popularity,omitempty"`
			ProfilePath        string  `json:"profile_path,omitempty"`
			Character          string  `json:"character,omitempty"`
			CreditId           string  `json:"credit_id,omitempty"`
			Order              int     `json:"order,omitempty"`
		} `json:"cast,omitempty"`
		Crew []struct {
			Adult              bool    `json:"adult,omitempty"`
			Gender             int     `json:"gender,omitempty"`
			Id                 int     `json:"id,omitempty"`
			KnownForDepartment string  `json:"known_for_department,omitempty"`
			Name               string  `json:"name,omitempty"`
			OriginalName       string  `json:"original_name,omitempty"`
			Popularity         float64 `json:"popularity,omitempty"`
			ProfilePath        *string `json:"profile_path,omitempty"`
			CreditId           string  `json:"credit_id,omitempty"`
			Department         string  `json:"department,omitempty"`
			Job                string  `json:"job,omitempty"`
		} `json:"crew,omitempty"`
	} `json:"credits,omitempty"`
	EpisodeGroups struct {
		Results []struct {
			Description  string      `json:"description,omitempty"`
			EpisodeCount int         `json:"episode_count,omitempty"`
			GroupCount   int         `json:"group_count,omitempty"`
			Id           string      `json:"id,omitempty"`
			Name         string      `json:"name,omitempty"`
			Network      interface{} `json:"network,omitempty"`
			Type         int         `json:"type,omitempty"`
		} `json:"results,omitempty"`
	} `json:"episode_groups,omitempty"`
	ExternalIds struct {
		ImdbId      string `json:"imdb_id,omitempty"`
		FreebaseMid string `json:"freebase_mid,omitempty"`
		FreebaseId  string `json:"freebase_id,omitempty"`
		TvdbId      int    `json:"tvdb_id,omitempty"`
		TvrageId    int    `json:"tvrage_id,omitempty"`
		FacebookId  string `json:"facebook_id,omitempty"`
		InstagramId string `json:"instagram_id,omitempty"`
		TwitterId   string `json:"twitter_id,omitempty"`
	} `json:"external_ids,omitempty"`
	Images struct {
		Backdrops []interface{} `json:"backdrops,omitempty"`
		Logos     []interface{} `json:"logos,omitempty"`
		Posters   []interface{} `json:"posters,omitempty"`
	} `json:"images,omitempty"`
	Keywords struct {
		Results []struct {
			Name string `json:"name,omitempty"`
			Id   int    `json:"id,omitempty"`
		} `json:"results,omitempty"`
	} `json:"keywords,omitempty"`
	Translations struct {
		Translations []struct {
			Iso31661    string `json:"iso_3166_1,omitempty"`
			Iso6391     string `json:"iso_639_1,omitempty"`
			Name        string `json:"name,omitempty"`
			EnglishName string `json:"english_name,omitempty"`
			Data        struct {
				Name     string `json:"name,omitempty"`
				Overview string `json:"overview,omitempty"`
				Homepage string `json:"homepage,omitempty"`
				Tagline  string `json:"tagline,omitempty"`
			} `json:"data,omitempty"`
		} `json:"translations,omitempty"`
	} `json:"translations,omitempty"`
	Videos struct {
		Results []struct {
			Iso6391     string    `json:"iso_639_1,omitempty"`
			Iso31661    string    `json:"iso_3166_1,omitempty"`
			Name        string    `json:"name,omitempty"`
			Key         string    `json:"key,omitempty"`
			Site        string    `json:"site,omitempty"`
			Size        int       `json:"size,omitempty"`
			Type        string    `json:"type,omitempty"`
			Official    bool      `json:"official,omitempty"`
			PublishedAt time.Time `json:"published_at,omitempty"`
			Id          string    `json:"id,omitempty"`
		} `json:"results,omitempty"`
	} `json:"videos,omitempty"`
}
