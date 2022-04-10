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

//TmdbShowBrowse Search show By Name and returns a list of Fetched Results
func ShowSearch(name string) (tmdbShowBrowse tmdbmodel.ShowMovieBrowse, err error) {

	link := fmt.Sprintf("%s&query=%s",
		constants.TmdbShowBrowseURL, name)
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
