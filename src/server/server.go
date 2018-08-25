package server

import (
	"fmt"
	"godbms/src/config"
	"godbms/src/pool"
	"godbms/src/storages"
	"net"
	"strconv"
)

type Server struct {
	config   config.Config
	listener net.Listener
	pool     *pool.Pool
	storages *storages.Storages
}

func NewServer(conf config.Config) *Server {
	return &Server{
		config:   conf,
		pool:     pool.NewPool(conf),
		storages: storages.NewStorages(conf),
	}
}

func (server *Server) Start() (err error) {
	portString := strconv.FormatInt(int64(server.config.Port), 10)
	server.listener, err = net.Listen("tcp", server.config.Host+":"+portString)
	fmt.Println("INFO: Listening on " + server.config.Host + ":" + portString)

	defer server.Close()

	for {
		conn, _ := server.listener.Accept()

		fmt.Println("INFO: connected " + conn.RemoteAddr().String())
		server.pool.AddToProcess(NewConnectionWrapper(server.storages, conn))
	}

	return err
}

func (server *Server) Close() {
	server.listener.Close()
}
