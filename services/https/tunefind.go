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

func TuneFindMovieBrowse() (tunefindMovieBrowse tunefind.MovieBrowse, err error) {

	response, err := http.Get(constants.TuneFindBrowseAllMovie)
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

func TuneFindShowBrowse() (tunefindShowBrowse tunefind.ShowBrowse, err error) {

	response, err := http.Get(constants.TuneFindBrowseAllShow)
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)
	if err != nil {
		return tunefindShowBrowse, err
	} else {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return tunefindShowBrowse, err
		} else {
			err = json.Unmarshal(body, &tunefindShowBrowse)
			if err != nil {
				return tunefindShowBrowse, err
			} else {
				return tunefindShowBrowse, err
			}
		}

	}
}

func TuneFindGameBrowse() (tunefindGameBrowse tunefind.GameBrowse, err error) {

	response, err := http.Get(constants.TuneFindBrowseAllGame)
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

func TuneFindShowHome() (tunefindShowHome tunefind.ShowHome, err error) {

	response, err := http.Get(constants.TuneFindShowHome)
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

func TuneFindMovieHome() (tunefindMovieHome tunefind.MovieGameHome, err error) {

	response, err := http.Get(constants.TuneFindMovieHome)
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)
	if err != nil {
		return tunefindMovieHome, err
	} else {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return tunefindMovieHome, err
		} else {
			err = json.Unmarshal(body, &tunefindMovieHome)
			if err != nil {
				return tunefindMovieHome, err
			} else {
				return tunefindMovieHome, err
			}
		}

	}
}

func TuneFindGameHome() (tunefindGameHome tunefind.MovieGameHome, err error) {

	response, err := http.Get(constants.TuneFindGameHome)
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

func TuneFindMovieSearch(movieName string) (tuneFindMovieSearch tunefind.MovieSearch, err error) {

	response, err := http.Get(constants.TuneFindMovieSearch + movieName + "?fields=song-events,hot-songs")
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)
	if err != nil {
		return tuneFindMovieSearch, err
	} else {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return tuneFindMovieSearch, err
		} else {
			err = json.Unmarshal(body, &tuneFindMovieSearch)
			if err != nil {
				return tuneFindMovieSearch, err
			} else {
				return tuneFindMovieSearch, err
			}
		}

	}
}

func TuneFindShowSearch(showName string) (tuneFindShowSearch tunefind.ShowSearch, err error) {

	response, err := http.Get(constants.TuneFindShowSearch + showName + "?fields=seasons&metatags=1")
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

func TuneFindSeasonSearch(showName string, seasonNumber int) (tuneFindSeasonSearch tunefind.SeasonSearch, err error) {

	response, err := http.Get(constants.TuneFindSeasonSearch + showName + "/season/" + strconv.Itoa(seasonNumber) + "?fields=episodes,theme-song,hot-songs,albums&metatags=1")
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

func TuneFindEpisodeSearch(episodeID int) (tuneFindEpisodeSearch tunefind.EpisodeSearch, err error) {

	response, err := http.Get(constants.TuneFindEpisodeSearch + strconv.Itoa(episodeID) + "?fields=song-events,questions,nextPrev")
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

func TuneFindTrending() (tunefindTrending tunefind.TrendingSongs, err error) {

	response, err := http.Get(constants.TuneFindTrending)
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
