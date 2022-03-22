package https

import (
	"encoding/json"
	"fmt"
	"github.com/shamith16/mindraeAPI/constants"
	"github.com/shamith16/mindraeAPI/entities/mindrae"
	"github.com/shamith16/mindraeAPI/entities/tmdb"
	"github.com/shamith16/mindraeAPI/entities/tunefind"
	"github.com/shamith16/mindraeAPI/utils"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

var movieList []string

var listOfMovie []mindrae.Movie

func MovieHome() {
	tunefindMovieHome, _ := TuneFindMovieHome()

	addStringToList := func(strList ...[]string) {
		for _, strings := range strList {
			for _, s := range strings {
				movieList = append(movieList, s)
			}
		}
	}

	addStringToList(tunefindMovieHome.Slider, tunefindMovieHome.Featured, tunefindMovieHome.RecentlyAdded)

	movieListLen := fmt.Sprintf("Total tundFindMovieHome Result: %d", len(movieList))

	fmt.Println(movieListLen)

	for _, movie := range movieList {

		fmt.Println("Searching for tunefind movie: ", movie)

		tunefindMovieSearch, _ := TuneFindMovieSearch(movie)

		fmt.Println("filterMovieFromResults for movie: ", tunefindMovieSearch.Movie.Name)

		TunefindSongLinkFetcher(&tunefindMovieSearch)

		filterMovieFromResults(&tunefindMovieSearch)

		fmt.Println("filterMovieFromResults for movie: ", tunefindMovieSearch.Movie.Name, " done.")

		time.Sleep(30 * time.Second)
	}
	var movieHome = mindrae.MovieHome{Movies: listOfMovie}

	marshal, _ := json.Marshal(movieHome)

	err := utils.WriteToFile("movie-home.json", marshal)
	if err != nil {
		return
	}

}

func TunefindSongLinkFetcher(movie *tunefind.MovieSearch) {

	var spotify string
	var applemusic string
	var itunes string
	var youtube string
	var amazon string

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

	//for i := range movie.HotSongs {
	//	spo, _ := client.Get(constants.TunefindBaseURL + movie.HotSongs[i].Spotify)
	//	spotify = spo.Header.Get("location")
	//
	//	app, _ := client.Get(constants.TunefindBaseURL + movie.HotSongs[i].Applemusic)
	//	applemusic = app.Header.Get("location")
	//
	//	itu, _ := client.Get(constants.TunefindBaseURL + movie.HotSongs[i].Itunes)
	//	itunes = itu.Header.Get("location")
	//
	//	you, _ := client.Get(constants.TunefindBaseURL + movie.HotSongs[i].Youtube)
	//	youtube = you.Header.Get("location")
	//
	//	ama, _ := client.Get(constants.TunefindBaseURL + movie.HotSongs[i].Amazon)
	//	amazon = ama.Header.Get("location")
	//
	//	movie.HotSongs[i].Spotify = spotify
	//	movie.HotSongs[i].Applemusic = applemusic
	//	movie.HotSongs[i].Itunes = itunes
	//	movie.HotSongs[i].Youtube = youtube
	//	movie.HotSongs[i].Amazon = amazon
	//}

}

func filterMovieFromResults(tunefind *tunefind.MovieSearch) {
	re := regexp.MustCompile(` \(aka.*`).ReplaceAll([]byte(tunefind.Movie.Name), []byte(""))

	name := url.QueryEscape(string(re))

	year := utils.YearStripper(tunefind.Movie.ReleaseDate)

	tmdbMovieBrowse, _ := TmdbMovieBrowse(name, year)

	listOfTResults := tmdbMovieBrowse.Results

	var isMatched = false

	var matchedResult tmdb.MovieBrowseResult

	var tmdbMovie tmdb.MovieSearch

	tmdbResult := fmt.Sprintf("Result of TmdbMovieBrowse is: %d", tmdbMovieBrowse.TotalResults)

	fmt.Println(tmdbResult)

	if tmdbMovieBrowse.TotalResults == 0 {

		resultNotFound := fmt.Sprintf("Movie %s not found in tmdbMovie skipping it", name)

		fmt.Println(resultNotFound)

		return

	} else {

		for _, result := range listOfTResults {

			if result.ReleaseDate[:4] == tunefind.Movie.ReleaseDate[:4] {

				fmt.Println(result.Title,
					"is matched with tunefind using release year")

				isMatched = true

				matchedResult = result

				break
			} else {
				fmt.Println(result.Title,
					"is not matched with tunefind using release year")

				continue
			}
		}
	}

	if isMatched {

		tmdbMovie, _ = TmdbMovieSearch(matchedResult.Id)

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

}

//TODO: HotSongs Proper link just check name and copy from SongEvents song oject
