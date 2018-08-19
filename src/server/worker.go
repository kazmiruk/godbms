package server

import (
	"bufio"
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

func (wrapper *WorkerWrapper) readln(r *bufio.Reader) (string, error) {
	input, err := r.ReadString('\n')

	if err != nil {
		return "", err
	}

	input = input[:len(input)-1]

	if len(input) > 0 && input[len(input)-1] == '\r' {
		input = input[:len(input)-1]
	}

	return input, nil
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

func (wrapper *WorkerWrapper) Run() {
	defer wrapper.conn.Close()
	reader := bufio.NewReader(wrapper.conn)

	for {
		commandStr, _ := wrapper.readln(reader)

		if strings.ToLower(commandStr) == "quit" {
			break
		}

		CommandProcessor(wrapper, commandStr)
	}
}
