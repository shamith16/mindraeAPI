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

func ShowSearch(showName string) (tuneFindShowSearch tunefindmodel.ShowSearch, err error) {

	link := fmt.Sprintf("%s%s?fields=seasons&metatags=1", constants.TuneFindShowSearch, showName)

	response, err := http.Get(link)

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	if err != nil {
		return tuneFindShowSearch, err
	} else {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return tuneFindShowSearch, err
		} else {
			err = json.Unmarshal(body, &tuneFindShowSearch)
			if err != nil {
				return tuneFindShowSearch, err
			} else {
				return tuneFindShowSearch, err
			}
		}

	}
}
