package https

import (
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"

	"github.com/shamith16/mindraeAPI/entities/mindrae"
	"github.com/shamith16/mindraeAPI/utils"
)

func MindraeMovieHome() {
	tunefindMovieHome, _ := TuneFindMovieHome()
	var slider []mindrae.MergeStructure
	var featured []mindrae.MergeStructure
	var recent []mindrae.MergeStructure

	for _, s := range tunefindMovieHome.Slider {
		fmt.Println("Searching full info on tunefind for movie", s)
		tunefindMovie, err := TuneFindMovieSearch(s)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Successfully got result for movie", tunefindMovie.Movie.Name, "Whose ReleaseDate is", tunefindMovie.Movie.ReleaseDate)
		re := regexp.MustCompile(` \(aka.*`)
		replace := re.ReplaceAll([]byte(tunefindMovie.Movie.Name), []byte(""))
		name := url.QueryEscape(string(replace))
		tmdbMovieBrowse, _ := TmdbMovieBrowse(name, utils.YearStripper(tunefindMovie.Movie.ReleaseDate))
		fmt.Println(len(tmdbMovieBrowse.Results), "Total length of result")
		if len(tmdbMovieBrowse.Results) == 0 {
			slider = append(slider, mindrae.MergeStructure{
				TuneFindMovie: tunefindMovie,
			})
		} else {
			for _, result := range tmdbMovieBrowse.Results {
				fmt.Println("Currently check tmdb result name and release date with tunefind")
				fmt.Println(result.Title, result.ReleaseDate, "to tunefind", tunefindMovie.Movie.Name, tunefindMovie.Movie.ReleaseDate)
				if result.ReleaseDate[:4] == tunefindMovie.Movie.ReleaseDate[:4] {
					fmt.Println("Successfully matched")
					fmt.Println(result.Title, result.ReleaseDate, "with tune-find's", tunefindMovie.Movie.Name, tunefindMovie.Movie.ReleaseDate)
					tmdbMovieSearch, _ := TmdbMovieSearch(result.Id)
					slider = append(slider, mindrae.MergeStructure{
						TuneFindMovie: tunefindMovie,
						TmdbMovie:     tmdbMovieSearch,
					})
					fmt.Println("Current length of slider is: ", len(slider))
					break
				}
			}
		}

	}
	for _, s := range tunefindMovieHome.Featured {
		fmt.Println("Searching full info on tunefind for movie", s)
		tunefindMovie, err := TuneFindMovieSearch(s)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Successfully got result for movie", tunefindMovie.Movie.Name, "Whose ReleaseDate is", tunefindMovie.Movie.ReleaseDate)
		re := regexp.MustCompile(` \(aka.*`)
		replace := re.ReplaceAll([]byte(tunefindMovie.Movie.Name), []byte(""))
		name := url.QueryEscape(string(replace))
		tmdbMovieBrowse, _ := TmdbMovieBrowse(name, utils.YearStripper(tunefindMovie.Movie.ReleaseDate))
		fmt.Println(len(tmdbMovieBrowse.Results), "Total length of result")
		if len(tmdbMovieBrowse.Results) == 0 {
			featured = append(featured, mindrae.MergeStructure{
				TuneFindMovie: tunefindMovie,
			})
		} else {
			for _, result := range tmdbMovieBrowse.Results {
				fmt.Println("Currently check tmdb result name and release date with tunefind")
				fmt.Println(result.Title, result.ReleaseDate, "to tunefind", tunefindMovie.Movie.Name, tunefindMovie.Movie.ReleaseDate)
				if result.ReleaseDate[:4] == tunefindMovie.Movie.ReleaseDate[:4] {
					fmt.Println("Successfully matched")
					fmt.Println(result.Title, result.ReleaseDate, "with tune-find's", tunefindMovie.Movie.Name, tunefindMovie.Movie.ReleaseDate)
					tmdbMovieSearch, _ := TmdbMovieSearch(result.Id)
					featured = append(featured, mindrae.MergeStructure{
						TuneFindMovie: tunefindMovie,
						TmdbMovie:     tmdbMovieSearch,
					})
					fmt.Println("Current length of slider is: ", len(featured))
					break
				}
			}
		}
	}
	for _, s := range tunefindMovieHome.RecentlyAdded {
		fmt.Println("Searching full info on tunefind for movie", s)
		tunefindMovie, err := TuneFindMovieSearch(s)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Successfully got result for movie", tunefindMovie.Movie.Name, "Whose ReleaseDate is", tunefindMovie.Movie.ReleaseDate)
		re := regexp.MustCompile(` \(aka.*`)
		replace := re.ReplaceAll([]byte(tunefindMovie.Movie.Name), []byte(""))
		name := url.QueryEscape(string(replace))
		tmdbMovieBrowse, _ := TmdbMovieBrowse(name, utils.YearStripper(tunefindMovie.Movie.ReleaseDate))
		fmt.Println(len(tmdbMovieBrowse.Results), "Total length of result")
		if len(tmdbMovieBrowse.Results) == 0 {
			recent = append(recent, mindrae.MergeStructure{
				TuneFindMovie: tunefindMovie,
			})
		} else {
			for _, result := range tmdbMovieBrowse.Results {
				fmt.Println("Currently check tmdb result name and release date with tunefind")
				fmt.Println(result.Title, result.ReleaseDate, "to tunefind", tunefindMovie.Movie.Name, tunefindMovie.Movie.ReleaseDate)
				if result.ReleaseDate[:4] == tunefindMovie.Movie.ReleaseDate[:4] {
					fmt.Println("Successfully matched")
					fmt.Println(result.Title, result.ReleaseDate, "with tune-find's", tunefindMovie.Movie.Name, tunefindMovie.Movie.ReleaseDate)
					tmdbMovieSearch, _ := TmdbMovieSearch(result.Id)
					recent = append(recent, mindrae.MergeStructure{
						TuneFindMovie: tunefindMovie,
						TmdbMovie:     tmdbMovieSearch,
					})
					fmt.Println("Current length of slider is: ", len(recent))
					break
				}
			}
		}
	}

	a := mindrae.MovieHome{
		Slider:   slider,
		Featured: featured,
		Recent:   recent,
	}
	marshal, _ := json.Marshal(a)
	utils.WriteToFile("movie-home.json", marshal)

}
