package constants

const (
	TuneFindApiBaseURL = "https://www.tunefind.com/api/frontend/"
	TunefindBaseURL    = "https://tunefind.com"
	tmdbApiKey         = "&api_key=2caaa89866fe5b08fcab57571825d956"
	TmdbImageBaseUrl   = "https://image.tmdb.org/t/p/"
	TmdbShowSearch     = "https://api.themoviedb.org/3/search/tv?language=en-US&page=1&include_adult=true" + tmdbApiKey
	TmdbMovieSearch    = "https://api.themoviedb.org/3/search/movie?language=en-US&page=1&include_adult=true" + tmdbApiKey
)
