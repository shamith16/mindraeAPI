package browse

type TmdbMovieBrowse struct {
	Page    int `json:"page,omitempty"`
	Results []struct {
		Adult            bool    `json:"adult,omitempty"`
		BackdropPath     *string `json:"backdrop_path,omitempty"`
		GenreIds         []int   `json:"genre_ids,omitempty"`
		Id               int     `json:"id,omitempty"`
		OriginalLanguage string  `json:"original_language,omitempty"`
		OriginalTitle    string  `json:"original_title,omitempty"`
		Overview         string  `json:"overview,omitempty"`
		Popularity       float64 `json:"popularity,omitempty"`
		PosterPath       *string `json:"poster_path,omitempty"`
		ReleaseDate      string  `json:"release_date,omitempty"`
		Title            string  `json:"title,omitempty"`
		Video            bool    `json:"video,omitempty"`
		VoteAverage      float64 `json:"vote_average,omitempty"`
		VoteCount        int     `json:"vote_count,omitempty"`
	} `json:"results,omitempty"`
	TotalPages   int `json:"total_pages,omitempty"`
	TotalResults int `json:"total_results,omitempty"`
}
