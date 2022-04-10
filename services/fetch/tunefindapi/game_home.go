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

func GameHome() (tunefindGameHome tunefindmodel.MovieGameHome, err error) {

	link := fmt.Sprintf("%s",
		constants.TuneFindGameHome)

	response, err := http.Get(link)

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	if err != nil {
		return tunefindGameHome, err
	} else {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return tunefindGameHome, err
		} else {
			err = json.Unmarshal(body, &tunefindGameHome)
			if err != nil {
				return tunefindGameHome, err
			} else {
				return tunefindGameHome, err
			}
		}

	}
}
