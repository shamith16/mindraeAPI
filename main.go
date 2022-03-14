package main

import "sync"

var wg3 sync.WaitGroup

func main() {
	wg3.Add(2)
	FetchTmdb(&wg3)
	fetchTuneFind(&wg3)
	wg3.Wait()
}
