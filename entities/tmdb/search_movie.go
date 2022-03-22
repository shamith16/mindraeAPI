package tmdb

type MovieSearch struct {
	Adult        bool   `json:"adult"`
	BackdropPath string `json:"backdrop_path"`
	//BelongsToCollection interface{} `json:"belongs_to_collection"`
	Budget           int      `json:"budget"`
	Genres           []Genres `json:"genres"`
	Homepage         string   `json:"homepage"`
	Id               int      `json:"id"`
	ImdbId           string   `json:"imdb_id"`
	OriginalLanguage string   `json:"original_language"`
	OriginalTitle    string   `json:"original_title"`
	Overview         string   `json:"overview"`
	Popularity       float64  `json:"popularity"`
	PosterPath       string   `json:"poster_path"`
	//ProductionCompanies []struct {
	//	Id            int     `json:"id"`
	//	LogoPath      *string `json:"logo_path"`
	//	Name          string  `json:"name"`
	//	OriginCountry string  `json:"origin_country"`
	//} `json:"production_companies,-"`
	//ProductionCountries []struct {
	//	Iso31661 string `json:"iso_3166_1"`
	//	Name     string `json:"name"`
	//} `json:"production_countries"`
	ReleaseDate     string            `json:"release_date"`
	Revenue         int               `json:"revenue"`
	Runtime         int               `json:"runtime"`
	SpokenLanguages []SpokenLanguages `json:"spoken_languages"`
	Status          string            `json:"status"`
	Tagline         string            `json:"tagline"`
	Title           string            `json:"title"`
	Video           bool              `json:"video"`
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
	ExternalIds ExternalIds `json:"external_ids"`
	Images      Images      `json:"images"`
	//Keywords struct {
	//	Keywords []struct {
	//		Id   int    `json:"id"`
	//		Name string `json:"name"`
	//	} `json:"keywords"`
	//} `json:"keywords"`
	//Translations struct {
	//	Translations []struct {
	//		Iso31661    string `json:"iso_3166_1"`
	//		Iso6391     string `json:"iso_639_1"`
	//		Name        string `json:"name"`
	//		EnglishName string `json:"english_name"`
	//		Data        struct {
	//			Homepage string `json:"homepage"`
	//			Overview string `json:"overview"`
	//			Runtime  int    `json:"runtime"`
	//			Tagline  string `json:"tagline"`
	//			Title    string `json:"title"`
	//		} `json:"data"`
	//	} `json:"translations"`
	//} `json:"translations"`
	Videos Videos `json:"videos"`
}

type Videos struct {
	Results []VideoResults `json:"results"`
}
