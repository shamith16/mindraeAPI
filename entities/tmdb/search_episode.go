package tmdb

type EpisodeSearch struct {
	AirDate string `json:"air_date"`
	//Crew    []struct {
	//	Department         string  `json:"department"`
	//	Job                string  `json:"job"`
	//	CreditId           string  `json:"credit_id"`
	//	Adult              bool    `json:"adult"`
	//	Gender             int     `json:"gender"`
	//	Id                 int     `json:"id"`
	//	KnownForDepartment string  `json:"known_for_department"`
	//	Name               string  `json:"name"`
	//	OriginalName       string  `json:"original_name"`
	//	Popularity         float64 `json:"popularity"`
	//	ProfilePath        *string `json:"profile_path"`
	//} `json:"crew"`
	EpisodeNumber int `json:"episode_number"`
	//GuestStars    []struct {
	//	Character          string  `json:"character"`
	//	CreditId           string  `json:"credit_id"`
	//	Order              int     `json:"order"`
	//	Adult              bool    `json:"adult"`
	//	Gender             int     `json:"gender"`
	//	Id                 int     `json:"id"`
	//	KnownForDepartment string  `json:"known_for_department"`
	//	Name               string  `json:"name"`
	//	OriginalName       string  `json:"original_name"`
	//	Popularity         float64 `json:"popularity"`
	//	ProfilePath        string  `json:"profile_path"`
	//} `json:"guest_stars"`
	Name           string  `json:"name"`
	Overview       string  `json:"overview"`
	Id             int     `json:"id"`
	ProductionCode string  `json:"production_code"`
	SeasonNumber   int     `json:"season_number"`
	StillPath      string  `json:"still_path"`
	VoteAverage    float64 `json:"vote_average"`
	VoteCount      int     `json:"vote_count"`
}
