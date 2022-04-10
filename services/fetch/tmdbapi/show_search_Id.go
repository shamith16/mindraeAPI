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

func ShowSearchById(showID int) (tmdbShowSearch *tmdbmodel.ShowSearch, err error) {

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
