package config

import "errors"

const (
	defaultPort int = 3333
)

type Config struct {
	Host       string
	Port       int
	MaxWorkers uint
}

func NewConfig(host string, port int, maxWorkers uint) (config *Config, err error) {
	if port < 0 || port > 65535 {
		err = errors.New("Invalid port number " + string(port) + " Should be from 0 till 65535")
		port = defaultPort
	}

	if maxWorkers < 1 {
		maxWorkers = 1
	}

	return &Config{
		host,
		port,
		maxWorkers,
	}, err
}
