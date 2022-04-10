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

func MovieBrowseAll() (tunefindMovieBrowse tunefindmodel.MovieBrowse, err error) {

	link := fmt.Sprintf("%s",
		constants.TuneFindBrowseAllMovie)

	response, err := http.Get(link)

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	if err != nil {
		return tunefindMovieBrowse, err
	} else {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return tunefindMovieBrowse, err
		} else {
			err = json.Unmarshal(body, &tunefindMovieBrowse)
			if err != nil {
				return tunefindMovieBrowse, err
			} else {
				return tunefindMovieBrowse, err
			}
		}

	}
}
