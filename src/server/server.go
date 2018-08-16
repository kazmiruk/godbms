package server

import (
	"fmt"
	"godbms/src/config"
	"godbms/src/pool"
	"net"
	"strconv"
)

type Server struct {
	config   config.Config
	listener net.Listener
	pool     *pool.Pool
}

func NewServer(conf config.Config) *Server {
	return &Server{
		config: conf,
		pool:   pool.NewPool(conf),
	}
}

func (server *Server) listen() {
	for {
		conn, _ := server.listener.Accept()

		server.pool.Add(NewConnectionWrapper(conn))
	}
}

func (server *Server) process() {
	for {
		server.pool.ProcessNext()
	}
}

func (server *Server) Start() (err error) {
	portString := strconv.FormatInt(int64(server.config.Port), 10)
	server.listener, err = net.Listen("tcp", server.config.Host+":"+portString)
	fmt.Println("Listening on " + server.config.Host + ":" + portString)

	defer server.Close()
	go server.process()
	server.listen()

	return err
}

func (server *Server) Close() {
	server.listener.Close()
}
