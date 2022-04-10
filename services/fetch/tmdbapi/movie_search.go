package tmdbapi

import (
	"encoding/json"
	"fmt"
	"github.com/shamith16/mindraeAPI/constants"
	"github.com/shamith16/mindraeAPI/entities/tmdbmodel"
	"io"
	"io/ioutil"
	"net/http"
)

func MovieSearch(name string, year string) (tmdbMovieBrowse tmdbmodel.ShowMovieBrowse, err error) {

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
