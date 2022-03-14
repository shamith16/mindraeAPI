package main

import (
	"encoding/json"
	"fmt"
	"github.com/shamith16/mindraeAPI/entities/tmdb"
	"github.com/shamith16/mindraeAPI/services/https"
	"log"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup

func FetchTmdb(w *sync.WaitGroup) {
	defer w.Done()
	var tmdbMovieSearch tmdb.MovieSearch
	var tmdbShowSearch tmdb.ShowSearch
	var tmdbSeasonSearch tmdb.SeasonSearch
	var tmdbEpisodeSearch tmdb.EpisodeSearch
	var tmdbMovie tmdb.MovieBrowse
	var tmdbShow tmdb.ShowBrowse

	now := time.Now()
	wg.Add(6)
	go func() {
		defer func() {
			fmt.Println("TmdbMovieBrowse finished at", time.Since(now))
			wg.Done()
		}()
		err := https.TmdbMovieBrowse("2012", "2009", &tmdbMovie)
		if err != nil {
			fmt.Println(err)

		} else {
			fmt.Println("TmdbMovieBrowse Done")
		}
	}()
	go func() {
		defer func() {
			fmt.Println("TmdbShowBrowse finished at", time.Since(now))
			wg.Done()
		}()
		err2 := https.TmdbShowBrowse("Friends", &tmdbShow)
		if err2 != nil {
			fmt.Println(err2)

		} else {
			fmt.Println("TmdbShowBrowse Done")

		}
	}()
	go func() {
		defer func() {
			fmt.Println("MovieSearch finished at", time.Since(now))
			wg.Done()
		}()
		err := https.TmdbMovieSearch(13804, &tmdbMovieSearch)
		if err != nil {
			fmt.Println(err)

		} else {
			fmt.Println("TmdbMovieSearch Done")
		}
	}()
	go func() {
		defer func() {
			fmt.Println("ShowSearch finished at", time.Since(now))
			wg.Done()
		}()
		err := https.TmdbShowSearch(1668, &tmdbShowSearch)
		if err != nil {
			fmt.Println(err)

		} else {
			fmt.Println("ShowSearch Done")
		}
	}()
	go func() {
		defer func() {
			fmt.Println("SeasonSearch finished at", time.Since(now))
			wg.Done()
		}()
		err := https.TmdbSeasonSearch(1668, 1, &tmdbSeasonSearch)
		if err != nil {
			fmt.Println(err)

		} else {
			fmt.Println("TmdbSeasonSearch Done")
		}
	}()
	go func() {
		defer func() {
			fmt.Println("EpisodeSearch finished at", time.Since(now))
			wg.Done()
		}()
		err := https.TmdbEpisodeSearch(1668, 1, 1, &tmdbEpisodeSearch)
		if err != nil {
			fmt.Println(err)

		} else {
			fmt.Println("TmdbEpisodeSearch Done")
		}
	}()
	wg.Wait()
	var tmdbSearchAll TmdbSearchAll = TmdbSearchAll{
		TmdbMovie:     tmdbMovie,
		TmdbShow:      tmdbShow,
		MovieSearch:   tmdbMovieSearch,
		ShowSearch:    tmdbShowSearch,
		SeasonSearch:  tmdbSeasonSearch,
		EpisodeSearch: tmdbEpisodeSearch,
	}
	customMarshal, _ := json.Marshal(tmdbSearchAll)
	func() {
		file, err := os.OpenFile(
			"tmdb.json",
			os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
			0666,
		)
		if err != nil {
			log.Fatal(err)
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {

			}
		}(file)

		// Write bytes to file
		byteSlice := customMarshal
		bytesWritten, err := file.Write(byteSlice)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Wrote %d bytes.\n", bytesWritten)
	}()
	elapsed := time.Since(now)
	fmt.Println("====================================================================")
	fmt.Println("Total Time Taken to Fetch Results and Write JSON to disk Tmdb: ", elapsed)
	fmt.Println("====================================================================")

}

type TmdbSearchAll struct {
	TmdbMovie     tmdb.MovieBrowse   `json:"tmdbMovie"`
	TmdbShow      tmdb.ShowBrowse    `json:"tmdbShow"`
	MovieSearch   tmdb.MovieSearch   `json:"tmdbMovieSearch"`
	ShowSearch    tmdb.ShowSearch    `json:"tmdbShowSearch"`
	SeasonSearch  tmdb.SeasonSearch  `json:"tmdbSeasonSearch"`
	EpisodeSearch tmdb.EpisodeSearch `json:"tmdbEpisodeSearch"`
}
