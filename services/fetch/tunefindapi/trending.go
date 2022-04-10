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

func Trending() (tunefindTrending tunefindmodel.TrendingSongs, err error) {

	link := fmt.Sprintf("%s", constants.TuneFindTrending)

	response, err := http.Get(link)

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	if err != nil {
		return tunefindTrending, err
	} else {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return tunefindTrending, err
		} else {
			err = json.Unmarshal(body, &tunefindTrending)
			if err != nil {
				return tunefindTrending, err
			} else {
				return tunefindTrending, err
			}
		}

	}
}
