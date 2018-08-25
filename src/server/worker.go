package server

import (
	"bufio"
	"fmt"
	"godbms/src/storages"
	"net"
	"strings"
)

type WorkerWrapper struct {
	conn     net.Conn
	storages *storages.Storages
}

func NewConnectionWrapper(storages *storages.Storages, conn net.Conn) *WorkerWrapper {
	return &WorkerWrapper{
		conn:     conn,
		storages: storages,
	}
}

func (wrapper *WorkerWrapper) OnAddToPool() {
}

func (wrapper *WorkerWrapper) OnGetFromPool() {
}

func (wrapper *WorkerWrapper) GetStorages() (storages *storages.Storages) {
	return wrapper.storages
}

func (wrapper *WorkerWrapper) GetConnection() (conn net.Conn) {
	return wrapper.conn
}

func (wrapper *WorkerWrapper) Close() {
	fmt.Println("INFO: connection " + wrapper.conn.RemoteAddr().String() + " closed")
	wrapper.conn.Close()
}

func (wrapper *WorkerWrapper) Run() {
	defer wrapper.Close()
	scanner := bufio.NewScanner(wrapper.conn)

	for scanner.Scan() {
		commandStr := scanner.Text()

		if strings.ToLower(commandStr) == "quit" {
			break
		}

		CommandProcessor(wrapper, commandStr)
	}
}
