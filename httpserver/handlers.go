package httpserver

import (
	"encoding/json"
	"net/http"
	"strconv"

	"finalistx.com/test/storage"
)

type TopWordsResult struct {
	Words []string `json:"top_words"`
}

func TopWords(w http.ResponseWriter, r *http.Request) {
	// GET param N shows words count
	n_str := r.URL.Query().Get("N")
	n, err := strconv.ParseInt(n_str, 10, 64)
	// check for an errors
	if err != nil {
		w.Write([]byte("specify ?N=<words count>"))
		return
	}
	if n <= 0 {
		w.Write([]byte("incorrect words count, specify greater than zero"))
		return
	}
	// show http header for json
	w.Header().Set("Content-Type", "application/json")
	// getting N top words
	words := storage.TopWords(n)
	// output results
	result, err := json.Marshal(&TopWordsResult{words})
	w.Write(result)
}
