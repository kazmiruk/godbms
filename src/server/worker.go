package server

import (
	"bufio"
	"fmt"
	"net"
)

type WorkerWrapper struct {
	conn net.Conn
}

func NewConnectionWrapper(conn net.Conn) *WorkerWrapper {
	return &WorkerWrapper{
		conn: conn,
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
	fmt.Fprintf(wrapper.conn, "Connection waits in the pool\r\n")
}

func (wrapper *WorkerWrapper) OnGetFromPool() {
	fmt.Fprintf(wrapper.conn, "Connection was takken from the pool\r\n")
}

func (wrapper *WorkerWrapper) Run() {
	defer wrapper.conn.Close()
	reader := bufio.NewReader(wrapper.conn)

	for {
		fmt.Fprintf(wrapper.conn, "Enter the word 'quit' (with no quotes) to exit.\r\n")

		str, _ := wrapper.readln(reader)

		if str == "quit" {
			fmt.Println("Quitting.")
			break
		}

		fmt.Println("Input:" + str)
		_, _ = fmt.Fprintf(wrapper.conn, "You said: %s\r\n", str)
	}
}
