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

func EpisodeSearch(episodeID int) (tuneFindEpisodeSearch tunefindmodel.EpisodeSearch, err error) {

	link := fmt.Sprintf("%s%s?fields=song-events,questions,nextPrev",
		constants.TuneFindEpisodeSearch, strconv.Itoa(episodeID))

	response, err := http.Get(link)

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	if err != nil {
		return tuneFindEpisodeSearch, err
	} else {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return tuneFindEpisodeSearch, err
		} else {
			err = json.Unmarshal(body, &tuneFindEpisodeSearch)
			if err != nil {
				return tuneFindEpisodeSearch, err
			} else {
				return tuneFindEpisodeSearch, err
			}
		}

	}
}
