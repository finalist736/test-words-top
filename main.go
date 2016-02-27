package main

import (
	"github.com/finalist736/test-words-top/httpserver"
	"github.com/finalist736/test-words-top/storage"
	"github.com/finalist736/test-words-top/tcpserver"
)

func main() {
	// start channel for word storage
	storage.StartDispatching()
	// start tcp server for wors receiving
	err := tcpserver.StartServer(":9000")
	if err != nil {
		panic(err)
	}
	// start http server for words result
	err = httpserver.StartServer(":8000")
	if err != nil {
		tcpserver.StopServer()
		panic(err)
	}
}
