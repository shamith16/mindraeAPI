package https

import (
	"encoding/json"
	"github.com/shamith16/mindraeAPI/constants"
	"github.com/shamith16/mindraeAPI/entities/browse/tunefind"
	"io"
	"io/ioutil"
	"net/http"
)

func TuneFindMovieBrowse(tunefindMovieBrowse *tunefind.TuneFindMovieBrowse) (err error) {

	response, err := http.Get(constants.TuneFindBrowseAllMovie)
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
			err = json.Unmarshal(body, tunefindMovieBrowse)
			if err != nil {
				return err
			} else {
				return err
			}
		}

	}
}

func TuneFindShowBrowse(tunefindShowBrowse *tunefind.TuneFindShowBrowse) (err error) {

	response, err := http.Get(constants.TuneFindBrowseAllShow)
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
			err = json.Unmarshal(body, tunefindShowBrowse)
			if err != nil {
				return err
			} else {
				return err
			}
		}

	}
}

func TuneFindGameBrowse(tunefindGameBrowse *tunefind.TuneFindGameBrowse) (err error) {

	response, err := http.Get(constants.TuneFindBrowseAllGame)
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
			err = json.Unmarshal(body, tunefindGameBrowse)
			if err != nil {
				return err
			} else {
				return err
			}
		}

	}
}
