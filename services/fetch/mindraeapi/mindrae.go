package mindraeapi

import (
	"encoding/json"
	"fmt"
	"github.com/shamith16/mindraeAPI/constants"
	"github.com/shamith16/mindraeAPI/entities/mindraemodel"
	"github.com/shamith16/mindraeAPI/entities/tmdbmodel"
	"github.com/shamith16/mindraeAPI/entities/tunefindmodel"
	tmdbapi "github.com/shamith16/mindraeAPI/services/fetch/tmdbapi"
	tunefindapi "github.com/shamith16/mindraeAPI/services/fetch/tunefindapi"
	"github.com/shamith16/mindraeAPI/utils"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

/*
//TODO: Generalise Song Link Fetcher
*/

func Home() {
	now := time.Now()

	var movieList []string

	var listOfMovie []mindraemodel.Movie

	var logs []string

	log := fmt.Sprintf("Fetching Home Started at %v\n", now)

	logs = append(logs, log)

	var sleepTime = 15

	var totalSongs int

	var songLinkFetcher = func(movie *tunefindmodel.MovieSearch) {
		//TODO: Handle link if any not present
		//TODO: itunes and apple music are essentially same Optimise it
		var spotify string
		var applemusic string
		var itunes string
		var youtube string
		//	var amazon string

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

			//			ama, _ := client.Get(constants.TunefindBaseURL + movie.SongEvents[i].Song.Amazon)
			//			amazon = ama.Header.Get("location")

			movie.SongEvents[i].Song.Spotify = spotify
			movie.SongEvents[i].Song.Applemusic = applemusic
			movie.SongEvents[i].Song.Itunes = itunes
			movie.SongEvents[i].Song.Youtube = youtube
			//			movie.SongEvents[i].Song.Amazon = amazon
		}

		log := fmt.Sprintf("Fetch Song Links for Movies %s Done\n", movie.Movie.Name)
		fmt.Println(log)
		logs = append(logs, log)

	}

	var filterMovies = func(tunefind *tunefindmodel.MovieSearch) (isMatched bool) {

		re := regexp.MustCompile(` \(aka.*`).ReplaceAll([]byte(tunefind.Movie.Name), []byte(""))

		name := url.QueryEscape(string(re))

		year := utils.YearStripper(tunefind.Movie.ReleaseDate)

		tmdbMovieBrowse, err := tmdbapi.MovieSearch(name, year)
		if err != nil {
			log := fmt.Sprintf("Error occurred at function Home() MovieSearch() is nil: %s\n", err)
			fmt.Println(log)
			logs = append(logs, log)
			return
		}

		listOfTResults := tmdbMovieBrowse.Results

		isMatched = false

		var matchedResult tmdbmodel.MovieBrowseResult

		var tmdbMovie tmdbmodel.MovieSearch

		log := fmt.Sprintf("Total Number of Results for movie %s is: %d\n", name, tmdbMovieBrowse.TotalResults)
		fmt.Println(log)
		logs = append(logs, log)

		if tmdbMovieBrowse.TotalResults == 0 {
			log := fmt.Sprintf("Movie %s not found in tmdbMovie skipping it\n", name)
			fmt.Println(log)
			logs = append(logs, log)

		} else {
			for _, result := range listOfTResults {
				if result.ReleaseDate[:4] == tunefind.Movie.ReleaseDate[:4] {
					log := fmt.Sprintf("%s Movie is matched using Year %s\n", name, year)
					fmt.Println(log)
					logs = append(logs, log)

					isMatched = true

					matchedResult = result
					break
				} else {
					log := fmt.Sprintf("%s Movie is coundn't be matched using Year %s\n", name, year)
					fmt.Println(log)
					logs = append(logs, log)

				}
			}
		}

		if isMatched {
			tmdbMovie, err = tmdbapi.MovieSearchById(matchedResult.Id)
			totalSongs = len(tunefind.SongEvents)
			log := fmt.Sprintf("Calling TunefindSongLinkFetcher total number of songs for %s movie is %d\n", tunefind.Movie.Name, totalSongs)
			fmt.Println(log)
			logs = append(logs, log)
			songLinkFetcher(tunefind)
			if err != nil {
				log := fmt.Sprintf("Error occurred at function Home() MovieSearchById() is nil: %s\n", err)
				fmt.Println(log)
				logs = append(logs, log)
				return
			}
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

	var addStringToList = func(listAppendTo *[]string, strList ...[]string) {
		for _, strings := range strList {
			for _, s := range strings {
				*listAppendTo = append(*listAppendTo, s)
			}
		}
	}

	log = fmt.Sprintf("Fetching Home Started At: %s\n", now)

	fmt.Println(log)

	logs = append(logs, log)

	tunefindMovieHome, err := tunefindapi.MovieHome()
	if err != nil {
		log := fmt.Sprintf("Error occurred at function Home() TunefindMovieHome() is nil: %s\n", err)
		fmt.Println(log)
		logs = append(logs, log)
		return
	}

	addStringToList(&movieList, tunefindMovieHome.Slider, tunefindMovieHome.Featured, tunefindMovieHome.RecentlyAdded)

	log = fmt.Sprintf("Total Number To Movies to be fetched from tunefind is: %d\n", len(movieList))
	fmt.Println(log)
	logs = append(logs, log)

	for _, movie := range movieList {

		tunefindMovieSearch, err := tunefindapi.MovieSearch(movie)
		if err != nil {
			log := fmt.Sprintf("Error occurred at function Home() TunefindMovieSearch() is nil: %s\n", err)
			fmt.Println(log)
			logs = append(logs, log)
			return
		}

		log = fmt.Sprintf("Checking if %s movie is in Tmdb: \n", tunefindMovieSearch.Movie.Name)
		fmt.Println(log)
		logs = append(logs, log)

		b := filterMovies(&tunefindMovieSearch)
		log = fmt.Sprintf("Movie has been filter result is %v\n", b)
		fmt.Println(log)
		logs = append(logs, log)

		log = fmt.Sprintf("Current sleep time is: %d\n", sleepTime)
		fmt.Println(log)
		logs = append(logs, log)
		time.Sleep(time.Duration(sleepTime) * time.Second)
	}

	var movieHome = mindraemodel.MovieHome{Movies: listOfMovie}

	log = fmt.Sprintf("Total Number of movies matched out of %d is %d\n", len(movieList), len(listOfMovie))
	fmt.Println(log)
	logs = append(logs, log)
	marshal, err := json.Marshal(movieHome)
	if err != nil {
		log := fmt.Sprintf("Error occurred at function Home() Marshal Error: %s\n", err)
		fmt.Println(log)
		logs = append(logs, log)
		return
	}

	err = utils.WriteToFile("movie-home.json", marshal, "json/", "new")
	if err != nil {
		log := fmt.Sprintf("Error occurred at function Home() Writing to file: %s\n", err)
		fmt.Println(log)
		logs = append(logs, log)
		return
	}

	_ = utils.WriteToFile("logs.txt", logs, "log/", "append")
	finished := fmt.Sprintf("Fetching Movie List took: %v", time.Until(now))
	fmt.Println(finished)
	logs = append(logs, finished)
	log = string(time.Until(now))
	logs = append(logs, log)
}
