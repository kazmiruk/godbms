package server

import (
	"fmt"
	"godbms/src/config"
	"godbms/src/worker"
	"net"
	"strconv"
)

type Server struct {
	config   config.Config
	listener net.Listener
}

func NewServer(conf config.Config) *Server {
	return &Server{
		config: conf,
	}
}

func (server *Server) Start() (err error) {
	portString := strconv.FormatInt(int64(server.config.Port), 10)
	server.listener, err = net.Listen("tcp", server.config.Host+":"+portString)
	fmt.Println("Listening on " + server.config.Host + ":" + portString)

	return err
}

func (server *Server) Listen(w worker.Worker) (err error) {
	for {
		conn, err := server.listener.Accept()

		if err != nil {
			break
		}

		go w.Execute(conn)
	}

	return err
}

func (server *Server) Close() {
	server.listener.Close()
}
