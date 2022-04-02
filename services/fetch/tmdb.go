package fetch

import (
	"encoding/json"
	"fmt"
	"github.com/shamith16/mindraeAPI/constants"
	"github.com/shamith16/mindraeAPI/entities/tmdb"
	"io"
	"io/ioutil"
	"net/http"
)

//TmdbShowBrowse Search Movie By Name and select Relevant Movie Object from a list of Fetched Results
func TmdbShowBrowse(name string) (tmdbShowBrowse tmdb.ShowMovieBrowse, err error) {
	link := fmt.Sprintf("%s&query=%s", constants.TmdbShowBrowseURL, name)
	response, err := http.Get(link)
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)
	if err != nil {
		return tmdbShowBrowse, err
	} else {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return tmdbShowBrowse, err
		} else {
			err = json.Unmarshal(body, &tmdbShowBrowse)
			if err != nil {
				return tmdbShowBrowse, err
			} else {
				return tmdbShowBrowse, err
			}
		}

	}
}

func TmdbMovieBrowse(name string, year string) (tmdbMovieBrowse tmdb.ShowMovieBrowse, err error) {
	link := fmt.Sprintf("%s&query=%s&year=%s&primary_release_year=%s",
		constants.TmdbMovieBrowseURL, name, year, year)
	response, err := http.Get(link)

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)
	if err != nil {
		return tmdbMovieBrowse, err
	} else {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return tmdbMovieBrowse, err
		} else {
			err = json.Unmarshal(body, &tmdbMovieBrowse)
			if err != nil {
				return tmdbMovieBrowse, err
			} else {
				return tmdbMovieBrowse, err
			}
		}

	}
}

func TmdbMovieSearch(id int) (tmdbMovieSearch tmdb.MovieSearch, err error) {
	link := fmt.Sprintf("%s%d?api_key=2caaa89866fe5b08fcab57571825d956&language=en-US&append_to_response=credits,external_ids,images,keywords,translations,videos",
		constants.TmdbMovieSearchURL, id)
	response, err := http.Get(link)
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)
	if err != nil {
		return tmdbMovieSearch, err
	} else {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return tmdbMovieSearch, err
		} else {
			err = json.Unmarshal(body, &tmdbMovieSearch)
			if err != nil {
				return tmdbMovieSearch, err
			} else {
				return tmdbMovieSearch, err
			}
		}

	}
}

func TmdbShowSearch(showID int) (tmdbShowSearch *tmdb.ShowSearch, err error) {
	link := fmt.Sprintf("%s%d?api_key=2caaa89866fe5b08fcab57571825d956&language=en-US&append_to_response=credits,episode_groups,external_ids,images,keywords,translations,videos",
		constants.TmdbShowURL, showID)
	response, err := http.Get(link)
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)
	if err != nil {
		return tmdbShowSearch, err
	} else {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return tmdbShowSearch, err
		} else {
			err = json.Unmarshal(body, tmdbShowSearch)
			if err != nil {
				return tmdbShowSearch, err
			} else {
				return tmdbShowSearch, err
			}
		}

	}
}

func TmdbSeasonSearch(showID int, seasonNumber int) (tmdbSeasonSearch *tmdb.SeasonSearch, err error) {
	link := fmt.Sprintf("%s%d/season/%d?%s%s",
		constants.TmdbShowURL, showID, seasonNumber, constants.TmdbApiKey, constants.TmdbLanguage)
	response, err := http.Get(link)
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)
	if err != nil {
		return tmdbSeasonSearch, err
	} else {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return tmdbSeasonSearch, err
		} else {
			err = json.Unmarshal(body, tmdbSeasonSearch)
			if err != nil {
				return tmdbSeasonSearch, err
			} else {
				return tmdbSeasonSearch, err
			}
		}

	}
}

func TmdbEpisodeSearch(showID int, seasonNumber int, episodeNumber int) (tmdbEpisodeSearch *tmdb.Episodes, err error) {
	link := fmt.Sprintf("%s%d/season/%d/episode/%d?%s%s", constants.TmdbShowURL, showID, seasonNumber, episodeNumber, constants.TmdbApiKey, constants.TmdbLanguage)
	response, err := http.Get(link)
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)
	if err != nil {
		return tmdbEpisodeSearch, err
	} else {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return tmdbEpisodeSearch, err
		} else {
			err = json.Unmarshal(body, tmdbEpisodeSearch)
			if err != nil {
				return tmdbEpisodeSearch, err
			} else {
				return tmdbEpisodeSearch, err
			}
		}

	}
}
