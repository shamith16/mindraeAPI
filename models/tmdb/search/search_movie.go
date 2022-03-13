package search

import "time"

type TmdbMovieSearch struct {
	Adult               bool        `json:"adult,omitempty"`
	BackdropPath        string      `json:"backdrop_path,omitempty"`
	BelongsToCollection interface{} `json:"belongs_to_collection,omitempty"`
	Budget              int         `json:"budget,omitempty"`
	Genres              []struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"genres,omitempty"`
	Homepage            string  `json:"homepage,omitempty"`
	Id                  int     `json:"id,omitempty"`
	ImdbId              string  `json:"imdb_id,omitempty"`
	OriginalLanguage    string  `json:"original_language,omitempty"`
	OriginalTitle       string  `json:"original_title,omitempty"`
	Overview            string  `json:"overview,omitempty"`
	Popularity          float64 `json:"popularity,omitempty"`
	PosterPath          string  `json:"poster_path,omitempty"`
	ProductionCompanies []struct {
		Id            int     `json:"id,omitempty"`
		LogoPath      *string `json:"logo_path,omitempty"`
		Name          string  `json:"name,omitempty"`
		OriginCountry string  `json:"origin_country,omitempty"`
	} `json:"production_companies,omitempty"`
	ProductionCountries []struct {
		Iso31661 string `json:"iso_3166_1,omitempty"`
		Name     string `json:"name,omitempty"`
	} `json:"production_countries,omitempty"`
	ReleaseDate     string `json:"release_date,omitempty"`
	Revenue         int    `json:"revenue,omitempty"`
	Runtime         int    `json:"runtime,omitempty"`
	SpokenLanguages []struct {
		EnglishName string `json:"english_name,omitempty"`
		Iso6391     string `json:"iso_639_1,omitempty"`
		Name        string `json:"name,omitempty"`
	} `json:"spoken_languages,omitempty"`
	Status      string  `json:"status,omitempty"`
	Tagline     string  `json:"tagline,omitempty"`
	Title       string  `json:"title,omitempty"`
	Video       bool    `json:"video,omitempty"`
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
			ProfilePath        *string `json:"profile_path,omitempty"`
			CastId             int     `json:"cast_id,omitempty"`
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
	ExternalIds struct {
		ImdbId      string      `json:"imdb_id,omitempty"`
		FacebookId  interface{} `json:"facebook_id,omitempty"`
		InstagramId interface{} `json:"instagram_id,omitempty"`
		TwitterId   interface{} `json:"twitter_id,omitempty"`
	} `json:"external_ids,omitempty"`
	Images struct {
		Backdrops []interface{} `json:"backdrops,omitempty"`
		Logos     []interface{} `json:"logos,omitempty"`
		Posters   []interface{} `json:"posters,omitempty"`
	} `json:"images,omitempty"`
	Keywords struct {
		Keywords []struct {
			Id   int    `json:"id,omitempty"`
			Name string `json:"name,omitempty"`
		} `json:"keywords,omitempty"`
	} `json:"keywords,omitempty"`
	Translations struct {
		Translations []struct {
			Iso31661    string `json:"iso_3166_1,omitempty"`
			Iso6391     string `json:"iso_639_1,omitempty"`
			Name        string `json:"name,omitempty"`
			EnglishName string `json:"english_name,omitempty"`
			Data        struct {
				Homepage string `json:"homepage,omitempty"`
				Overview string `json:"overview,omitempty"`
				Runtime  int    `json:"runtime,omitempty"`
				Tagline  string `json:"tagline,omitempty"`
				Title    string `json:"title,omitempty"`
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
