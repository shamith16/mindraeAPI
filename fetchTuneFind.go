package main

import (
	"encoding/json"
	"fmt"
	"github.com/shamith16/mindraeAPI/entities/tunefind"
	"github.com/shamith16/mindraeAPI/services/https"
	"log"
	"os"
	"sync"
	"time"
)

var wg1 sync.WaitGroup

func fetchTuneFind(w *sync.WaitGroup) {
	defer w.Done()
	wg1.Add(11)
	var tunefindMovieBrowseAll tunefind.MovieBrowse
	var tunefindShowBrowseAll tunefind.ShowBrowse
	var tunefindGameBrowseAll tunefind.GameBrowse
	var tunefindMovieHome tunefind.MovieGameHome
	var tunefindGameHome tunefind.MovieGameHome
	var tunefindShowHome tunefind.ShowHome
	var tunefindMovieSearch tunefind.MovieSearch
	var tunefindShowSearch tunefind.ShowSearch
	var tunefindSeasonSearch tunefind.SeasonSearch
	var tunefindEpisodeSearch tunefind.EpisodeSearch
	var tunefindTrending tunefind.TrendingSongs

	var customBrowseAll CustomBrowseAll

	now := time.Now()

	go func() {
		defer func() {
			fmt.Println("TuneFindShowBrowse finished at", time.Since(now))
			wg1.Done()
		}()
		err2 := https.TuneFindShowBrowse(&tunefindShowBrowseAll)
		if err2 != nil {
			fmt.Println(err2)

		} else {
			fmt.Println("TuneFindShowBrowse Done")

		}
	}()

	go func() {
		defer func() {
			fmt.Println("TuneFindMovieBrowse finished at", time.Since(now))
			wg1.Done()
		}()
		err2 := https.TuneFindMovieBrowse(&tunefindMovieBrowseAll)
		if err2 != nil {
			fmt.Println(err2)

		} else {
			fmt.Println("TuneFindMovieBrowse Done")

		}
	}()

	go func() {
		defer func() {
			fmt.Println("TuneFindGameBrowse finished at", time.Since(now))
			wg1.Done()
		}()
		err2 := https.TuneFindGameBrowse(&tunefindGameBrowseAll)
		if err2 != nil {
			fmt.Println(err2)

		} else {
			fmt.Println("TuneFindGameBrowse Done")

		}
	}()

	go func() {
		defer func() {
			fmt.Println("TuneFindTrending finished at", time.Since(now))
			wg1.Done()
		}()
		err2 := https.TuneFindTrending(&tunefindTrending)
		if err2 != nil {
			fmt.Println(err2)

		} else {
			fmt.Println("TuneFindTrending Done")

		}
	}()

	go func() {
		defer func() {
			fmt.Println("TuneFindShowHome finished at", time.Since(now))
			wg1.Done()
		}()
		err2 := https.TuneFindShowHome(&tunefindShowHome)
		if err2 != nil {
			fmt.Println(err2)

		} else {
			fmt.Println("TuneFindTrending Done")

		}
	}()

	go func() {
		defer func() {
			fmt.Println("TuneFindGameHome finished at", time.Since(now))
			wg1.Done()
		}()
		err2 := https.TuneFindGameHome(&tunefindGameHome)
		if err2 != nil {
			fmt.Println(err2)

		} else {
			fmt.Println("TuneFindGameHome Done")

		}
	}()

	go func() {
		defer func() {
			fmt.Println("TuneFindMovieHome finished at", time.Since(now))
			wg1.Done()
		}()
		err2 := https.TuneFindMovieHome(&tunefindMovieHome)
		if err2 != nil {
			fmt.Println(err2)

		} else {
			fmt.Println("TuneFindMovieHome Done")

		}
	}()

	go func() {
		defer func() {
			fmt.Println("TuneFindSearchMovie finished at", time.Since(now))
			wg1.Done()
		}()
		err2 := https.TuneFindMovieSearch("sea-change-2017", &tunefindMovieSearch)
		if err2 != nil {
			fmt.Println(err2)

		} else {
			fmt.Println("TuneFindSearchMovie Done")

		}
	}()

	go func() {
		defer func() {
			fmt.Println("TuneFindSearchShow finished at", time.Since(now))
			wg1.Done()
		}()
		err2 := https.TuneFindShowSearch("friends", &tunefindShowSearch)
		if err2 != nil {
			fmt.Println(err2)

		} else {
			fmt.Println("TuneFindSearchShow Done")

		}
	}()

	go func() {
		defer func() {
			fmt.Println("TuneFindSearchSeason finished at", time.Since(now))
			wg1.Done()
		}()
		err2 := https.TuneFindSeasonSearch("friends", 1, &tunefindSeasonSearch)
		if err2 != nil {
			fmt.Println(err2)

		} else {
			fmt.Println("TuneFindSearchSeason Done")

		}
	}()

	go func() {
		defer func() {
			fmt.Println("TuneFindSearchEpisode finished at", time.Since(now))
			wg1.Done()
		}()
		err2 := https.TuneFindEpisodeSearch(22189, &tunefindEpisodeSearch)
		if err2 != nil {
			fmt.Println(err2)

		} else {
			fmt.Println("TuneFindSearchEpisode Done")

		}
	}()

	wg1.Wait()
	customBrowseAll = CustomBrowseAll{
		TunefindMovieBrowseAll: tunefindMovieBrowseAll,
		TunefindShowBrowseAll:  tunefindShowBrowseAll,
		TunefindGameBrowseAll:  tunefindGameBrowseAll,
		TunefindMovieHome:      tunefindMovieHome,
		TunefindGameHome:       tunefindGameHome,
		TunefindShowHome:       tunefindShowHome,
		TunefindSearchMovie:    tunefindMovieSearch,
		TunefindSearchShow:     tunefindShowSearch,
		TunefindSearchSeason:   tunefindSeasonSearch,
		TunefindSearchEpisode:  tunefindEpisodeSearch,
		TunefindTrending:       tunefindTrending,
	}
	customMarshal, _ := json.Marshal(customBrowseAll)
	func() {
		file, err := os.OpenFile(
			"tunefind.json",
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
	fmt.Println("Total Time Taken to Fetch Results and Write JSON to disk TuneFind: ", elapsed)
	fmt.Println("====================================================================")
}

type CustomBrowseAll struct {
	TunefindMovieBrowseAll tunefind.MovieBrowse   `json:"tunefindMovieBrowseAll"`
	TunefindShowBrowseAll  tunefind.ShowBrowse    `json:"tunefindShowBrowseAll"`
	TunefindGameBrowseAll  tunefind.GameBrowse    `json:"tunefindGameBrowseAll"`
	TunefindSearchMovie    tunefind.MovieSearch   `json:"tunefindMovieSearch"`
	TunefindSearchShow     tunefind.ShowSearch    `json:"tunefindShowSearch"`
	TunefindSearchSeason   tunefind.SeasonSearch  `json:"tunefindSeasonSearch"`
	TunefindSearchEpisode  tunefind.EpisodeSearch `json:"tunefindEpisodeSearch"`
	TunefindMovieHome      tunefind.MovieGameHome `json:"tunefindMovieHome"`
	TunefindShowHome       tunefind.ShowHome      `json:"tunefindShowHome"`
	TunefindGameHome       tunefind.MovieGameHome `json:"tunefindGameHome"`
	TunefindTrending       tunefind.TrendingSongs `json:"tunefindTrending"`
}
