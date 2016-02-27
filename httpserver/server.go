package httpserver

import (
	"net/http"
)

func StartServer(addr string) error {
	// add root path handler
	http.HandleFunc("/", TopWords)
	// start server and return error
	return http.ListenAndServe(addr, nil)
}
