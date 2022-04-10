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

func ShowHome() (tunefindShowHome tunefindmodel.ShowHome, err error) {

	link := fmt.Sprintf("%s", constants.TuneFindShowHome)

	response, err := http.Get(link)

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	if err != nil {
		return tunefindShowHome, err
	} else {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return tunefindShowHome, err
		} else {
			err = json.Unmarshal(body, &tunefindShowHome)
			if err != nil {
				return tunefindShowHome, err
			} else {
				return tunefindShowHome, err
			}
		}

	}
}
