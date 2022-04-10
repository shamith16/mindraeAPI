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

func SeasonSearch(showID int, seasonNumber int) (tmdbSeasonSearch *tmdbmodel.SeasonSearch, err error) {

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
