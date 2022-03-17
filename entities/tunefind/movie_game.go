package tunefind

type MovieGameHome struct {
	Slider        []string     `json:"slider,omitempty"`
	Featured      []string     `json:"featured,omitempty"`
	RecentlyAdded []string     `json:"recently_added,omitempty"`
	Contribute    []string     `json:"contribute,omitempty"`
	Games         []Games      `json:"games,omitempty"`
	Movies        []ShowMovies `json:"movies,omitempty"`
}
