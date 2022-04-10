package tunefindmodel

type ShowHome struct {
	Shows        []ShowMovies `json:"shows"`
	AiringShows  []string     `json:"airing_shows"`
	NewShows     []string     `json:"new_shows"`
	PopularShows []string     `json:"popular_shows"`
}
