package config

import "errors"

const (
	defaultPort int = 3333
)

type Config struct {
	Host string
	Port int
}

func NewConfig(host string, port int) (config *Config, err error) {
	if port < 0 || port > 65535 {
		err = errors.New("Invalid port number " + string(port) + " Should be from 0 till 65535")
		port = defaultPort
	}

	return &Config{
		host,
		port,
	}, err
}
