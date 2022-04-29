package mindraeapi

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/shamith16/mindraeAPI/entities/mindraemodel"
	"github.com/shamith16/mindraeAPI/services/fetch/tunefindapi"
	"github.com/shamith16/mindraeAPI/utils"
	"time"
)

const routeName = "movie_home.go"

func MovieHome() {
	var (
		startTime   = time.Now()
		movieList   []string
		listOfMovie []mindraemodel.Movie
		movieHome   mindraemodel.MovieHome
	)

	tunefindMovieHome, err := tunefindapi.MovieHome()
	if err != nil {
		addToLog(fmt.Sprintf("Error occurred at function Home() TunefindMovieHome() is nil: %s\n", err), routeName)
		return
	}

	addStringToList(&movieList, tunefindMovieHome.Slider, tunefindMovieHome.Featured, tunefindMovieHome.RecentlyAdded)

	addToLog(fmt.Sprintf("Total Number To Movies to be fetched from tunefind is: %d\n", len(movieList)), routeName)

	for _, movie := range movieList {
		tunefindMovieSearch, err := tunefindapi.MovieSearch(movie)
		if err != nil {
			addToLog(fmt.Sprintf("Error occurred at function Home() TunefindMovieSearch() is nil: %s\n", err), routeName)
			return
		}

		addToLog(fmt.Sprintf("Checking if %s movie is in Tmdb: \n", tunefindMovieSearch.Movie.Name), routeName)
		matchedBool := filterMovie(&tunefindMovieSearch, &listOfMovie, routeName)
		addToLog(fmt.Sprintf("Movie has been filter result is %v\n", matchedBool), routeName)
	}

	movieHome = mindraemodel.MovieHome{Movies: listOfMovie}
	addToLog(fmt.Sprintf("Total Number of movies matched out of %d is %d\n", len(movieList), len(listOfMovie)), routeName)
	marshal, err := jsoniter.Marshal(movieHome)
	if err != nil {
		addToLog(fmt.Sprintf("Error occurred at function Home() Marshal Error: %s\n", err), routeName)
		return
	}
	err = utils.WriteToFile("movie-home.json", marshal, "json/", "new")
	if err != nil {
		addToLog(fmt.Sprintf("Error occurred at function Home() Writing to file: %s\n", err), routeName)
		return
	}
	addToLog(fmt.Sprintf("Fetching Movie List took: %v", time.Until(startTime)), routeName)
}
