package browse

type TmdbShowBrowse struct {
	Page    int `json:"page,omitempty"`
	Results []struct {
		BackdropPath     *string  `json:"backdrop_path,omitempty"`
		FirstAirDate     string   `json:"first_air_date,omitempty"`
		GenreIds         []int    `json:"genre_ids,omitempty"`
		Id               int      `json:"id,omitempty"`
		Name             string   `json:"name,omitempty"`
		OriginCountry    []string `json:"origin_country,omitempty"`
		OriginalLanguage string   `json:"original_language,omitempty"`
		OriginalName     string   `json:"original_name,omitempty"`
		Overview         string   `json:"overview,omitempty"`
		Popularity       float64  `json:"popularity,omitempty"`
		PosterPath       *string  `json:"poster_path,omitempty"`
		VoteAverage      float64  `json:"vote_average,omitempty"`
		VoteCount        int      `json:"vote_count,omitempty"`
	} `json:"results,omitempty"`
	TotalPages   int `json:"total_pages,omitempty"`
	TotalResults int `json:"total_results,omitempty"`
}
