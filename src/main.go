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
		fmt.Println("ERROR: Error building config: ", err.Error())
		os.Exit(1)
	}

	err = server.NewServer(*conf).Start()

	if err != nil {
		fmt.Println("ERROR: Error starting:", err.Error())
		os.Exit(2)
	}
}
