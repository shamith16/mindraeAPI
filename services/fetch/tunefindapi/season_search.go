package tunefindapi

import (
	"encoding/json"
	"fmt"
	"github.com/shamith16/mindraeAPI/constants"
	"github.com/shamith16/mindraeAPI/entities/tunefindmodel"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

func SeasonSearch(showName string, seasonNumber int) (tuneFindSeasonSearch tunefindmodel.SeasonSearch, err error) {

	link := fmt.Sprintf("%s%s/season/%s?fields=episodes,theme-song,hot-songs,albums&metatags=1",
		constants.TuneFindSeasonSearch, showName, strconv.Itoa(seasonNumber))

	response, err := http.Get(link)

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	if err != nil {
		return tuneFindSeasonSearch, err
	} else {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return tuneFindSeasonSearch, err
		} else {
			err = json.Unmarshal(body, &tuneFindSeasonSearch)
			if err != nil {
				return tuneFindSeasonSearch, err
			} else {
				return tuneFindSeasonSearch, err
			}
		}

	}
}
