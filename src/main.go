package main

import (
	"fmt"
	"godbms/src/config"
	"godbms/src/server"
	"godbms/src/worker"
	"os"
)

func main() {
	conf, err := config.NewConfig("0.0.0.0", 3333)

	if err != nil {
		fmt.Println("Error building config: ", err.Error())
		os.Exit(1)
	}

	serv := server.NewServer(*conf)
	err = serv.Start()

	if err != nil {
		fmt.Println("Error starting:", err.Error())
		os.Exit(2)
	}

	defer serv.Close()
	err = serv.Listen(worker.Worker{})

	if err != nil {
		fmt.Println("Error accepting: ", err.Error())
		os.Exit(3)
	}
}
