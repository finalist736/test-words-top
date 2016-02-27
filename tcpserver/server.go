package tcpserver

import (
	"bufio"
	"net"
	"runtime"

	"finalistx.com/test/storage"
)

var listener net.Listener

func StartServer(addr string) error {
	var err error
	listener, err = net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	go listening()
	return nil
}

func StopServer() {
	listener.Close()
}

func listening() {
	for {
		conn, err := listener.Accept()
		if err != nil {
			runtime.Gosched()
			continue
		}
		go handleConnection(conn)
		runtime.Gosched()
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	line, err := reader.ReadString('\n')
	if err == nil {
		go storage.HandleString(line)
	}
	runtime.Goexit()
}
