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

func MovieSearchById(id int) (tmdbMovieSearch tmdbmodel.MovieSearch, err error) {

	link := fmt.Sprintf("%s%d?api_key=2caaa89866fe5b08fcab57571825d956&language=en-US&append_to_response=credits,external_ids,images,keywords,translations,videos&include_image_language=en,null",
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
