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

func MovieHome() (tunefindMovieHome tunefindmodel.MovieGameHome, err error) {

	link := fmt.Sprintf("%s",
		constants.TuneFindMovieHome)

	response, err := http.Get(link)

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	if err != nil {
		return tunefindMovieHome, err
	} else {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return tunefindMovieHome, err
		} else {
			err = json.Unmarshal(body, &tunefindMovieHome)
			if err != nil {
				return tunefindMovieHome, err
			} else {
				return tunefindMovieHome, err
			}
		}

	}
}
