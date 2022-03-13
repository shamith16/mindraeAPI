package search

type TmdbSeasonSearch struct {
	Id       string `json:"_id,omitempty"`
	AirDate  string `json:"air_date,omitempty"`
	Episodes []struct {
		AirDate       string `json:"air_date,omitempty"`
		EpisodeNumber int    `json:"episode_number,omitempty"`
		Crew          []struct {
			Department         string  `json:"department,omitempty"`
			Job                string  `json:"job,omitempty"`
			CreditId           string  `json:"credit_id,omitempty"`
			Adult              bool    `json:"adult,omitempty"`
			Gender             int     `json:"gender,omitempty"`
			Id                 int     `json:"id,omitempty"`
			KnownForDepartment string  `json:"known_for_department,omitempty"`
			Name               string  `json:"name,omitempty"`
			OriginalName       string  `json:"original_name,omitempty"`
			Popularity         float64 `json:"popularity,omitempty"`
			ProfilePath        *string `json:"profile_path,omitempty"`
		} `json:"crew,omitempty"`
		GuestStars []struct {
			Character          string  `json:"character,omitempty"`
			CreditId           string  `json:"credit_id,omitempty"`
			Order              int     `json:"order,omitempty"`
			Adult              bool    `json:"adult,omitempty"`
			Gender             int     `json:"gender,omitempty"`
			Id                 int     `json:"id,omitempty"`
			KnownForDepartment string  `json:"known_for_department,omitempty"`
			Name               string  `json:"name,omitempty"`
			OriginalName       string  `json:"original_name,omitempty"`
			Popularity         float64 `json:"popularity,omitempty"`
			ProfilePath        *string `json:"profile_path,omitempty"`
		} `json:"guest_stars,omitempty"`
		Id             int     `json:"id,omitempty"`
		Name           string  `json:"name,omitempty"`
		Overview       string  `json:"overview,omitempty"`
		ProductionCode string  `json:"production_code,omitempty"`
		SeasonNumber   int     `json:"season_number,omitempty"`
		StillPath      string  `json:"still_path,omitempty"`
		VoteAverage    float64 `json:"vote_average,omitempty"`
		VoteCount      int     `json:"vote_count,omitempty"`
	} `json:"episodes,omitempty"`
	Name         string `json:"name,omitempty"`
	Overview     string `json:"overview,omitempty"`
	Id1          int    `json:"id,omitempty"`
	PosterPath   string `json:"poster_path,omitempty"`
	SeasonNumber int    `json:"season_number,omitempty"`
}
