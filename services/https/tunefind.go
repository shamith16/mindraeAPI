package https

import (
	"encoding/json"
	"github.com/shamith16/mindraeAPI/constants"
	"github.com/shamith16/mindraeAPI/entities/tunefind"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

func TuneFindMovieBrowse(tunefindMovieBrowse *tunefind.MovieBrowse) (err error) {

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

func TuneFindShowBrowse(tunefindShowBrowse *tunefind.ShowBrowse) (err error) {

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

func TuneFindGameBrowse(tunefindGameBrowse *tunefind.GameBrowse) (err error) {

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

func TuneFindShowHome(tunefindShowHome *tunefind.ShowHome) (err error) {

	response, err := http.Get(constants.TuneFindShowHome)
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
			err = json.Unmarshal(body, tunefindShowHome)
			if err != nil {
				return err
			} else {
				return err
			}
		}

	}
}

func TuneFindMovieHome(tunefindMovieHome *tunefind.MovieGameHome) (err error) {

	response, err := http.Get(constants.TuneFindMovieHome)
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
			err = json.Unmarshal(body, tunefindMovieHome)
			if err != nil {
				return err
			} else {
				return err
			}
		}

	}
}

func TuneFindGameHome(tunefindGameHome *tunefind.MovieGameHome) (err error) {

	response, err := http.Get(constants.TuneFindGameHome)
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
			err = json.Unmarshal(body, tunefindGameHome)
			if err != nil {
				return err
			} else {
				return err
			}
		}

	}
}

func TuneFindMovieSearch(movieName string, tuneFindMovieSearch *tunefind.MovieSearch) (err error) {

	response, err := http.Get(constants.TuneFindMovieSearch + movieName + "?fields=song-events,hot-songs")
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
			err = json.Unmarshal(body, tuneFindMovieSearch)
			if err != nil {
				return err
			} else {
				return err
			}
		}

	}
}

func TuneFindShowSearch(showName string, tuneFindShowSearch *tunefind.ShowSearch) (err error) {

	response, err := http.Get(constants.TuneFindShowSearch + showName + "?fields=seasons&metatags=1")
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
			err = json.Unmarshal(body, tuneFindShowSearch)
			if err != nil {
				return err
			} else {
				return err
			}
		}

	}
}

func TuneFindSeasonSearch(showName string, seasonNumber int, tuneFindSeasonSearch *tunefind.SeasonSearch) (err error) {

	response, err := http.Get(constants.TuneFindSeasonSearch + showName + "/season/" + strconv.Itoa(seasonNumber) + "?fields=episodes,theme-song,hot-songs,albums&metatags=1")
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
			err = json.Unmarshal(body, tuneFindSeasonSearch)
			if err != nil {
				return err
			} else {
				return err
			}
		}

	}
}

func TuneFindEpisodeSearch(episodeID int, tuneFindEpisodeSearch *tunefind.EpisodeSearch) (err error) {

	response, err := http.Get(constants.TuneFindEpisodeSearch + strconv.Itoa(episodeID) + "?fields=song-events,questions,nextPrev")
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
			err = json.Unmarshal(body, tuneFindEpisodeSearch)
			if err != nil {
				return err
			} else {
				return err
			}
		}

	}
}

func TuneFindTrending(tunefindTrending *tunefind.TrendingSongs) (err error) {

	response, err := http.Get(constants.TuneFindTrending)
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
			err = json.Unmarshal(body, tunefindTrending)
			if err != nil {
				return err
			} else {
				return err
			}
		}

	}
}
