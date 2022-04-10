package tunefindmodel

type TrendingSongs struct {
	Songs []Songs `json:"songs"`
}

type Events struct {
	Rank  int   `json:"rank"`
	Event Event `json:"event"`
}

type Songs struct {
	Rank   int      `json:"rank"`
	Trend  string   `json:"trend"`
	Song   Song     `json:"song"`
	Events []Events `json:"events"`
}
