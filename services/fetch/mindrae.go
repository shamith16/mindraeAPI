package fetch

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"time"

	"github.com/shamith16/mindraeAPI/constants"
	"github.com/shamith16/mindraeAPI/entities/mindrae"
	"github.com/shamith16/mindraeAPI/entities/tmdb"
	"github.com/shamith16/mindraeAPI/entities/tunefind"
	"github.com/shamith16/mindraeAPI/utils"
)

/*
//TODO: Generalise Song Link Fetcher
*/

func MovieHome() {
	now := time.Now()
	var movieList []string
	var listOfMovie []mindrae.Movie
	var logs []string

	var sleepTime = 15
	var totalSongs int
	var totalTime = func(t time.Time, logs []string) {
		until := time.Until(t)
		l := fmt.Sprintf("Total Time taken is: %v\n", until.Minutes())
		fmt.Println(l)
		logs = append(logs, l)
	}

	var filterMovies = func(tunefind *tunefind.MovieSearch) (isMatched bool) {
		re := regexp.MustCompile(` \(aka.*`).ReplaceAll([]byte(tunefind.Movie.Name), []byte(""))

		name := url.QueryEscape(string(re))

		year := utils.YearStripper(tunefind.Movie.ReleaseDate)

		tmdbMovieBrowse, err := TmdbMovieBrowse(name, year)
		if err != nil {
			curErr := fmt.Sprintf("Error occurred at function MovieHome() TmdbMovieBrowse() is nil: %s\n", err)
			fmt.Println(curErr)
			logs = append(logs, curErr)
			totalTime(now, logs)
			return
		}

		listOfTResults := tmdbMovieBrowse.Results

		isMatched = false

		var matchedResult tmdb.MovieBrowseResult

		var tmdbMovie tmdb.MovieSearch

		tmdbResult := fmt.Sprintf("Total Number of Results for movie %s is: %d\n", name, tmdbMovieBrowse.TotalResults)
		fmt.Println(tmdbResult)
		logs = append(logs, tmdbResult)

		if tmdbMovieBrowse.TotalResults == 0 {
			resultNotFound := fmt.Sprintf("Movie %s not found in tmdbMovie skipping it\n", name)
			fmt.Println(resultNotFound)
			logs = append(logs, resultNotFound)

		} else {
			for _, result := range listOfTResults {
				if result.ReleaseDate[:4] == tunefind.Movie.ReleaseDate[:4] {
					a := fmt.Sprintf("%s Movie is matched using Year %s\n", name, year)
					fmt.Println(a)
					logs = append(logs, a)

					isMatched = true

					matchedResult = result
					break
				} else {
					a := fmt.Sprintf("%s Movie is coundn't be matched using Year %s\n", name, year)
					fmt.Println(a)
					logs = append(logs, a)

				}
			}
		}

		if isMatched {
			tmdbMovie, err = TmdbMovieSearch(matchedResult.Id)
			if err != nil {
				curErr := fmt.Sprintf("Error occurred at function MovieHome() TmdbMovieSearch() is nil: %s\n", err)
				fmt.Println(curErr)
				logs = append(logs, curErr)
				totalTime(now, logs)
				return
			}
			listOfMovie = append(listOfMovie, mindrae.Movie{
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
	var songLinkFetcher = func(movie *tunefind.MovieSearch) {
		//TODO: Handle link if any not present
		//TODO: itunes and apple music are essentially same Optimise it
		var spotify string
		var applemusic string
		var itunes string
		var youtube string
		var amazon string

		if totalSongs <= 1 {
			sleepTime = 5
		} else if totalSongs <= 10 {
			sleepTime = 10
		} else if totalSongs > 10 && totalSongs <= 15 {
			sleepTime = 15
		} else if totalSongs > 15 {
			sleepTime = 20
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

			ama, _ := client.Get(constants.TunefindBaseURL + movie.SongEvents[i].Song.Amazon)
			amazon = ama.Header.Get("location")

			movie.SongEvents[i].Song.Spotify = spotify
			movie.SongEvents[i].Song.Applemusic = applemusic
			movie.SongEvents[i].Song.Itunes = itunes
			movie.SongEvents[i].Song.Youtube = youtube
			movie.SongEvents[i].Song.Amazon = amazon
		}

		a := fmt.Sprintf("Fetch Song Links for Movies %s Done\n", movie.Movie.Name)
		fmt.Println(a)
		logs = append(logs, a)

	}
	var addStringToList = func(listAppendTo []string, strList ...[]string) {
		for _, strings := range strList {
			for _, s := range strings {
				listAppendTo = append(listAppendTo, s)
			}
		}
	}
	a := fmt.Sprintf("Fetching TuneFindMovieHome Started At: %s\n", now)
	fmt.Println(a)
	logs = append(logs, a)
	tunefindMovieHome, err := TuneFindMovieHome()
	if err != nil {
		curErr := fmt.Sprintf("Error occurred at function MovieHome() TunefindMovieHome() is nil: %s\n", err)
		fmt.Println(curErr)
		logs = append(logs, curErr)
		totalTime(now, logs)
		return
	}

	addStringToList(movieList, tunefindMovieHome.Slider, tunefindMovieHome.Featured, tunefindMovieHome.RecentlyAdded)

	{
		movieListLen := fmt.Sprintf("Total Number To Movies to be fetched from tunefind is: %d\n", len(movieList))
		fmt.Println(movieListLen)
		logs = append(logs, movieListLen)
	}

	for _, movie := range movieList {

		tunefindMovieSearch, err := TuneFindMovieSearch(movie)
		if err != nil {
			curErr := fmt.Sprintf("Error occurred at function MovieHome() TunefindMovieSearch() is nil: %s\n", err)
			fmt.Println(curErr)
			logs = append(logs, curErr)
			totalTime(now, logs)
			return
		}

		{
			totalSongs = len(tunefindMovieSearch.SongEvents)
			a := fmt.Sprintf("Calling TunefindSongLinkFetcher total number of songs for %s movie is %d\n", tunefindMovieSearch.Movie.Name, totalSongs)
			fmt.Println(a)
			logs = append(logs, a)
			songLinkFetcher(&tunefindMovieSearch)

		}

		a = fmt.Sprintf("Checking if %s movie is in Tmdb: \n", tunefindMovieSearch.Movie.Name)
		fmt.Println(a)
		logs = append(logs, a)

		b := filterMovies(&tunefindMovieSearch)
		a = fmt.Sprintf("Movie has been filter result is %v\n", b)
		fmt.Println(a)
		logs = append(logs, a)

		a = fmt.Sprintf("Current sleep time is: %d\n", sleepTime)
		fmt.Println(a)
		logs = append(logs, a)
		time.Sleep(time.Duration(sleepTime) * time.Second)
	}

	var movieHome = mindrae.MovieHome{Movies: listOfMovie}
	a = fmt.Sprintf("Total Number of movies matched is: %d", len(movieList)-len(listOfMovie))
	fmt.Println(a)
	logs = append(logs, a)
	marshal, err := json.Marshal(movieHome)
	if err != nil {
		curErr := fmt.Sprintf("Error occurred at function MovieHome() Marshal Error: %s\n", err)
		fmt.Println(curErr)
		logs = append(logs, curErr)
		totalTime(now, logs)
		return
	}

	err = utils.WriteToFile("movie-home.json", marshal, "jsons", "new")
	if err != nil {
		curErr := fmt.Sprintf("Error occurred at function MovieHome() Writing to file: %s\n", err)
		fmt.Println(curErr)
		logs = append(logs, curErr)
		totalTime(now, logs)
		return
	}

	var ctx []byte
	for _, s := range logs {
		ctx = append(ctx, []byte(s)...)
	}
	_ = utils.WriteToFile("logs.txt", ctx, "logs", "append")
	finished := fmt.Sprintf("Fetching Movie List took: %v", time.Until(now))
	fmt.Println(finished)
	logs = append(logs, finished)
	totalTime(now, logs)
}
