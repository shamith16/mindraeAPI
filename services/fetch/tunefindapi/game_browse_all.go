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

func GameBrowseAll() (tunefindGameBrowse tunefindmodel.GameBrowse, err error) {

	link := fmt.Sprintf("%s",
		constants.TuneFindBrowseAllGame)

	response, err := http.Get(link)

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	if err != nil {
		return tunefindGameBrowse, err
	} else {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return tunefindGameBrowse, err
		} else {
			err = json.Unmarshal(body, &tunefindGameBrowse)
			if err != nil {
				return tunefindGameBrowse, err
			} else {
				return tunefindGameBrowse, err
			}
		}

	}
}
