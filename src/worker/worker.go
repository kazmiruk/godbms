package worker

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

type Worker struct {
}

func readln(r *bufio.Reader) (string, error) {
	input, err := r.ReadString('\n')

	if err != nil {
		return "", err
	}

	input = input[:len(input)-1]

	if input[len(input)-1] == '\r' {
		input = input[:len(input)-1]
	}

	return input, nil
}

func (w *Worker) Execute(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		fmt.Fprintf(conn, "Enter the word 'quit' (with no quotes) to exit.\r\n")

		str, _ := readln(reader)

		if str == "quit" {
			fmt.Println("Quitting.")
			os.Exit(0)
		}

		fmt.Println("Input:" + str)
		_, _ = fmt.Fprintf(conn, "You said: %s\r\n", str)
	}
}
