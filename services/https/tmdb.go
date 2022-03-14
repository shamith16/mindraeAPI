package https

import (
	"encoding/json"
	"github.com/shamith16/mindraeAPI/constants"
	"github.com/shamith16/mindraeAPI/entities/browse/tmdb"
	"io"
	"io/ioutil"
	"net/http"
)

func TmdbMovieBrowse(name string, year string, tmdbMovieBrowse *tmdb.TmdbMovieBrowse) (err error) {

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

func TmdbShowBrowse(name string, tmdbShowBrowse *tmdb.TmdbShowBrowse) (err error) {

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
