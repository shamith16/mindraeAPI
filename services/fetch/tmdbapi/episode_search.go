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

func EpisodeSearch(showID int, seasonNumber int, episodeNumber int) (tmdbEpisodeSearch *tmdbmodel.Episodes, err error) {

	link := fmt.Sprintf("%s%d/season/%d/episode/%d?%s%s",
		constants.TmdbShowURL, showID, seasonNumber, episodeNumber, constants.TmdbApiKey, constants.TmdbLanguage)

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
