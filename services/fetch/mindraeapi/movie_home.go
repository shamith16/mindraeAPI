package mindraeapi

import (
	"encoding/json"
	"fmt"
	"github.com/shamith16/mindraeAPI/constants"
	"github.com/shamith16/mindraeAPI/entities/mindraemodel"
	"github.com/shamith16/mindraeAPI/entities/tmdbmodel"
	"github.com/shamith16/mindraeAPI/entities/tunefindmodel"
	"github.com/shamith16/mindraeAPI/services/fetch/tmdbapi"
	"github.com/shamith16/mindraeAPI/services/fetch/tunefindapi"
	"github.com/shamith16/mindraeAPI/utils"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

func MovieHome() {

	var (
		startTime   = time.Now()
		movieList   []string
		listOfMovie []mindraemodel.Movie
		logs        []string
		sleepTime   = 0
		totalSongs  int
		movieHome   mindraemodel.MovieHome
	)

	var (
		addToLog = func(s string) {
			fmt.Println(s)
			logs = append(logs, s)
		}

		songLinkFetcher = func(movie *tunefindmodel.MovieSearch) {
			//TODO: Handle link if any not present
			//TODO: itunes and apple music are essentially same Optimise it
			var spotify string
			var applemusic string
			var itunes string
			var youtube string
			//var amazon string

			if totalSongs <= 1 {
				sleepTime = 4
			} else if totalSongs <= 10 {
				sleepTime = 8
			} else if totalSongs > 10 && totalSongs <= 15 {
				sleepTime = 15
			} else if totalSongs > 15 {
				sleepTime = 19
			}

			RedirectHandler := func(req *http.Request, via []*http.Request, times int) error {
				err := fmt.Errorf("redirect policy: stopped after %d times", times)
				if len(via) >= times {
					return err
				}
				return nil
			}

			client := &http.Client{
				CheckRedirect: func(req *http.Request, via []*http.Request) error {
					return RedirectHandler(req, via, 2)
				},
			}

			for i := range movie.SongEvents {
				spo, _ := client.Get(constants.TunefindBaseURL + movie.SongEvents[i].Song.Spotify)
				spotify = spo.Header.Get("location")

				app, _ := client.Get(constants.TunefindBaseURL + movie.SongEvents[i].Song.Applemusic)
				applemusic = app.Header.Get("location")

				itu, _ := client.Get(constants.TunefindBaseURL + movie.SongEvents[i].Song.Itunes)
				itunes = itu.Header.Get("location")

				you, _ := client.Get(constants.TunefindBaseURL + movie.SongEvents[i].Song.Youtube)
				youtube = you.Header.Get("location")

				//ama, _ := client.Get(constants.TunefindBaseURL + movie.SongEvents[i].Song.Amazon)
				//amazon = ama.Header.Get("location")

				movie.SongEvents[i].Song.Spotify = spotify
				movie.SongEvents[i].Song.Applemusic = applemusic
				movie.SongEvents[i].Song.Itunes = itunes
				movie.SongEvents[i].Song.Youtube = youtube
				//movie.SongEvents[i].Song.Amazon = amazon
			}

			addToLog(fmt.Sprintf("Fetch Song Links for Movies %s Done\n", movie.Movie.Name))

		}

		filterMovies = func(tunefind *tunefindmodel.MovieSearch) (isMatched bool) {

			re := regexp.MustCompile(` \(aka.*`).ReplaceAll([]byte(tunefind.Movie.Name), []byte(""))

			name := url.QueryEscape(string(re))

			year := utils.YearStripper(tunefind.Movie.ReleaseDate)

			tmdbMovieSearch, err := tmdbapi.MovieSearch(name, year)
			if err != nil {
				addToLog(fmt.Sprintf("Error occurred at function Home() MovieSearch() is nil: %s\n", err))
				return
			}

			listOfTResults := tmdbMovieSearch.Results

			isMatched = false

			var (
				matchedResult tmdbmodel.MovieBrowseResult
				tmdbMovie     tmdbmodel.MovieSearch
			)

			addToLog(fmt.Sprintf("Total Number of Results for movie %s is: %d\n", name, tmdbMovieSearch.TotalResults))

			if tmdbMovieSearch.TotalResults == 0 {
				addToLog(fmt.Sprintf("Movie %s not found in tmdbMovie skipping it\n", name))

			} else {
				for _, result := range listOfTResults {
					if result.ReleaseDate[:4] == tunefind.Movie.ReleaseDate[:4] {
						addToLog(fmt.Sprintf("%s Movie is matched using Year %s\n", name, year))

						isMatched = true

						matchedResult = result
						break
					} else {
						addToLog(fmt.Sprintf("%s Movie is couldn't be matched using Year %s\n", name, year))
						continue
					}
				}
			}

			if isMatched {
				tmdbMovie, err = tmdbapi.MovieSearchById(matchedResult.Id)
				if err != nil {
					addToLog(fmt.Sprintf("Error occurred at function Home() MovieSearchById() is nil: %s\n", err))
					return
				}
				totalSongs = len(tunefind.SongEvents)
				addToLog(fmt.Sprintf("Calling TunefindSongLinkFetcher total number of songs for %s movie is %d\n", tunefind.Movie.Name, totalSongs))
				songLinkFetcher(tunefind)
				listOfMovie = append(listOfMovie, mindraemodel.Movie{
					TuneFindId:            tunefind.Movie.ID,
					TuneFindName:          tunefind.Movie.Name,
					TuneFindNameStub:      tunefind.Movie.NameStub,
					TuneFindType:          tunefind.Movie.Type,
					TuneFindTypeFull:      tunefind.Movie.TypeFull,
					TuneFindSongsCount:    tunefind.Movie.SongsCount,
					TuneFindReleaseDate:   tunefind.Movie.ReleaseDate,
					TuneFindAirDay:        tunefind.Movie.Event.AirdateDay,
					TuneFindAirMonth:      tunefind.Movie.Event.AirdateMonth,
					TuneFindAirMonthShort: tunefind.Movie.Event.AirdateMonthShort,
					TuneFindAirYear:       tunefind.Movie.Event.AirdateYear,
					TuneFindIsAired:       tunefind.Movie.Event.IsAired,
					SongEvents:            tunefind.SongEvents,
					HotSongs:              tunefind.HotSongs,
					TmdbIsAdult:           tmdbMovie.Adult,
					TmdbBackDropPath:      tmdbMovie.BackdropPath,
					TmdbBudget:            int64(tmdbMovie.Budget),
					Genres:                tmdbMovie.Genres,
					MovieHomePageUrl:      tmdbMovie.Homepage,
					TmdbId:                int64(tmdbMovie.Id),
					ImdbId:                tmdbMovie.ImdbId,
					OriginalLanguage:      tmdbMovie.OriginalLanguage,
					OriginalTitle:         tmdbMovie.OriginalTitle,
					MovieOverview:         tmdbMovie.Overview,
					Popularity:            tmdbMovie.Popularity,
					PosterPath:            tmdbMovie.PosterPath,
					TmdbReleaseDate:       tmdbMovie.ReleaseDate,
					Revenue:               int64(tmdbMovie.Revenue),
					Runtime:               int64(tmdbMovie.Runtime),
					TmdbStatus:            tmdbMovie.Status,
					TagLine:               tmdbMovie.Tagline,
					Title:                 tmdbMovie.Title,
					Video:                 tmdbMovie.Video,
					VoteAverage:           tmdbMovie.VoteAverage,
					VoteCount:             int64(tmdbMovie.VoteCount),
					ExternalIds:           tmdbMovie.ExternalIds,
					Videos:                tmdbMovie.Videos,
					Images:                tmdbMovie.Images,
				})
			}
			return
		}

		addStringToList = func(listAppendTo *[]string, strList ...[]string) {
			for _, strings := range strList {
				for _, s := range strings {
					*listAppendTo = append(*listAppendTo, s)
				}
			}
		}
	)

	addToLog(fmt.Sprintf("Fetching MindraeMovieHome Started at %v\n", startTime))

	tunefindMovieHome, err := tunefindapi.MovieHome()
	if err != nil {
		addToLog(fmt.Sprintf("Error occurred at function Home() TunefindMovieHome() is nil: %s\n", err))
		return
	}

	addStringToList(&movieList, tunefindMovieHome.Slider, tunefindMovieHome.Featured, tunefindMovieHome.RecentlyAdded)

	addToLog(fmt.Sprintf("Total Number To Movies to be fetched from tunefind is: %d\n", len(movieList)))

	for _, movie := range movieList {

		tunefindMovieSearch, err := tunefindapi.MovieSearch(movie)
		if err != nil {
			addToLog(fmt.Sprintf("Error occurred at function Home() TunefindMovieSearch() is nil: %s\n", err))
			return
		}

		addToLog(fmt.Sprintf("Checking if %s movie is in Tmdb: \n", tunefindMovieSearch.Movie.Name))

		b := filterMovies(&tunefindMovieSearch)
		addToLog(fmt.Sprintf("Movie has been filter result is %v\n", b))

		addToLog(fmt.Sprintf("Current sleep time is: %d\n", sleepTime))
		time.Sleep(time.Duration(sleepTime) * time.Second)
	}

	movieHome = mindraemodel.MovieHome{Movies: listOfMovie}

	addToLog(fmt.Sprintf("Total Number of movies matched out of %d is %d\n", len(movieList), len(listOfMovie)))

	marshal, err := json.Marshal(movieHome)
	if err != nil {
		addToLog(fmt.Sprintf("Error occurred at function Home() Marshal Error: %s\n", err))
		return
	}

	err = utils.WriteToFile("movie-home.json", marshal, "json/", "new")
	if err != nil {
		addToLog(fmt.Sprintf("Error occurred at function Home() Writing to file: %s\n", err))
		return
	}
	addToLog(fmt.Sprintf("Fetching Movie List took: %v", time.Until(startTime)))

	_ = utils.WriteToFile("logs.txt", logs, "log/", "append")

}
