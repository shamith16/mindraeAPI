package tmdbapi

import (
	"fmt"
	tmdb "github.com/cyruzin/golang-tmdb"
	"github.com/shamith16/mindraeAPI/constants"
	"net/http"
	"time"
)

var tmdbClient *tmdb.Client
var err error

func init() {
	tmdbClient, err = tmdb.Init(constants.TmdbApiKey)
	if err != nil {
		fmt.Println(err)
	}

	h := http.Client{
		Timeout: 50 * time.Second,
	}
	tmdbClient.SetClientConfig(h)
}
func GetMovieDetails(movieId int) (tmdbMovieDetails tmdb.MovieDetails, err error) {
	options := map[string]string{
		"include_video_language": "en",
		"include_image_language": "en",
		//alternative_titles,changes,external_ids,keywords,lists,recommendations,reviews,similar,translations,
		"append_to_response": "credits,images,release_dates,videos,watch/providers",
	}

	tmdbMovie, err := tmdbClient.GetMovieDetails(movieId, options)
	if err != nil {
		fmt.Println(err)
	}
	tmdbMovieDetails = *tmdbMovie
	return
}

func GetTvDetails(showID int) (tmdbTvDetails tmdb.TVDetails, err error) {
	tmdbTv, err := tmdbClient.GetTVDetails(showID, nil)
	tmdbTvDetails = *tmdbTv
	return
}

func GetTvSeasonDetails(showID int, seasonNumber int) (tmdbTvSeasonDetails tmdb.TVSeasonDetails, err error) {
	tmdbTvSeason, err := tmdbClient.GetTVSeasonDetails(showID, seasonNumber, nil)
	tmdbTvSeasonDetails = *tmdbTvSeason
	return
}

func GetTvEpisodeDetails(showID int, seasonNumber int, episodeNumber int) (tmdbTvEpisodeDetails tmdb.TVEpisodeDetails, err error) {
	tmdbEpisode, err := tmdbClient.GetTVEpisodeDetails(showID, seasonNumber, episodeNumber, nil)
	tmdbTvEpisodeDetails = *tmdbEpisode
	return
}

func SearchTmdbMovie(movieName string, urlOptions map[string]string) (tmdbSearchMovies tmdb.SearchMovies, err error) {
	movieSearch, err := tmdbClient.GetSearchMovies(movieName, urlOptions)
	tmdbSearchMovies = *movieSearch
	return
}

func SearchTmdbTvShows(showName string) (tmdbSearchTVShows tmdb.SearchTVShows, err error) {
	tvShowSearch, err := tmdbClient.GetSearchTVShow(showName, nil)
	tmdbSearchTVShows = *tvShowSearch
	return
}
