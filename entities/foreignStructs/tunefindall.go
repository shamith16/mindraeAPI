package foreignStructs

import "github.com/shamith16/mindraeAPI/entities/tunefind"

type CustomBrowseAll struct {
	TunefindMovieBrowseAll tunefind.MovieBrowse   `json:"tunefindMovieBrowseAll"`
	TunefindShowBrowseAll  tunefind.ShowBrowse    `json:"tunefindShowBrowseAll"`
	TunefindGameBrowseAll  tunefind.GameBrowse    `json:"tunefindGameBrowseAll"`
	TunefindSearchMovie    tunefind.MovieSearch   `json:"tunefindMovieSearch"`
	TunefindSearchShow     tunefind.ShowSearch    `json:"tunefindShowSearch"`
	TunefindSearchSeason   tunefind.SeasonSearch  `json:"tunefindSeasonSearch"`
	TunefindSearchEpisode  tunefind.EpisodeSearch `json:"tunefindEpisodeSearch"`
	TunefindMovieHome      tunefind.MovieGameHome `json:"tunefindMovieHome"`
	TunefindShowHome       tunefind.ShowHome      `json:"tunefindShowHome"`
	TunefindGameHome       tunefind.MovieGameHome `json:"tunefindGameHome"`
	TunefindTrending       tunefind.TrendingSongs `json:"tunefindTrending"`
}
