package tmdb

type SeasonSearch struct {
	Id           string     `json:"_id"`
	AirDate      string     `json:"air_date"`
	Episodes     []Episodes `json:"episodes"`
	Name         string     `json:"name"`
	Overview     string     `json:"overview"`
	Id1          int        `json:"id"`
	PosterPath   string     `json:"poster_path"`
	SeasonNumber int        `json:"season_number"`
}
