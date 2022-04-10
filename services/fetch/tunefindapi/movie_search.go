package tunefindapi

import (
	"encoding/json"
	"fmt"
	"github.com/shamith16/mindraeAPI/constants"
	"github.com/shamith16/mindraeAPI/entities/tunefindmodel"
	"io"
	"io/ioutil"
	"net/http"
)

func MovieSearch(movieName string) (tuneFindMovieSearch tunefindmodel.MovieSearch, err error) {

	link := fmt.Sprintf("%s%s?fields=song-events,hot-songs",
		constants.TuneFindMovieSearch, movieName)

	response, err := http.Get(link)

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	if err != nil {
		return tuneFindMovieSearch, err
	} else {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return tuneFindMovieSearch, err
		} else {
			err = json.Unmarshal(body, &tuneFindMovieSearch)
			if err != nil {
				return tuneFindMovieSearch, err
			} else {
				return tuneFindMovieSearch, err
			}
		}

	}
}
