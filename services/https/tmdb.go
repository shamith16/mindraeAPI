package https

import (
	"encoding/json"
	"github.com/shamith16/mindraeAPI/constants"
	"github.com/shamith16/mindraeAPI/entities/tmdb"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

func TmdbMovieBrowse(name string, year string, tmdbMovieBrowse *tmdb.MovieBrowse) (err error) {

	response, err := http.Get(constants.TmdbMovieBrowse + "&query=" + name + "&year=" + year + "&primary_release_year=" + year)
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)
	if err != nil {
		return err
	} else {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		} else {
			err = json.Unmarshal(body, tmdbMovieBrowse)
			if err != nil {
				return err
			} else {
				return err
			}
		}

	}
}

func TmdbShowBrowse(name string, tmdbShowBrowse *tmdb.ShowBrowse) (err error) {

	response, err := http.Get(constants.TmdbShowBrowse + "&query=" + name)
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)
	if err != nil {
		return err
	} else {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		} else {
			err = json.Unmarshal(body, tmdbShowBrowse)
			if err != nil {
				return err
			} else {
				return err
			}
		}

	}
}

func TmdbMovieSearch(id int, tmdbMovieSearch *tmdb.MovieSearch) (err error) {

	response, err := http.Get(constants.TmdbMovieSearch + strconv.Itoa(id) + "?api_key=2caaa89866fe5b08fcab57571825d956&language=en-US&append_to_response=credits,external_ids,images,keywords,translations,videos")
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)
	if err != nil {
		return err
	} else {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		} else {
			err = json.Unmarshal(body, tmdbMovieSearch)
			if err != nil {
				return err
			} else {
				return err
			}
		}

	}
}

func TmdbShowSearch(showID int, tmdbShowSearch *tmdb.ShowSearch) (err error) {

	response, err := http.Get(constants.TmdbShowSearch + strconv.Itoa(showID) + "?api_key=2caaa89866fe5b08fcab57571825d956&language=en-US&append_to_response=credits,episode_groups,external_ids,images,keywords,translations,videos,")
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)
	if err != nil {
		return err
	} else {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		} else {
			err = json.Unmarshal(body, tmdbShowSearch)
			if err != nil {
				return err
			} else {
				return err
			}
		}

	}
}

func TmdbSeasonSearch(showID int, seasonNumber int, tmdbSeasonSearch *tmdb.SeasonSearch) (err error) {

	response, err := http.Get(constants.TmdbSeasonSearch + strconv.Itoa(showID) + "/season/" + strconv.Itoa(seasonNumber) + "?api_key=2caaa89866fe5b08fcab57571825d956&language=en-US")
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)
	if err != nil {
		return err
	} else {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		} else {
			err = json.Unmarshal(body, tmdbSeasonSearch)
			if err != nil {
				return err
			} else {
				return err
			}
		}

	}
}

func TmdbEpisodeSearch(showID int, seasonNumber int, episodeNumber int, tmdbEpisodeSearch *tmdb.EpisodeSearch) (err error) {

	response, err := http.Get(constants.TmdbEpisodeSearch + strconv.Itoa(showID) + "/season/" + strconv.Itoa(seasonNumber) + "/episode/" + strconv.Itoa(episodeNumber) + "?api_key=2caaa89866fe5b08fcab57571825d956&language=en-US")
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)
	if err != nil {
		return err
	} else {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		} else {
			err = json.Unmarshal(body, tmdbEpisodeSearch)
			if err != nil {
				return err
			} else {
				return err
			}
		}

	}
}
