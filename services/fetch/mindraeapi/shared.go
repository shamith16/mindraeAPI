package mindraeapi

import (
	"fmt"
	tmdb "github.com/cyruzin/golang-tmdb"
	"github.com/shamith16/mindraeAPI/constants"
	"github.com/shamith16/mindraeAPI/entities/mindraemodel"
	"github.com/shamith16/mindraeAPI/entities/tmdbmodel"
	"github.com/shamith16/mindraeAPI/entities/tunefindmodel"
	"github.com/shamith16/mindraeAPI/services/fetch/tmdbapi"
	"github.com/shamith16/mindraeAPI/utils"
	"net/http"
	"regexp"
	"time"
)

func addStringToList(listAppendTo *[]string, strList ...[]string) {
	for _, strings := range strList {
		for _, s := range strings {
			*listAppendTo = append(*listAppendTo, s)
		}
	}
}

func addToLog(s string, Name string) {
	fmt.Println(s)
	utils.Logger(s, Name)
}

func songFetcher(event tunefindmodel.SongEvent) (spotify, applemusic, itunes, youtube string) {
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

	spo, _ := client.Get(constants.TunefindBaseURL + event.Song.Spotify)
	spotify = spo.Header.Get("location")

	app, _ := client.Get(constants.TunefindBaseURL + event.Song.Applemusic)
	applemusic = app.Header.Get("location")

	itu, _ := client.Get(constants.TunefindBaseURL + event.Song.Itunes)
	itunes = itu.Header.Get("location")

	you, _ := client.Get(constants.TunefindBaseURL + event.Song.Youtube)
	youtube = you.Header.Get("location")

	return
}

func eventsLinkFetcher(events *tunefindmodel.MovieSearch, sleepTime *int, totalSongs int, routeName string) {

	var (
		s, a, itu, y string
	)

	if totalSongs <= 1 {
		*sleepTime = 6
	} else if totalSongs <= 10 {
		*sleepTime = 8
	} else if totalSongs > 10 && totalSongs <= 15 {
		*sleepTime = 15
	} else if totalSongs > 15 {
		*sleepTime = 19
	}
	for i := range events.SongEvents {
		s, a, itu, y = songFetcher(events.SongEvents[i])
		events.SongEvents[i].Song.Spotify = s
		events.SongEvents[i].Song.Applemusic = a
		events.SongEvents[i].Song.Itunes = itu
		events.SongEvents[i].Song.Youtube = y

	}
	for i := range events.HotSongs {
		for i2 := range events.SongEvents {
			if events.HotSongs[i].Name == events.SongEvents[i2].Song.Name {
				events.HotSongs[i] = events.SongEvents[i2].Song
				break
			}
		}
	}

	addToLog(fmt.Sprintf("Fetch Song Links for Movies %s Done\n", events.Movie.Name), routeName)

}

func filterMovie(tunefind *tunefindmodel.MovieSearch, listOfMovie *[]mindraemodel.Movie, routeName string) (isMatched bool) {

	sleepTime := 0
	isMatched = false

	re := regexp.MustCompile(` \(aka.*`).ReplaceAll([]byte(tunefind.Movie.Name), []byte(""))

	name := string(re)

	year := utils.YearStripper(tunefind.Movie.ReleaseDate)

	tmdbMovieSearch, err := tmdbapi.SearchTmdbMovie(name, map[string]string{
		"year": year,
	})
	if err != nil {
		addToLog(fmt.Sprintf("Error occurred at function Home() MovieSearch() is nil: %s\n", err), routeName)
		return
	}

	listOfTResults := tmdbMovieSearch.Results

	var (
		matchedResult tmdbmodel.SearchMovieResult
		tmdbMovie     tmdb.MovieDetails
	)

	addToLog(fmt.Sprintf("Total Number of Results for movie %s is: %d\n", name, tmdbMovieSearch.TotalResults), routeName)

	if tmdbMovieSearch.TotalResults == 0 {
		addToLog(fmt.Sprintf("Movie %s not found in tmdbMovie skipping it\n", name), routeName)
	} else {
		for _, result := range listOfTResults {
			if result.ReleaseDate[:4] == tunefind.Movie.ReleaseDate[:4] {
				addToLog(fmt.Sprintf("%s Movie is matched using Year %s\n", name, year), routeName)

				isMatched = true

				matchedResult = result
				break
			} else {
				addToLog(fmt.Sprintf("%s Movie is couldn't be matched using Year %s\n", name, year), routeName)
				continue
			}
		}
	}

	if isMatched {
		tmdbMovie, err = tmdbapi.GetMovieDetails(int(matchedResult.ID))
		if err != nil {
			addToLog(fmt.Sprintf("Error occurred at function Home() MovieSearchById() is nil: %s\n", err), routeName)
			return
		}

		totalSongs := len(tunefind.SongEvents)
		addToLog(fmt.Sprintf("Calling TunefindSongLinkFetcher total number of songs for %s movie is %d\n", tunefind.Movie.Name, totalSongs), routeName)
		eventsLinkFetcher(tunefind, &sleepTime, totalSongs, routeName)

		addToLog(fmt.Sprintf("Current sleep time is: %d\n", sleepTime), routeName)
		time.Sleep(time.Duration(sleepTime) * time.Second)

		tunefind := mindraemodel.TuneFind{
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
		}
		*listOfMovie = append(*listOfMovie, mindraemodel.Movie{
			TuneFind: tunefind,
			Tmdb:     tmdbMovie,
		})

	}

	return
}
