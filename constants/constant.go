package constants

import "fmt"

var (
	TmdbLanguage           = "&language=en-US"
	tmdbAdult              = "&include_adult=true"
	tmdbRegion, tmdbPage   = "&region=US", "&page=1"
	tuneFindApiBaseURL     = "https://www.tunefind.com/api/frontend"
	TunefindBaseURL        = "https://tunefind.com"
	tmdbAPIBaseURL         = "https://api.themoviedb.org/3"
	TmdbApiKey             = "api_key=2caaa89866fe5b08fcab57571825d956"
	tmdbImageBaseUrl       = "https://image.tmdb.org/t/p/"
	TmdbShowBrowseURL      = fmt.Sprintf(tmdbAPIBaseURL+"/search/tv?%s%s%s%s", TmdbApiKey, TmdbLanguage, tmdbPage, tmdbAdult)
	TmdbMovieBrowseURL     = fmt.Sprintf(tmdbAPIBaseURL+"/search/movie?%s%s%s%s%s", TmdbApiKey, TmdbLanguage, tmdbPage, tmdbAdult, tmdbRegion)
	TmdbMovieSearchURL     = tmdbAPIBaseURL + "/movie/"
	TmdbShowURL            = tmdbAPIBaseURL + "/tv/"
	TuneFindBrowseAllMovie = tuneFindApiBaseURL + "/movie/browse"
	TuneFindBrowseAllShow  = tuneFindApiBaseURL + "/show/browse"
	TuneFindBrowseAllGame  = tuneFindApiBaseURL + "/game/browse"
	TuneFindShowHome       = tuneFindApiBaseURL + "/show"
	TuneFindMovieHome      = tuneFindApiBaseURL + "/movie"
	TuneFindGameHome       = tuneFindApiBaseURL + "/game"
	TuneFindMovieSearch    = tuneFindApiBaseURL + "/movie/"
	TuneFindShowSearch     = tuneFindApiBaseURL + "/show/"
	TuneFindSeasonSearch   = tuneFindApiBaseURL + "/show/"
	TuneFindEpisodeSearch  = tuneFindApiBaseURL + "/episode/"
	TuneFindTrending       = tuneFindApiBaseURL + "/trending/now"
)
