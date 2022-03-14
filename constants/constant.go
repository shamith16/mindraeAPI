package constants

const (
	TuneFindApiBaseURL     = "https://www.tunefind.com/api/frontend/"
	TunefindBaseURL        = "https://tunefind.com"
	tmdbApiKey             = "api_key=2caaa89866fe5b08fcab57571825d956"
	TmdbImageBaseUrl       = "https://image.tmdb.org/t/p/"
	TmdbShowBrowse         = "https://api.themoviedb.org/3/search/tv?" + tmdbApiKey + "&language=en-US&page=1&include_adult=true"
	TmdbMovieBrowse        = "https://api.themoviedb.org/3/search/movie?" + tmdbApiKey + "&language=en-US&page=1&include_adult=true"
	TuneFindBrowseAllMovie = "https://www.tunefind.com/api/frontend/movie/browse"
	TuneFindBrowseAllShow  = "https://www.tunefind.com/api/frontend/show/browse"
	TuneFindBrowseAllGame  = "https://www.tunefind.com/api/frontend/game/browse"
)
