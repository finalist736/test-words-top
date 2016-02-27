package storage

import (
	"math"
	"runtime"
	"sync"
)

type Word struct {
	Word  string
	Count int64
}

var stringsChannel chan []string

//var wordsStorage map[string]int64
var wordsStorage []*Word
var mux sync.RWMutex

func StartDispatching() {
	//wordsStorage = make(map[string]int64)
	stringsChannel = make(chan []string, 10)
	go dispatch()
}

func TopWords(n int64) []string {
	// lock mutex for reading
	mux.RLock()
	defer mux.RUnlock()

	var wordsResult []string
	// calc words in storage
	wordsCount := float64(len(wordsStorage))
	num := int64(math.Min(wordsCount, float64(n)))
	// get first <num> words
	for i := int64(0); i < num; i++ {
		wordsResult = append(wordsResult, wordsStorage[i].Word)
	}
	return wordsResult
}

func dispatch() {
	for {
		select {
		case words := <-stringsChannel:
			// received words list, lock mutex for writing and add words to storage
			mux.Lock()
			for _, iterator := range words {
				var this_word_new bool = true // is this word first time in storage
				for _, storage_iterator := range wordsStorage {
					if storage_iterator.Word == iterator {
						// this word exists, let add 1 to count
						this_word_new = false
						storage_iterator.Count++
						break
					}
				}
				if this_word_new { // this word not exists, add it with count = 1
					wordsStorage = append(wordsStorage, &Word{iterator, 1})
				}
			}
			// sort words storage for fast getting top
			SortWordsByCount(bycount).Sort(wordsStorage)
			// print after sorting
			//for _, iterator := range wordsStorage {
			//fmt.Printf("%+v\n", iterator)
			//}
			mux.Unlock()
		}
		runtime.Gosched()
	}
}
