package mindraeapi

import (
	"fmt"
	"github.com/shamith16/mindraeAPI/constants"
	"github.com/shamith16/mindraeAPI/entities/mindraemodel"
	"github.com/shamith16/mindraeAPI/entities/tmdbmodel"
	"github.com/shamith16/mindraeAPI/entities/tunefindmodel"
	"github.com/shamith16/mindraeAPI/services/fetch/tmdbapi"
	"github.com/shamith16/mindraeAPI/utils"
	"net/http"
	"net/url"
	"regexp"
	"sync"
	"time"
)

func songFetcher(event *tunefindmodel.SongEvent, wg *sync.WaitGroup) {

	defer wg.Done()

	var spotify string
	var applemusic string
	var itunes string
	var youtube string

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

	event.Song.Spotify = spotify
	event.Song.Applemusic = applemusic
	event.Song.Itunes = itunes
	event.Song.Youtube = youtube

}

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

func eventsLinkFetcher(events *tunefindmodel.MovieSearch, sleepTime *int, totalSongs int, routeName string) {

	if totalSongs <= 1 {
		*sleepTime = 4
	} else if totalSongs <= 10 {
		*sleepTime = 8
	} else if totalSongs > 10 && totalSongs <= 15 {
		*sleepTime = 15
	} else if totalSongs > 15 {
		*sleepTime = 19
	}

	wg.Add(totalSongs)
	for _, event := range events.SongEvents {
		songFetcher(&event, &wg)
	}
	wg.Wait()

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

func filterMovies(tunefind *tunefindmodel.MovieSearch, listOfMovie *[]mindraemodel.Movie, routeName string) (isMatched bool) {

	sleepTime := 0

	re := regexp.MustCompile(` \(aka.*`).ReplaceAll([]byte(tunefind.Movie.Name), []byte(""))

	name := url.QueryEscape(string(re))

	year := utils.YearStripper(tunefind.Movie.ReleaseDate)

	tmdbMovieSearch, err := tmdbapi.MovieSearch(name, year)
	if err != nil {
		addToLog(fmt.Sprintf("Error occurred at function Home() MovieSearch() is nil: %s\n", err), routeName)
		return
	}

	listOfTResults := tmdbMovieSearch.Results

	isMatched = false

	var (
		matchedResult tmdbmodel.MovieBrowseResult
		tmdbMovie     tmdbmodel.MovieSearch
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
		tmdbMovie, err = tmdbapi.MovieSearchById(matchedResult.Id)
		if err != nil {
			addToLog(fmt.Sprintf("Error occurred at function Home() MovieSearchById() is nil: %s\n", err), routeName)
			return
		}
		totalSongs := len(tunefind.SongEvents)
		addToLog(fmt.Sprintf("Calling TunefindSongLinkFetcher total number of songs for %s movie is %d\n", tunefind.Movie.Name, totalSongs), routeName)
		eventsLinkFetcher(tunefind, &sleepTime, totalSongs, routeName)
		addToLog(fmt.Sprintf("Current sleep time is: %d\n", sleepTime), routeName)
		time.Sleep(time.Duration(sleepTime) * time.Second)
		*listOfMovie = append(*listOfMovie, mindraemodel.Movie{
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
