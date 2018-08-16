package main

import (
	"fmt"
	"godbms/src/config"
	"godbms/src/server"
	"os"
)

func main() {
	conf, err := config.NewConfig("0.0.0.0", 3333, 1)

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
}
