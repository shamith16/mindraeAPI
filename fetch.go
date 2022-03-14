package main

import (
	"encoding/json"
	"fmt"
	"github.com/shamith16/mindraeAPI/entities/browse/tmdb"
	"github.com/shamith16/mindraeAPI/entities/browse/tunefind"
	"github.com/shamith16/mindraeAPI/services/https"
	"log"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup

func fetch() {
	{
		wg.Add(5)
		var tmdbMovie tmdb.TmdbMovieBrowse
		var tmdbShow tmdb.TmdbShowBrowse
		var tunefindMovieBrowseAll tunefind.TuneFindMovieBrowse
		var tunefindShowBrowseAll tunefind.TuneFindShowBrowse
		var tunefindGameBrowseAll tunefind.TuneFindGameBrowse
		var customBrowseAll CustomBrowseAll

		now := time.Now()
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
				fmt.Println("TuneFindShowBrowse finished at", time.Since(now))
				wg.Done()
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
				wg.Done()
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
				wg.Done()
			}()
			err2 := https.TuneFindGameBrowse(&tunefindGameBrowseAll)
			if err2 != nil {
				fmt.Println(err2)

			} else {
				fmt.Println("TuneFindGameBrowse Done")

			}
		}()

		wg.Wait()
		customBrowseAll = CustomBrowseAll{
			TmdbMovie:              tmdbMovie,
			TmdbShow:               tmdbShow,
			TunefindMovieBrowseAll: tunefindMovieBrowseAll,
			TunefindShowBrowseAll:  tunefindShowBrowseAll,
			TunefindGameBrowseAll:  tunefindGameBrowseAll,
		}
		customMarshal, _ := json.Marshal(customBrowseAll)
		func() {
			file, err := os.OpenFile(
				"test.json",
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
		fmt.Println("Total Time Taken to Fetch Results and Write JSON to disk: ", elapsed)
		fmt.Println("====================================================================")
	}
}

type CustomBrowseAll struct {
	TmdbMovie              tmdb.TmdbMovieBrowse         `json:"tmdbMovie,omitempty"`
	TmdbShow               tmdb.TmdbShowBrowse          `json:"tmdbShow,omitempty"`
	TunefindMovieBrowseAll tunefind.TuneFindMovieBrowse `json:"tunefindMovieBrowseAll,omitempty"`
	TunefindShowBrowseAll  tunefind.TuneFindShowBrowse  `json:"tunefindShowBrowseAll,omitempty"`
	TunefindGameBrowseAll  tunefind.TuneFindGameBrowse  `json:"tunefindGameBrowseAll,omitempty"`
}
