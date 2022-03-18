package https

import (
	"encoding/json"
	"fmt"
	"github.com/shamith16/mindraeAPI/constants"
	"github.com/shamith16/mindraeAPI/entities/tmdb"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

func TmdbMovieBrowse(name string, year int) (tmdbMovieBrowse tmdb.ShowMovieBrowse, err error) {
	link := fmt.Sprintf("%s&query=%s&year=%d&primary_release_year=%d", constants.TmdbMovieBrowseURL, name, year, year)
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

func TmdbShowBrowse(name string) (tmdbShowBrowse tmdb.ShowMovieBrowse, err error) {

	response, err := http.Get(constants.TmdbShowBrowseURL + "&query=" + name)
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

func TmdbMovieSearch(id int) (tmdbMovieSearch tmdb.MovieSearch, err error) {

	response, err := http.Get(constants.TmdbMovieSearchURL + strconv.Itoa(id) + "?api_key=2caaa89866fe5b08fcab57571825d956&language=en-US&append_to_response=credits,external_ids,images,keywords,translations,videos")
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

	response, err := http.Get(constants.TmdbShowURL + strconv.Itoa(showID) + "?api_key=2caaa89866fe5b08fcab57571825d956&language=en-US&append_to_response=credits,episode_groups,external_ids,images,keywords,translations,videos,")
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

	response, err := http.Get(constants.TmdbShowURL + strconv.Itoa(showID) + "/season/" + strconv.Itoa(seasonNumber) + "?api_key=2caaa89866fe5b08fcab57571825d956&language=en-US")
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

	response, err := http.Get(constants.TmdbShowURL + strconv.Itoa(showID) + "/season/" + strconv.Itoa(seasonNumber) + "/episode/" + strconv.Itoa(episodeNumber) + "?api_key=2caaa89866fe5b08fcab57571825d956&language=en-US")
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
